package main

import (
	"testing"
)

func TestParsed6(t *testing.T) {
	rollSpec, err := Parse("d6")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 1 {
		t.Fatalf("Expected die count of 1, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 6 {
		t.Fatalf("Expected sides of 6, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 0 {
		t.Fatalf("Expected modifier of 0, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 0 {
		t.Fatalf("Expected best of of 0, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 1 {
		t.Fatalf("Expected Times of 1, got %d", rollSpec.Times)
	}
}

func TestParse1d6(t *testing.T) {
	rollSpec, err := Parse("1d6")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 1 {
		t.Fatalf("Expected die count of 1, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 6 {
		t.Fatalf("Expected sides of 6, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 0 {
		t.Fatalf("Expected modifier of 0, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 0 {
		t.Fatalf("Expected best of of 0, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 1 {
		t.Fatalf("Expected Times of 1, got %d", rollSpec.Times)
	}
}

func TestParse1d6Plus1(t *testing.T) {
	rollSpec, err := Parse("d6+1")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 1 {
		t.Fatalf("Expected die count of 1, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 6 {
		t.Fatalf("Expected sides of 6, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 1 {
		t.Fatalf("Expected modifier of 1, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 0 {
		t.Fatalf("Expected best of of 0, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 1 {
		t.Fatalf("Expected Times of 1, got %d", rollSpec.Times)
	}
}

func TestParse3d6(t *testing.T) {
	rollSpec, err := Parse("3d6")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 3 {
		t.Fatalf("Expected die count of 3, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 6 {
		t.Fatalf("Expected sides of 6, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 0 {
		t.Fatalf("Expected modifier of 0, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 0 {
		t.Fatalf("Expected best of of 0, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 1 {
		t.Fatalf("Expected Times of 1, got %d", rollSpec.Times)
	}
}

func TestParse3d6SixTimes(t *testing.T) {
	rollSpec, err := Parse("3d6x6")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 3 {
		t.Fatalf("Expected die count of 3, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 6 {
		t.Fatalf("Expected sides of 6, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 0 {
		t.Fatalf("Expected modifier of 0, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 0 {
		t.Fatalf("Expected best of of 0, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 6 {
		t.Fatalf("Expected Times of 6, got %d", rollSpec.Times)
	}
}

func TestParseBest3Of4d6SixTimes(t *testing.T) {
	rollSpec, err := Parse("3,4d6x6")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 4 {
		t.Fatalf("Expected die count of 4, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 6 {
		t.Fatalf("Expected sides of 6, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 0 {
		t.Fatalf("Expected modifier of 0, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 3 {
		t.Fatalf("Expected best of of 3, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 6 {
		t.Fatalf("Expected Times of 6, got %d", rollSpec.Times)
	}
}

func TestParseTwoBestOfThreed10Plus3FiveTimes(t *testing.T) {
	rollSpec, err := Parse("2,3d10+3x5")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 3 {
		t.Fatalf("Expected die count of 3, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 10 {
		t.Fatalf("Expected sides of 10, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 3 {
		t.Fatalf("Expected modifier of 3, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 2 {
		t.Fatalf("Expected best of of 2, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 5 {
		t.Fatalf("Expected Times of 5, got %d", rollSpec.Times)
	}
}

func TestBasicParse2d8Plus1TwoTimes(t *testing.T) {
	rollSpec, err := Parse("2d8+3x2")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if rollSpec.DieCount != 2 {
		t.Fatalf("Expected die count of 2, got %d", rollSpec.DieCount)
	}
	if rollSpec.Sides != 8 {
		t.Fatalf("Expected sides of 8, got %d", rollSpec.Sides)
	}
	if rollSpec.Modifier != 3 {
		t.Fatalf("Expected modifier of 3, got %d", rollSpec.Modifier)
	}
	if rollSpec.BestOf != 0 {
		t.Fatalf("Expected best of of 0, got %d", rollSpec.BestOf)
	}
	if rollSpec.Times != 2 {
		t.Fatalf("Expected Times of 2, got %d", rollSpec.Times)
	}
}
