package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string // ชื่อเคส (เพื่อให้รู้ว่าเทสเรื่องอะไร)
		a, b int    // โจทย์
		want int    // คำตอบที่คาดหวัง
	}{
		{"Positive numbers", 2, 3, 5},
		{"Negative numbers", -1, -2, -3},
		{"Mixed numbers", -5, 5, 0},
		{"Zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Positive", 2, 4, -2},
		{"Negative", -1, -2, 1},
		{"Mixed", -2, 2, -4},
		{"Zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d;want = %d", tt.a, tt.b, got, tt.want)
			}
		})

	}
}
