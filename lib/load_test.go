package lib

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_compareVersion(t *testing.T) {
	type args struct {
		constraint string
		actual     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"correct", args{"1.0.0", "v1.0.0"}, false},
		{"correct", args{"^1.0.0", "v1.1.2"}, false},
		{"correct", args{"^1.5.0", "v1.6.2"}, false},
		{"correct", args{"=1.5.0", "v1.5.0"}, false},
		{"correct", args{"=1.5.0", "v1.5.5"}, true},
		{"correct", args{"!=1.5.0", "v1.5.5"}, false},
		{"correct", args{"~1.5", "v1.5.5"}, false},
		{"no err on empty actual", args{"~1.5", ""}, false},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := compareVersion(tt.args.constraint, tt.args.actual); (err != nil) != tt.wantErr {
					t.Errorf("compareVersion() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func createTestFile(content []byte) {
	ioutil.WriteFile("test1.hcl", content, 0644)
}

func removeTestFile() {
	os.Remove("test1.hcl")
}

func Test_readAndDecodeBody(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		rootCfg *RootCfg
		wantErr bool
	}{
		{"with version", []byte(`version = "1.0.0"`), &RootCfg{Version: "1.0.0"}, false},
		{"with error", []byte(`unknown thing`), nil, true},
		{"with diag", []byte(`unknown_thing = "hi"`), nil, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				createTestFile(tt.content)
				got, err := readAndDecodeBody()

				if (err != nil) != tt.wantErr {
					t.Errorf("readAndDecodeBody() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(got, tt.rootCfg) {
					t.Errorf("readAndDecodeBody() got = %v, want %v", got, tt.rootCfg)
				}

				removeTestFile()
			},
		)
	}
}

func TestListRequestsAndTests(t *testing.T) {
	testHcl := `
	version = "1.0.0"
	request test-1 {}
	request test-2 {}
	test test-1 {}
	test test-2 {}
`
	type args struct {
		lsType  string
		execCtx *ExecutionContext
	}
	tests := []struct {
		name           string
		args           args
		wantList       []ListObject
		wantMaxNameLen int
		contents       []byte
		wantErr        bool
	}{
		{
			"only requests", args{"request", &ExecutionContext{Version: "1.0.0"}},
			[]ListObject{{"test-1", "request"}, {"test-2", "request"}}, 6, []byte(testHcl), false,
		},
		{
			"only tests", args{"test", &ExecutionContext{Version: "1.0.0"}},
			[]ListObject{{"test-1", "test"}, {"test-2", "test"}}, 6, []byte(testHcl), false,
		},
	}
	for _, tt := range tests {
		createTestFile(tt.contents)
		defer removeTestFile()
		t.Run(
			tt.name, func(t *testing.T) {
				gotList, gotMaxNameLen, err := ListRequestsAndTests(tt.args.lsType, tt.args.execCtx)
				if (err != nil) != tt.wantErr {
					t.Errorf("ListRequestsAndTests() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotList, tt.wantList) {
					t.Errorf("ListRequestsAndTests() gotList = %v, want %v", gotList, tt.wantList)
				}
				if gotMaxNameLen != tt.wantMaxNameLen {
					t.Errorf("ListRequestsAndTests() gotMaxNameLen = %v, want %v", gotMaxNameLen, tt.wantMaxNameLen)
				}
			},
		)
	}
}
