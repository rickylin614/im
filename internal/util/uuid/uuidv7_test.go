package uuid

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"check len", 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); len(got) != (tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}
