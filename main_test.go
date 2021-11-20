package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_withFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "al-count",
			args: args{filePath: "./tests/al-count-0.b"},
			want: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},	{
			name: "al-count-1",
			args: args{filePath: "./tests/al-count-1.b"},
			want: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},{
			name: "hello",
			args: args{filePath: "./tests/Hello.b"},
			want: "Hello World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData := withFile(tt.args.filePath, "")
			if !reflect.DeepEqual(strings.TrimSpace(gotData), strings.TrimSpace(tt.want)) {
				t.Errorf("withFile() gotData = %v, want %v", gotData, tt.want)
			}
		})
	}
}
