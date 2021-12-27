package main

import (
	"reflect"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		source []string
		target []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Normal",
			args: args{
				source: []string{"1", "15", "17", "200", "2", "302", "5"},
				target: []string{"2", "100", "15", "302", "5"},
			},
			want: []string{"15", "2", "302", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.source, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
