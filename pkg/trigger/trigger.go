package trigger

import (
	"image"

	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vladl2c/pixel_trigger/pkg/color"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	VK_XBUTTON1          = 0x05
)

type Config struct {
	Tolerence   uint32
	TargetColor string
	TriggerKey  int
	IsKeyHeld   bool
}

type Colorer interface {
	GetColor(val string) (r, g, b uint32)
}

const (
	// alt key
	leftAltRawCode = 88
	// screenshot size
	rectWidth  = 5
	rectHeight = 5
)

type trigger struct {
	captureScreen chan *image.Image
	colorDetected chan bool
	screen        screen
	config        *Config
	color         Colorer
}

func Init(config *Config) *trigger {
	screen := NewScreen()

	if config == nil {
		config = GenerateDefaultConfig()
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
	// start routines to do processing
	go t.screenshot()
	go t.scanImage()

	for isDetected := range t.colorDetected {
		if isDetected && GetAsyncKeyState(VK_XBUTTON1) {
			robotgo.KeyPress("k")
		}
	}
}

func (t *trigger) screenshot() {
	for {
		img := robotgo.CaptureImg(t.screen.startX, t.screen.startY, rectWidth, rectHeight)
		t.captureScreen <- &img
		time.Sleep(8 * time.Millisecond)
	}
}

func (t *trigger) scanImage() {
	for img := range t.captureScreen {
		t.colorDetected <- t.detectTargetColor(img)
	}
}

// scanImage scans the image checks target
func (t *trigger) detectTargetColor(img *image.Image) bool {
	bounds := (*img).Bounds()

	// Iterate over each pixel in the image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := (*img).At(x, y)
			r, g, b, _ := c.RGBA()
			// Normalize to 8-bit color values
			r8, g8, b8 := r>>8, g>>8, b>>8

			if t.isTarget(r8, g8, b8) {
				return true
			}
		}
	}

	// No color found
	return false
}

func (t *trigger) isTarget(r, g, b uint32) bool {
	// Define target color and tolerance
	targetR, targetG, targetB := t.color.GetColor(t.config.TargetColor)
	tolerance := t.config.Tolerence

	// Check if each color component is within tolerance of the target color
	rMatch := (r >= targetR-tolerance) && (r <= targetR+tolerance)
	gMatch := (g >= targetG-tolerance) && (g <= targetG+tolerance)
	bMatch := (b >= targetB-tolerance) && (b <= targetB+tolerance)

	// Return true if the color matches the color within tolerance
	return rMatch && !gMatch && !bMatch
}

func GetAsyncKeyState(vk int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vk))
	return ret != 0
}
