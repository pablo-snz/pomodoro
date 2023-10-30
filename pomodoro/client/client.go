package client

import (
	"pomodoro/daemon"
)

type Client struct {
	daemon *daemon.Daemon
}

func NewClient(d *daemon.Daemon) *Client {
	return &Client{
		daemon: d,
	}
}

func (c *Client) StartDaemon(workTime, breakTime float64) {
	c.daemon.Start(workTime, breakTime)
}

func (c *Client) StopDaemon() {
	c.daemon.Stop()
}
