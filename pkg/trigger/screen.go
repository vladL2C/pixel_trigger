package trigger

import "github.com/go-vgo/robotgo"

type screen struct {
	rectWidth, rectHeight int
	startX, startY        int
}

func NewScreen() screen {
	screenWidth, screenHeight := robotgo.GetScreenSize()

	// Calculate the center of the screen
	centerX := (screenWidth / 2)
	centerY := (screenHeight / 2)

	// find top left corner from the middle of screen to draw square
	startX := centerX - rectWidth/2
	startY := centerY - rectHeight/2

	return screen{
		rectWidth:  rectWidth,
		rectHeight: rectHeight,
		startX:     startX,
		startY:     startY,
	}
}
