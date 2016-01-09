package brightnessgrid

type Grid struct {
	lights [][]int
}

func NewGrid(size int) *Grid {
	lights := make([][]int, size)
	for i := range lights {
		lights[i] = make([]int, size)
	}
	return &Grid{lights: lights}
}

func (g *Grid) Brightness(x, y int) int {
	return g.lights[x][y]
}

func (g *Grid) TurnOnSquare(x1, y1, x2, y2 int) {
	g.mapRectangle(x1, y1, x2, y2, func(light int) int {
		return light + 1
	})
}

func (g *Grid) TurnOffSquare(x1, y1, x2, y2 int) {
	g.mapRectangle(x1, y1, x2, y2, func(light int) int {
		if light == 0 {
			return light
		}
		return light - 1
	})
}

func (g *Grid) ToggleSquare(x1, y1, x2, y2 int) {
	g.mapRectangle(x1, y1, x2, y2, func(light int) int {
		return light + 2
	})
}

func (g *Grid) TotalBrightness() int {
	total := 0
	for i := 0; i < len(g.lights); i++ {
		for j := 0; j < len(g.lights[i]); j++ {
			total += g.lights[i][j]
		}
	}
	return total
}

func (g *Grid) mapRectangle(x1, y1, x2, y2 int, mapper func(brightness int) int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			g.lights[i][j] = mapper(g.lights[i][j])
		}
	}
}
