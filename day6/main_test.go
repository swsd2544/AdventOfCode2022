package main

import (
	"fmt"
	"testing"
)

func TestSolveI(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{
			text: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 7,
		},
		{
			text: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 5,
		},
		{
			text: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 6,
		},
		{
			text: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 10,
		},
		{
			text: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 11,
		},
	}

	for id, tt := range tests {
		t.Run(fmt.Sprint(id), func(t *testing.T) {
			got := solveI(tt.text)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}

func TestSolveII(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{
			text: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 19,
		},
		{
			text: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 23,
		},
		{
			text: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 23,
		},
		{
			text: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 29,
		},
		{
			text: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 26,
		},
	}

	for id, tt := range tests {
		t.Run(fmt.Sprint(id), func(t *testing.T) {
			got := solveII(tt.text)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
