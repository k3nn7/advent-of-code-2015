package main

import "testing"

func TestCompareHouses(t *testing.T) {
	house1 := House{4, 5}
	house2 := House{3, 4}
	house3 := House{4, 5}

	if house1 == house2 {
		t.Errorf("House %#v and %#v should be different", house1, house2)
	}

	if house1 != house3 {
		t.Errorf("House %#v and %#v should be the same", house1, house3)
	}
}

func TestNewWorldHasNoVistiedHouses(t *testing.T) {
	world := World{}
	if world.VisitedHouses() != 0 {
		t.Errorf("Visited houses should be 0 got %d", world.VisitedHouses())
	}
}

func TestVisitOneHouse(t *testing.T) {
	world := World{}
	world.Visit(House{1, 1})
	if world.VisitedHouses() != 1 {
		t.Errorf("Visited houses should be 1 got %d", world.VisitedHouses())
	}
}

func TestVisitSameHouseTwice(t *testing.T) {
	world := World{}
	world.Visit(House{1, 1})
	world.Visit(House{1, 1})
	if world.VisitedHouses() != 1 {
		t.Errorf("Visited houses should be 1 got %d", world.VisitedHouses())
	}
}

func TestCursorStartsAt0(t *testing.T) {
	cursor := Cursor{}
	assertPosition(0, 0, cursor, t)
}

func TestParseCursorCommand(t *testing.T) {
	cursor := Cursor{}
	cursor.Cmd("^")
	assertPosition(0, 1, cursor, t)
	cursor.Cmd("v")
	assertPosition(0, 0, cursor, t)
	cursor.Cmd(">")
	assertPosition(1, 0, cursor, t)
	cursor.Cmd("<")
	assertPosition(0, 0, cursor, t)
}

func assertPosition(x, y int, c Cursor, t *testing.T) {
	if c.X != x || c.Y != y {
		t.Errorf("Cursor should be at %d,%d but is at %d,%d", x, y, c.X, c.Y)
	}
}
