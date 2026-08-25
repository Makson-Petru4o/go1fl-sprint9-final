package main

import "testing"

func TestMaximum(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "positive elements", data: []int{3, 17, 8, 5}, want: 17},
		{name: "negative elements", data: []int{-9, -2, -15, -4}, want: -2},
		{name: "mixed elements", data: []int{-10, 0, 25, -3}, want: 25},
		{name: "duplicate maximum", data: []int{7, 2, 7, 1}, want: 7},
		{name: "integer limits", data: []int{minInt, 0, maxInt}, want: maxInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d; want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "positive size", size: 10, want: 10},
		{name: "zero size", size: 0, want: 0},
		{name: "negative size", size: -1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if len(data) != tt.want {
				t.Fatalf("generateRandomElements(%d) returned a slice of length %d; want %d", tt.size, len(data), tt.want)
			}

			if data == nil {
				t.Fatalf("generateRandomElements(%d) returned nil; want a non-nil slice", tt.size)
			}
		})
	}
}
