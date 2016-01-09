package main

type LightGrid interface {
	TurnOnSquare(x1, y1, x2, y2 int)
	TurnOffSquare(x1, y1, x2, y2 int)
	ToggleSquare(x1, y1, x2, y2 int)
	TotalBrightness() int
}
