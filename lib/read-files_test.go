package lib

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/hcl/v2"
)

func TestReadFilesNoFile(t *testing.T) {
	body, err := readFiles()

	if err != nil {
		t.Errorf("Error reading files, \n%s", err)
	}

	if !reflect.DeepEqual(body, hcl.EmptyBody()) {
		t.Errorf("Invalid body, \ngot: %s, \nwant: %s.", body, hcl.EmptyBody())
	}
}

func TestReadFilesWithFilesNoBody(t *testing.T) {
	err := ioutil.WriteFile("test1.hcl", []byte{}, 0644)
	if err != nil {
		t.Errorf("Error creating test files, \n%s", err)
	}

	err = ioutil.WriteFile("test2.hcl", []byte{}, 0644)

	if err != nil {
		t.Errorf("Error creating test files, \n%s", err)
	}

	_, err = readFiles()

	if err != nil {
		t.Errorf("Error reading files, \n%s", err)
	}

	err = os.Remove("test1.hcl")
	if err != nil {
		t.Errorf("Error deleting test files, \n%s", err)
	}

	err = os.Remove("test2.hcl")
	if err != nil {
		t.Errorf("Error deleting test files, \n%s", err)
	}
}

func TestReadFilesWithFileInvalidBody(t *testing.T) {
	err := ioutil.WriteFile("test1.hcl", []byte("Ooh La La Laa."), 0644)

	if err != nil {
		t.Errorf("Error creating test files, \n%s", err)
	}

	_, err = readFiles()

	if err == nil {
		t.Errorf("Should have throw an error, \n%s", err)
	}

	err = os.Remove("test1.hcl")
	if err != nil {
		t.Errorf("Error deleting test files, \n%s", err)
	}
}
