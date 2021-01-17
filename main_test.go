package main

import (
	"os"
	"strings"
	"testing"
)

func TestRecursiveScanFolder(t *testing.T) {
	dir, _ := os.Getwd()

	outPutDir := strings.Join(recursiveScanFolder(dir), "")

	if outPutDir != dir {
		t.Error("Has not found correct repo foolder", dir)
	}
}
