package lib

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// A trick to disable log messages during testing
func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}
