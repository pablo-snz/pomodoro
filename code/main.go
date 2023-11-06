package main

import (
	"code/client"
	"code/config_parser"
	"code/daemon"
	"code/pomodoro"
	"fmt"
	"github.com/spf13/cobra"
)

var sound bool
var startCmd = &cobra.Command{
	Use:     "start [\"STATE:TIME STATE:TIME ...\"]",
	Short:   "Start the Pomodoro timer",
	Example: "start \"work:25 break:5\"",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var pomodoro_states []pomodoro.PomodoroStates
		var err error

		if len(args) == 0 {
			pomodoro_states, err = config_parser.GetPomodoroStates()
			fmt.Println("Starting the Pomodoro timer with default settings:")
		} else {
			arg := args[0]
			pomodoro_states, err = config_parser.Parse(arg)
			fmt.Println("Starting the Pomodoro timer with custom settings:")
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		d := daemon.NewDaemon(pomodoro_states, sound)
		d.Start()

		for _, state := range pomodoro_states {
			fmt.Printf("%v: %v min\n", state.State, state.Time)
		}

	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the Pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
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
		client, err := client.NewPomodoroIPCClient()
		if err != nil {
			fmt.Println(err)
			return
		}
		response, err := client.SendQuery("status")
		if err != nil {
			fmt.Println(err)
			return
		}
		var output string

		for _, status := range response.States {
			output = fmt.Sprintf("State: %v, Time: %v", status.State, status.Time)
			if status.State == response.Current {
				output += fmt.Sprintf(" <- Current: %d min %d sec remaining\n", response.Minutes, response.Seconds)
			} else {
				output += "\n"
			}
			fmt.Print(output)
		}

	},
}

var setCmd = &cobra.Command{
	Use:   "set [STATUS]",
	Short: "Set the Pomodoro timer status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.NewPomodoroIPCClient()
		if err != nil {
			fmt.Println(err)
			return
		}
		response, err := client.SendCommand(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(response)
	},
}

func main() {
	var rootCmd = &cobra.Command{Use: "pomodoro"}
	startCmd.Flags().BoolVarP(&sound, "sound", "s", false, "Play a sound when the timer starts")
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(setCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
