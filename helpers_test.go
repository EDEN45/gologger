package gologger

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_toJSON(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want []byte
	}{
		{name: "test", args: []any{"One", 1, "Two", "dsds"}, want: []byte(`{"One":"1","Two":"dsds"}` + "\r\n")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toJSON(tt.args); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(reflect.DeepEqual(string(got), string(tt.want)))
				t.Errorf("toJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
