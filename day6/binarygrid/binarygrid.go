package binarygrid

type Grid struct {
	lights [][]bool
}

func NewGrid(size int) *Grid {
	lights := make([][]bool, size)
	for i := range lights {
		lights[i] = make([]bool, size)
	}
	return &Grid{lights: lights}
}

func (g *Grid) IsOn(x, y int) bool {
	return g.lights[x][y]
}

func (g *Grid) TurnOnSquare(x1, y1, x2, y2 int) {
	g.mapRectangle(x1, y1, x2, y2, func(light bool) bool {
		return true
	})
}

func (g *Grid) TurnOffSquare(x1, y1, x2, y2 int) {
	g.mapRectangle(x1, y1, x2, y2, func(light bool) bool {
		return false
	})
}

func (g *Grid) ToggleSquare(x1, y1, x2, y2 int) {
	g.mapRectangle(x1, y1, x2, y2, func(light bool) bool {
		return !light
	})
}

func (g *Grid) TotalBrightness() int {
	total := 0
	for i := 0; i < len(g.lights); i++ {
		for j := 0; j < len(g.lights[i]); j++ {
			if g.lights[i][j] {
				total++
			}
		}
	}
	return total
}

func (g *Grid) mapRectangle(x1, y1, x2, y2 int, mapper func(light bool) bool) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			g.lights[i][j] = mapper(g.lights[i][j])
		}
	}
}
