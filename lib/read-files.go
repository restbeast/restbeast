package lib

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"path/filepath"
)

func readFiles() (body hcl.Body, err error) {
	files, err := filepath.Glob(filepath.Join("./", "*.hcl"))
	if err != nil {
		return nil, err
	}

	var parsedFiles []*hcl.File

	// Create a HCL parser and parse all files.
	parser := hclparse.NewParser()
	for _, fileName := range files {
		f, diags := parser.ParseHCLFile(fileName)
		if diags != nil {
			return nil, fmt.Errorf("can not parse parsing hcl file %s\n%s", fileName, diags)
		}

		parsedFiles = append(parsedFiles, f)
	}

	// Merge bodies of parsed files
	mergedBody := hcl.MergeFiles(parsedFiles)

	return mergedBody, nil
}
