package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTask322(t *testing.T){
	var tests = []struct{
		n string
		want int
	}{
		{"10000", 9240},
		{"1", 1},
		{"-100", 0},
		{"20000", 0},
		{"5",4},
		{"17",16},
		{"32",30},
		{"256",240},
	}

	for _, tt := range tests{
		testName := fmt.Sprintf("%s", tt.n)
		t.Run(testName, func(t *testing.T) {
			new, _ := strconv.Atoi(tt.n)
			if new > 1 || new < 10000 {
				ans := task322(tt.n)
				if ans != tt.want {
					t.Errorf("got %d, want %d", ans, tt.want)
				}
			} else {
				t.Errorf("Need a number between 1 .. 10000")
			}
		})
	}
}