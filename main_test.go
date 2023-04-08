package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("epilson")

	wg.Wait()

	if msg != "epilson" {
		t.Error("Expected to find epilson, but is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epilson"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "epilson") {
		t.Error("Expected to find epilson, but is not there")
	}
}

func TestMain(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world!, but is not there")
	}

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe!, but is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos!, but is not there")
	}
}
