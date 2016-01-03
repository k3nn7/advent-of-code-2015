package main

import (
	"reflect"
	"testing"
)

func TestParseDimensions(t *testing.T) {
	input := "30x5x18"
	expected := Dimensions{30, 5, 18}
	actual := ParseDimensions(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expecting %#v but get %#v", expected, actual)
	}
}

func TestPaperRequired(t *testing.T) {
	d := Dimensions{2, 3, 4}
	expected := 58
	actual := d.PaperRequired()
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestPaperRequired2(t *testing.T) {
	d := Dimensions{1, 1, 10}
	expected := 43
	actual := d.PaperRequired()
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestRibbonRequired(t *testing.T) {
	d := Dimensions{2, 3, 4}
	expected := 34
	actual := d.RibbonRequired()
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestRibbonRequired2(t *testing.T) {
	d := Dimensions{1, 1, 10}
	expected := 14
	actual := d.RibbonRequired()
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
