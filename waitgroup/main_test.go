package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestPrintsome(t *testing.T) {
	stdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	intput := []string{"Valu", "bai", "Mitra", "789", "Deba mm"}
	go Printsome(intput[0], &wg)
	wg.Add(1)

	wg.Wait()
	writer.Close()

	result, _ := io.ReadAll(reader)
	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, intput[0]) {
		t.Errorf("Not expected result")
	}
}
