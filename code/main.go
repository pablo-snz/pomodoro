package main

import (
	"code/client"
	"code/config_parser"
	"code/daemon"
	"code/pomodoro"
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Pomodoro timer",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var pomodoro_states []pomodoro.PomodoroStates
		var err error

		if len(args) == 0 {
			fmt.Println("Starting the Pomodoro timer with default settings...")
			pomodoro_states, err = config_parser.GetPomodoroStates()
		} else {
			arg := args[0]
			fmt.Printf("Starting the Pomodoro timer with custom settings: %v\n", arg)
			pomodoro_states, err = config_parser.Parse(arg)
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		d := daemon.NewDaemon(pomodoro_states)
		d.Start()

	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the Pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping the Pomodoro timer...")
		client, err := client.NewPomodoroIPCClient()
		if err != nil {
			fmt.Println(err)
			return
		}
		response, err := client.SendCommand("stop")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(response)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the Pomodoro timer status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting the Pomodoro timer status...")
		client, err := client.NewPomodoroIPCClient()
		if err != nil {
			fmt.Println(err)
			return
		}
		response, err := client.SendCommand("status")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Status: ", response)
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "pomodoro"}
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(statusCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
