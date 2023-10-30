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
	app.Name = "Pomodoro"
	app.Usage = "Pomodoro daemon & client"
	app.Version = "1.0.0"

	daemonInstance := daemon.NewDaemon()
	clientInstance := client.NewClient(daemonInstance)

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the daemon",
			Flags: []cli.Flag{
				cli.Float64Flag{
					Name:  "work, w",
					Usage: "Set work time in minutes (decimals allowed)",
				},
				cli.Float64Flag{
					Name:  "break, b",
					Usage: "Set break time in minutes (decimals allowed)",
				},
			},
			Action: func(c *cli.Context) error {
				workTime := c.Float64("work")
				breakTime := c.Float64("break")

				if workTime <= 0 || breakTime <= 0 {
					fmt.Println("Work and break times must be greater than 0")
					return nil
				}

				clientInstance.StartDaemon(workTime, breakTime)
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
