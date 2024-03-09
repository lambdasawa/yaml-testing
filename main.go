package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"time"
)

type Result struct {
	DataName    string
	ProgramName string
	Data        string
	ExitCode    int
	Timeout     bool
	Stdout      string
	Stderr      string
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	const (
		programDirName  = "program"
		testdataDirName = "testdata"
	)

	var results []Result

	testdata, err := os.ReadDir(testdataDirName)
	if err != nil {
		return err
	}

	programs, err := os.ReadDir(programDirName)
	if err != nil {
		return err
	}

	for _, data := range testdata {
		dataContent, err := os.ReadFile(path.Join(testdataDirName, data.Name()))
		if err != nil {
			return err
		}

		for _, program := range programs {
			result := execute(string(dataContent), data.Name(), program.Name(), testdataDirName, programDirName)

			log.Printf("data:%-40s program:%-20s", result.DataName, result.ProgramName)

			results = append(results, result)
		}
	}

	if err := exportAsTSV(results); err != nil {
		return err
	}

	return nil
}

func execute(data string, dataName string, programName string, testdataDirName string, programDirName string) Result {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill)
	defer cancel()

	var stdout, stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, "bash", "main.sh", path.Join("..", "..", testdataDirName, dataName))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = path.Join(programDirName, programName)

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	var timeout bool
	select {
	case <-done:
		timeout = false
	case <-time.After(3 * time.Second):
		go cmd.Process.Kill()
		timeout = true
	}

	return Result{
		DataName:    dataName,
		ProgramName: programName,
		Data:        data,
		ExitCode:    cmd.ProcessState.ExitCode(),
		Timeout:     timeout,
		Stdout:      takeInitText(stdout.String()),
		Stderr:      stderr.String(),
	}
}

func takeInitText(text string) string {
	const maxLength = 512

	if len(text) < maxLength {
		return text
	}

	return text[:maxLength]
}

func exportAsTSV(results []Result) error {
	file, err := os.Create("results.tsv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Comma = '\t'

	if err := writer.Write([]string{"data name", "program name", "data", "stdout", "stderr", "exit code", "timeout"}); err != nil {
		return err
	}

	for _, result := range results {
		if err := writer.Write([]string{
			result.DataName,
			result.ProgramName,
			result.Data,
			result.Stdout,
			result.Stderr,
			fmt.Sprintf("%v", result.ExitCode),
			fmt.Sprintf("%v", result.Timeout),
		}); err != nil {
			return err
		}
	}

	return nil
}
