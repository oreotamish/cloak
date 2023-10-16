package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = origStdout

	expected := "hello world\n"
	if string(out) != expected {
		t.Errorf("Expected %q but got %q", expected, out)
	}
}
