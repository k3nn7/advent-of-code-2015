package brightnessgrid

import "testing"

func TestAllLightsOffAtBeginning(t *testing.T) {
	grid := NewGrid(10)

	eachInRectangle(*grid, 0, 0, 9, 9, func(x, y, brightness int) {
		if brightness != 0 {
			t.Errorf("Light at %d,%d should has brightness 0, has: %d", x, y, brightness)
		}
	})

	if grid.TotalBrightness() != 0 {
		t.Errorf("Lights on should be 0 but got %d", grid.TotalBrightness())
	}
}

func TestTurnOnSquare(t *testing.T) {
	grid := NewGrid(10)

	grid.TurnOnSquare(0, 0, 2, 2)

	eachInRectangle(*grid, 0, 0, 2, 2, func(x, y int, brightness int) {
		if brightness != 1 {
			t.Errorf("Light at %d,%d should has brighness 1, has: %d", x, y, brightness)
		}
	})

	eachInRectangle(*grid, 3, 3, 9, 9, func(x, y int, brightness int) {
		if brightness != 0 {
			t.Errorf("Light at %d,%d should has brighness 0, has: %d", x, y, brightness)
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

	eachInRectangle(*grid, 0, 0, 1, 1, func(x, y int, brightness int) {
		if brightness != 0 {
			t.Errorf("Light at %d,%d should has brighness 0, has: %d", x, y, brightness)
		}
	})

	eachInRectangle(*grid, 0, 2, 0, 2, func(x, y int, brightness int) {
		if brightness != 1 {
			t.Errorf("Light at %d,%d should has brighness 1, has: %d", x, y, brightness)
		}
	})

	eachInRectangle(*grid, 2, 0, 2, 2, func(x, y int, brightness int) {
		if brightness != 1 {
			t.Errorf("Light at %d,%d should has brighness 1, has: %d", x, y, brightness)
		}
	})

	eachInRectangle(*grid, 3, 3, 9, 9, func(x, y int, brightness int) {
		if brightness != 0 {
			t.Errorf("Light at %d,%d should has brighness 0, has: %d", x, y, brightness)
		}
	})
}

func TestToggleSquare(t *testing.T) {
	grid := NewGrid(10)

	grid.TurnOnSquare(0, 0, 2, 2)
	grid.TurnOffSquare(0, 0, 1, 1)
	grid.ToggleSquare(0, 0, 9, 9)

	eachInRectangle(*grid, 0, 0, 1, 1, func(x, y int, brightness int) {
		if brightness != 2 {
			t.Errorf("Light at %d,%d should has brighness 2, has: %d", x, y, brightness)
		}
	})

	eachInRectangle(*grid, 0, 2, 2, 2, func(x, y int, brightness int) {
		if brightness != 3 {
			t.Errorf("Light at %d,%d should has brighness 3, has: %d", x, y, brightness)
		}
	})

	eachInRectangle(*grid, 2, 0, 2, 2, func(x, y int, brightness int) {
		if brightness != 3 {
			t.Errorf("Light at %d,%d should has brighness 3, has: %d", x, y, brightness)
		}
	})

	eachInRectangle(*grid, 3, 3, 9, 9, func(x, y int, brightness int) {
		if brightness != 2 {
			t.Errorf("Light at %d,%d should has brighness 2, has: %d", x, y, brightness)
		}
	})
}

func TestTotalBrightness(t *testing.T) {
	grid := NewGrid(1000)
	grid.ToggleSquare(0, 0, 999, 999)
	expectedTotalBrightness := 2000000
	totalBrightness := grid.TotalBrightness()
	if expectedTotalBrightness != totalBrightness {
		t.Errorf("Expected total brightness: %d got: %d", expectedTotalBrightness, totalBrightness)
	}
}

func eachInRectangle(grid Grid, x1, y1, x2, y2 int, handler func(x, y, brightness int)) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			handler(i, j, grid.Brightness(i, j))
		}
	}
}
