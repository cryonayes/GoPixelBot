package ScreenAnalyzer

import (
	"GoPixelBot/Arduino"
	"GoPixelBot/DebugConsole"
	"fmt"
	"image"
	"time"

	"github.com/kbinani/screenshot"
)

type ScreenAnalyzer struct{}

var clickChan = make(chan bool)
var enemyCount = 0
var isRunning = false
var isDebugEnabled = false

func (sa *ScreenAnalyzer) StartPixelBot(bounds image.Rectangle, debugConsole DebugConsole.DebugConsole, arduino *Arduino.Arduino) {
	if (DebugConsole.DebugConsole{}) != debugConsole {
		isDebugEnabled = true
	}
	isRunning = true

	go func() {
		for {
			select {
			case <-clickChan:
				if isDebugEnabled {
					enemyCount++
					debugConsole.PrintToDebugConsole(fmt.Sprintf("Enemy detected %d\n", enemyCount))
				}
				// Arduino should listen for serial data, each char means left click
				arduino.SendString("F")
			}
		}
	}()

	for isRunning {
		ss := sa.takeScreenShot(bounds)
		sa.analyzeScreenshot(ss)
	}

}

func (sa *ScreenAnalyzer) StopPixelBot() {
	isRunning = false
}

func (sa *ScreenAnalyzer) takeScreenShot(bounds image.Rectangle) *image.RGBA {
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic("Error while capturing screenshot!")
	}
	return img
}

func (sa *ScreenAnalyzer) analyzeScreenshot(screenShot *image.RGBA) {
	for i := 0; i < screenShot.Rect.Max.X; i++ {
		for j := 0; j < screenShot.Rect.Max.Y; j++ {
			r, g, b, _ := screenShot.At(i, j).RGBA()
			r = r / 257
			g = g / 257
			b = b / 257

			// These conditions should be changed for your usage !
			if r > 255 && g > 255 && g < 255 && b > 255 {
				clickChan <- true
				time.Sleep(time.Second / 60)
				return
			}

			if r > 255 && r < 255 && g > 255 && g < 255 && b > 255 && b < 255 {
				clickChan <- true
				time.Sleep(time.Second / 60)
				return
			}

			if r > 255 && r < 255 && g > 255 && g < 255 && b > 255 && b < 255 {
				clickChan <- true
				time.Sleep(time.Second / 60)
				return
			}
		}
	}
}
