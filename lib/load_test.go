package lib

import "testing"

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
		t.Run(tt.name, func(t *testing.T) {
			if err := compareVersion(tt.args.constraint, tt.args.actual); (err != nil) != tt.wantErr {
				t.Errorf("compareVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
