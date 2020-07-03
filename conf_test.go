package conf

import (
	"os"
	"testing"
)

type testOptions struct {
	Hello string `json:"hello"`
}

func TestLoad(t *testing.T) {
	os.Setenv("CONF_DIR", "testdata")
	var opts testOptions
	err := Load("module2", &opts)
	if err != nil {
		t.Fatal(err)
	}
	if opts.Hello != "world" {
		t.Fatal("not equal world")
	}
	err = Load("module1", &opts)
	if err != nil {
		t.Fatal(err)
	}
	if opts.Hello != "world" {
		t.Fatal("not equal world")
	}
	os.Unsetenv("CONF_DIR")
	err = Load("module1", &opts)
	if err != os.ErrNotExist {
		t.Fatal(err)
	}
}
