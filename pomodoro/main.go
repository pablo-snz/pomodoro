package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"pomodoro/client"
	"pomodoro/daemon"
)

func main() {
	app := cli.NewApp()
	app.Name = "My Daemon Control"
	app.Usage = "Control my daemon"
	app.Version = "1.0.0"

	daemonInstance := daemon.NewDaemon()
	clientInstance := client.NewClient(daemonInstance)

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the daemon",
			Action: func(c *cli.Context) error {
				clientInstance.StartDaemon()
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "Stop the daemon",
			Action: func(c *cli.Context) error {
				clientInstance.StopDaemon()
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
	}
}
