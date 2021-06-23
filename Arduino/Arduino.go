package Arduino

import (
	serial "github.com/tarm/goserial"
	"io"
)

type Arduino struct {
	Port string
	BaudRate int
	Serial io.ReadWriteCloser
}

var IsConnected bool

func (ar *Arduino) ConnectArduino(port string, baudRate int) {
	ar.Port = port
	ar.BaudRate = baudRate

	config := &serial.Config{
		Name: port,
		Baud: baudRate,
	}

	serialArduino, err := serial.OpenPort(config)
	if err != nil {
		panic("Cannot connect to Arduino!")
	}
	IsConnected = true
	ar.Serial = serialArduino
}

func (ar *Arduino) SendString(data string) {
	if !IsConnected {
		panic("Cannot connect to Arduino")
	}
	_, err := ar.Serial.Write([]byte(data))
	if err != nil {
		panic("Cannot send data to Arduino")
	}
}

func (ar *Arduino) SendBytes(data []byte) {
	if !IsConnected {
		panic("Cannot connect to Arduino")
	}
	_, err := ar.Serial.Write(data)
	if err != nil {
		panic("Cannot send data to Arduino")
	}
}

func (ar *Arduino) CloseConnection() {
	_ = ar.Serial.Close()
}