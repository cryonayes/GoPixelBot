# GoPixelBot
### Arduino assisted auto clicker based on pixel detection.

## How it works ?
The HID compatible atmega32u4 microchip used in this project can send mouse/keyboard inputs to the computer.
When a certain pixel color is detected on the screen, this program will send a signal to the arduino over the serial connection.
At the receiving end, the arduino will receive these signals and make clicks to fire in game.
