package godev

import (
	"os"
	"os/exec"

	"github.com/cdvelop/compiler"
	"github.com/cdvelop/dev_browser"
	"github.com/cdvelop/token"
	"github.com/cdvelop/watch_files"
)

type Dev struct {
	app_path string //ej: app.exe

	*dev_browser.Browser
	*watch_files.WatchFiles
	*compiler.Compiler

	*exec.Cmd

	// Scanner   *bufio.Scanner
	Interrupt chan os.Signal

	ProgramStartedMessages chan string

	run_arguments []string

	*token.TwoKeys
}
