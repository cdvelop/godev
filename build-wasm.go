package godev

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const js_wasm_format = `const go = new Go();
WebAssembly.instantiateStreaming(fetch("static/app.wasm"), go.importObject).then((result) => {
	go.run(result.instance);
});`

const wasm_js_file = "/wasm/wasm_exec.js"

const WASM_FILE_NAME = "/wasm_main.go"

func (u *ui) webAssemblyCheck() {
	// chequear si existe main.go en la ruta de trabajo ej: frontend/main.go
	_, err := os.Open(WORK_FOLDER + WASM_FILE_NAME)
	if err == nil {
		var compiler string

		u.wasm_build = true

		// tiny Go Check
		_, err := os.ReadFile(u.theme_folder + "/wasm/wasm_exec_tinygo.js")
		if err == nil {
			u.with_tinyGo = true
			compiler = "TinyGo"
		} else {
			compiler = "Go"
		}

		fmt.Printf("*** Compilador: [%v] WebAssembly Activado ***\n", compiler)
	}

}

func (u *ui) addWasmJS(out_js *bytes.Buffer) {
	var err error
	// si existen los archivos js wasm agregamos la llamada a estos
	err = readFile(u.theme_folder+"/wasm/wasm_exec_tinygo.js", out_js)
	if err == nil {
		// fmt.Println("*** COMPILACIÓN WASM TINYGO ***")
		out_js.WriteString(js_wasm_format)

	} else {

		err = readFile(u.theme_folder+wasm_js_file, out_js)
		if err == nil {
			// fmt.Println("*** COMPILACIÓN WASM GO ***")
			out_js.WriteString(js_wasm_format)
		}
	}

	if err != nil {
		log.Println("addWasmJS error: ", err)
	}

}

func (u ui) BuildWASM() {
	err := u.buildWASM(WORK_FOLDER+WASM_FILE_NAME, STATIC_FOLDER+"/app.wasm")
	if err != nil {
		log.Println("BuildWASM error: ", err)
	}
}

func (u ui) buildWASM(input_go_file string, out_wasm_file string) error {

	var cmd *exec.Cmd

	// fmt.Println("WITH TINY GO?: ", u.with_tinyGo)
	// Ajustamos los parámetros de compilación según la configuración
	if u.with_tinyGo {
		// fmt.Println("*** COMPILACIÓN WASM TINYGO ***")
		cmd = exec.Command("tinygo", "build", "-o", out_wasm_file, "-target", "wasm", input_go_file)

	} else {
		// compilación normal...
		// fmt.Println("*** COMPILACIÓN WASM GO ***")
		cmd = exec.Command("go", "build", "-o", out_wasm_file, "-tags", "dev", "-ldflags", "-s -w", "-v", input_go_file)
		cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		// if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, string(output))
		return fmt.Errorf("al compilar a WebAssembly: %v", err)
	}

	// Verificamos si el archivo wasm se creó correctamente
	if _, err := os.Stat(out_wasm_file); err != nil {
		return fmt.Errorf("el archivo WebAssembly no se creó correctamente: %v", err)
	}

	// fmt.Printf("WebAssembly compilado correctamente y guardado en %s\n", out_wasm_file)

	return nil
}
