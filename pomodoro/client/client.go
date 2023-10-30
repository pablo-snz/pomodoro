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

func (c *Client) StartDaemon() {
	c.daemon.Start()
}

func (c *Client) StopDaemon() {
	c.daemon.Stop()
}
