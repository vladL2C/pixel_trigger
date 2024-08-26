package trigger

import (
	"fmt"
	"image"
	"time"

	"github.com/go-vgo/robotgo"
	gohook "github.com/robotn/gohook"
	"github.com/vladl2c/pixel_trigger/pkg/color"
)

type config struct {
	tolerence   uint32
	targetColor string
	triggerKey  int
	isKeyHeld   bool
}

type screen struct {
	rectWidth, rectHeight int
	startX, startY        int
}

type Colorer interface {
	GetColor(val string) (r, g, b uint32)
}

const (
	leftAltRawCode = 58
)

type trigger struct {
	captureScreen chan *image.Image
	colorDetected chan bool
	screen        screen
	config        *config
	color         Colorer
}

func Init() *trigger {
	// Set the size of the rectangle
	rectWidth, rectHeight := 5, 5
	screenWidth, screenHeight := robotgo.GetScreenSize()

	// Calculate the center of the screen
	centerX := (screenWidth / 2)
	centerY := (screenHeight / 2)

	// find top left corner from the middle of screen to draw square
	startX := centerX - rectWidth/2
	startY := centerY - rectHeight/2

	screen := screen{
		rectWidth:  rectWidth,
		rectHeight: rectHeight,
		startX:     startX,
		startY:     startY,
	}

	config := &config{
		tolerence:   80,
		targetColor: "red",
		triggerKey:  leftAltRawCode,
	}

	return &trigger{
		screen:        screen,
		captureScreen: make(chan *image.Image),
		colorDetected: make(chan bool),
		config:        config,
		color:         color.New(),
	}
}

func (t *trigger) Run() {
	go t.setKeyState()
	go t.CaptureScreen(t.screen.startX, t.screen.startY, t.screen.rectWidth, t.screen.rectHeight, t.captureScreen)
	go t.ScanImage(t.captureScreen, t.colorDetected)
	for isRed := range t.colorDetected {
		if isRed && t.config.isKeyHeld {
			start := time.Now()
			robotgo.Click()
			time.Sleep(14 * time.Millisecond)
			elapsed := time.Since(start)

			fmt.Println(fmt.Sprintf("elapsed time %v", elapsed))
		}
	}
}

func (t *trigger) CaptureScreen(startX, startY, rectWidth, rectHeight int, captureScreen chan *image.Image) {
	for {
		img := robotgo.CaptureImg(startX, startY, rectWidth, rectHeight)
		captureScreen <- &img
	}
}

func (t *trigger) ScanImage(screnshotCh chan *image.Image, hasRed chan bool) {

	for img := range screnshotCh {
		hasRed <- t.scanImage(img)
	}
}

// scanImageForRed scans the image for any red pixels
func (t *trigger) scanImage(img *image.Image) bool {
	// Get image bounds
	bounds := (*img).Bounds()

	// Iterate over each pixel in the image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the color of the pixel at (x, y)
			c := (*img).At(x, y)

			// Convert color to RGBA format
			r, g, b, _ := c.RGBA()

			// Normalize to 8-bit color values
			r8, g8, b8 := r>>8, g>>8, b>>8
			// Check if the pixel is "red"
			if t.isTarget(r8, g8, b8) {
				return true
			}
		}
	}

	// No red color found
	return false
}

// isRed checks if the given RGB values correspond to a red color
func (t *trigger) isTarget(r, g, b uint32) bool {
	// Define target color for red and tolerance
	targetR, targetG, targetB := t.color.GetColor(t.config.targetColor)
	tolerance := t.config.tolerence

	// Check if each color component is within tolerance of the target color
	rMatch := (r >= targetR-tolerance) && (r <= targetR+tolerance)
	gMatch := (g >= targetG-tolerance) && (g <= targetG+tolerance)
	bMatch := (b >= targetB-tolerance) && (b <= targetB+tolerance)

	// Return true if the color matches the red color within tolerance
	return rMatch && !gMatch && !bMatch
}

func (t *trigger) setKeyState() {
	eventHook := gohook.Start()
	var e gohook.Event

	for e = range eventHook {
		switch e.Kind {
		case gohook.KeyHold, gohook.KeyDown:
			if e.Rawcode == leftAltRawCode {
				t.config.isKeyHeld = true
			}
		case gohook.KeyUp:
			if e.Rawcode == leftAltRawCode {
				t.config.isKeyHeld = false
			}
		}

	}
}
