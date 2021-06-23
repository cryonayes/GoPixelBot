package main

import (
	"GoPixelBot/Arduino"
	"GoPixelBot/DebugConsole"
	"GoPixelBot/ScreenAnalyzer"
	"image"
)

func main() {

	var arduino = Arduino.Arduino{}
	// Port value can be different on your system
	arduino.ConnectArduino("COM5", 9600)
	defer arduino.CloseConnection()

	var debugConsole = DebugConsole.DebugConsole{}
	/*
	// I used this to debug my code with my smartphone while gaming. nc -lp 8888
	debugConsole = DebugConsole.DebugConsole{Network: "tcp", Address: "192.168.1.10:8888"}
	debugConsole.InitConnection()
	defer debugConsole.CloseConnection()
	*/

	// 4 Pixel from center of the screen, my screen is 1080p
	bounds := image.Rect(958, 538, 962, 542)
	sa := ScreenAnalyzer.ScreenAnalyzer{}
	sa.StartPixelBot(bounds, debugConsole, &arduino)

}
