package DebugConsole

import (
	"fmt"
	"net"
)

type Connection struct {
	Connection net.Conn
	IsConnected bool
}

type DebugConsole struct {
	Network string
	Address string
	Connection
}

func (console *DebugConsole) PrintToDebugConsole(data string) {
	if !console.Connection.IsConnected {
		fmt.Println("Debug console not connected!")
		return
	}
	_, err := console.Connection.Connection.Write([]byte(data))
	if err != nil {
		fmt.Println("Cannot write to debug console!")
		return
	}
}

func (console *DebugConsole) InitConnection() {
	if len(console.Network) == 0 {
		panic("Please provide a network type for debug console! (tcp/udp)")
	}
	if len(console.Address) == 0 {
		panic("Please provide an address for debug console!")
	}

	conn, err := net.Dial(console.Network, console.Address)
	if err != nil {
		panic("Connection to debug console failed!")
	}
	console.Connection.Connection = conn
	console.Connection.IsConnected = true
}

func (console *DebugConsole) CloseConnection() {
	if !console.Connection.IsConnected {
		fmt.Println("Debug console not connected!")
		return
	}
	err := console.Connection.Connection.Close()
	if err != nil {
		return
	}
}