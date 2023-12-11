package main

import (
	"reflect"
	"testing"
)

func TestExpandUniverse(t *testing.T) {
	input := [][]string{
		{".", ".", ".", "#", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		{".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{"#", ".", ".", ".", "#", ".", ".", ".", ".", "."},
	}
	expected := [][]string{
		{".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", ".", "."},
		{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		{"#", ".", ".", ".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
	}
	got := expandUniverse(input)
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("test expandUniverse failed. got=(%d,%d) expected=(%d,%d)", len(got), len(got[0]), len(expected), len(expected[0]))
	}
}

func TestGetGalaxiesCoords(t *testing.T) {
	input := [][]string{
		{".", ".", ".", "#", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		{".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
		{"#", ".", ".", ".", "#", ".", ".", ".", ".", "."},
	}
	expected := map[int]Coord{
		1: {x: 0, y: 3},
		2: {x: 1, y: 7},
		3: {x: 2, y: 0},
		4: {x: 4, y: 6},
		5: {x: 5, y: 1},
		6: {x: 6, y: 9},
		7: {x: 8, y: 7},
		8: {x: 9, y: 0},
		9: {x: 9, y: 4},
	}
	got := getGalaxiesCoords(input)
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("test getGalaxiesCoords failed. got=%v expected=%v", got, expected)
	}
}
