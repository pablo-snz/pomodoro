package daemon

import (
	"code/pomodoro"
	"fmt"
	"os"

	"github.com/sevlyar/go-daemon"
)

type Daemon struct {
	server *PomodoroIPCServer
}

func NewDaemon(pomodoroStates []pomodoro.PomodoroStates) *Daemon {
	
	server := NewPomodoroIPCServer(pomodoroStates)

	return &Daemon{
		server: server,
	}
}

func (d *Daemon) Start(workTime, breakTime float64) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pomodoroDir := homeDir + "/.pomodoro/"
	if _, err := os.Stat(pomodoroDir); os.IsNotExist(err) {
		err := os.Mkdir(pomodoroDir, 0755)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	cntxt := &daemon.Context{
		PidFileName: "pomodoro.pid",
		PidFilePerm: 0644,
		LogFileName: "pomodoro.log",
		LogFilePerm: 0640,
		WorkDir:     pomodoroDir,
		Umask:       027,
	}

	demon, err := cntxt.Reborn()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if demon != nil {
		fmt.Println("Daemon started with PID", demon.Pid)
		return
	}
	defer cntxt.Release()

	d.server.Start()
}


