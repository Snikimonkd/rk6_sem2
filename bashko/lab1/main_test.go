package main

import (
	"reflect"
	"testing"
)

func Test_newPos(t *testing.T) {
	type args struct {
		cords Cord
		n     int
		m     int
	}
	tests := []struct {
		name string
		args args
		want []Cord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPos(tt.args.cords, tt.args.n, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
