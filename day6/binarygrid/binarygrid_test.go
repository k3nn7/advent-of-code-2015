package binarygrid

import "testing"

func TestAllLightsOffAtBeginning(t *testing.T) {
	grid := NewGrid(10)

	eachInRectangle(*grid, 0, 0, 9, 9, func(x, y int, isOn bool) {
		if isOn {
			t.Errorf("Light at %d,%d should be off", x, y)
		}
	})

	if grid.TotalBrightness() != 0 {
		t.Errorf("Lights on should be 0 but got %d", grid.TotalBrightness())
	}
}

func TestTurnOnSquare(t *testing.T) {
	grid := NewGrid(10)

	grid.TurnOnSquare(0, 0, 2, 2)

	eachInRectangle(*grid, 0, 0, 2, 2, func(x, y int, isOn bool) {
		if !isOn {
			t.Errorf("Light at %d,%d should be on", x, y)
		}
	})

	eachInRectangle(*grid, 3, 3, 9, 9, func(x, y int, isOn bool) {
		if isOn {
			t.Errorf("Light at %d,%d should be off", x, y)
		}
	})

	if grid.TotalBrightness() != 9 {
		t.Errorf("Lights on should be 9 but got %d", grid.TotalBrightness())
	}
}

func TestTurnOffSquare(t *testing.T) {
	grid := NewGrid(10)

	grid.TurnOnSquare(0, 0, 2, 2)
	grid.TurnOffSquare(0, 0, 1, 1)

	eachInRectangle(*grid, 0, 0, 1, 1, func(x, y int, isOn bool) {
		if isOn {
			t.Errorf("Light at %d,%d should be off", x, y)
		}
	})

	eachInRectangle(*grid, 0, 2, 0, 2, func(x, y int, isOn bool) {
		if !isOn {
			t.Errorf("Light at %d,%d should be on", x, y)
		}
	})

	eachInRectangle(*grid, 2, 0, 2, 2, func(x, y int, isOn bool) {
		if !isOn {
			t.Errorf("Light at %d,%d should be on", x, y)
		}
	})

	eachInRectangle(*grid, 3, 3, 9, 9, func(x, y int, isOn bool) {
		if isOn {
			t.Errorf("Light at %d,%d should be off", x, y)
		}
	})
}

func TestToggleSquare(t *testing.T) {
	grid := NewGrid(10)

	grid.TurnOnSquare(0, 0, 2, 2)
	grid.TurnOffSquare(0, 0, 1, 1)
	grid.ToggleSquare(0, 0, 9, 9)

	eachInRectangle(*grid, 0, 0, 1, 1, func(x, y int, isOn bool) {
		if !isOn {
			t.Errorf("Light at %d,%d should be on", x, y)
		}
	})

	eachInRectangle(*grid, 0, 2, 2, 2, func(x, y int, isOn bool) {
		if isOn {
			t.Errorf("Light at %d,%d should be off", x, y)
		}
	})

	eachInRectangle(*grid, 2, 0, 2, 2, func(x, y int, isOn bool) {
		if isOn {
			t.Errorf("Light at %d,%d should be off", x, y)
		}
	})

	eachInRectangle(*grid, 3, 3, 9, 9, func(x, y int, isOn bool) {
		if !isOn {
			t.Errorf("Light at %d,%d should be on", x, y)
		}
	})
}

func eachInRectangle(grid Grid, x1, y1, x2, y2 int, handler func(x, y int, isOn bool)) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			handler(i, j, grid.IsOn(i, j))
		}
	}
}
