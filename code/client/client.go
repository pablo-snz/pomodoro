package client

import (
	"net"
)

type PomodoroIPCClient struct {
	serverSocket string
	conn         net.Conn
}

func NewPomodoroIPCClient() (*PomodoroIPCClient, error) {
	serverSocket := "/tmp/pomodoro.sock"

	conn, err := net.Dial("unix", serverSocket)
	if err != nil {
		return nil, err
	}

	return &PomodoroIPCClient{
		serverSocket: serverSocket,
		conn:         conn,
	}, nil
}

func (c *PomodoroIPCClient) SendCommand(command string) (string, error) {
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		return "", err
	}

	response := make([]byte, 1024)
	_, err = c.conn.Read(response)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

func (c *PomodoroIPCClient) Close() {
	c.conn.Close()
}
