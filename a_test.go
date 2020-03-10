package main

import (
	"reflect"
	"testing"
)

func Test_a_isA(t *testing.T) {
	var tt *bool

	type fields struct {
		bool *bool
	}
	var tests = []struct {
		name   string
		fields fields
		want   *bool
	}{
		{
			name: "false!",
			fields: fields{
				bool:tt,
			},
			want: tt,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := a{
				bool: tt.fields.bool,
			}
			if got := a.isA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isA() = %v, want %v", got, tt.want)
			}
		})
	}
}