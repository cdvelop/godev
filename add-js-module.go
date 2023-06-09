package godev

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/cdvelop/model"
)

// path ej: "modules/users/js_module","ui/components/form/js_module"
func attachJsToModuleFromFolder(comp *model.Object, module_name string, funtions, listener_add, listener_rem *bytes.Buffer) {

	// adjuntar js desde carpeta
	if comp.Path != nil && comp.Path.FolderPath() != "" {
		path := comp.Path.FolderPath() + "/js_module"

		files, err := os.ReadDir(path)
		if err == nil {
			// fmt.Printf("directorio %v de %v no encontrado\n", path, module.MainName)

			for _, file := range files {

				data, err := os.ReadFile(path + "/" + file.Name())
				if err != nil {
					log.Fatalf("error: archivo %v/%v no existe %v", path, file.Name(), err)
				}

				parsed_js := parseModuleJS(parseJS{
					ModuleName: module_name,
					FieldName:  "",
				}, data)

				if strings.Contains(file.Name(), "add-listener") {

					listener_add.WriteString(parsed_js.String() + "\n")

					// reemplazar todas las ocurrencias de "addEventListener" por "removeEventListener"
					rem_listeners := strings.ReplaceAll(parsed_js.String(), "addEventListener", "removeEventListener")

					listener_rem.WriteString(rem_listeners + "\n")
				} else {

					funtions.WriteString(parsed_js.String() + "\n")
				}

			}

		}

	}

}
