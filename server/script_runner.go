package server

import (
	"log"
	"os"
	"path"

	"github.com/fer2d2/micro-yodeler/command"
)

type ScriptRunner struct {
	Path string
}

func NewScriptRunner(scriptsDirectoryPath string) ScriptRunner {
	scriptRunner := new(ScriptRunner)
	scriptRunner.Path = path.Clean(scriptsDirectoryPath)

	if _, err := os.Stat(scriptRunner.Path); err != nil {
		log.Fatalf("Error found: %s", err)
	}

	return *scriptRunner
}

func (scriptRunner ScriptRunner) Execute(fileName string, result *command.Result) error {
	fileName = path.Base(fileName) // Prevent traversal path attacks

	stdOut, stdErr, exitCode := command.Run("sh", "-c", scriptRunner.Path+"/"+fileName)

	result.StdOut = stdOut
	result.StdErr = stdErr
	result.ExitCode = exitCode

	return nil
}
