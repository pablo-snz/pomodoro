package daemon

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

type Daemon struct {
	PIDFile string
}

func NewDaemon() *Daemon {
	return &Daemon{
		PIDFile: "daemon-example.pid",
	}
}

func (d *Daemon) Start(workTime, breakTime float64) {
	cntxt := &daemon.Context{
		PidFileName: d.PIDFile,
		PidFilePerm: 0644,
		LogFileName: "daemon-example.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
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

	d.runPomodoro(workTime, breakTime)
}

func (d *Daemon) runPomodoro(workTime, breakTime float64) {
	
	workTimeSeconds := workTime * 60
	breakTimeSeconds := breakTime * 60

	for {

		d.notify("Pomodoro", "Work Time has started")

		time.Sleep(time.Duration(workTimeSeconds) * time.Second)

		d.notify("Pomodoro", "Break Time has started")

		time.Sleep(time.Duration(breakTimeSeconds) * time.Second)
	}
}

func (d *Daemon) notify(title, message string) {
	cmd := exec.Command("notify-send", title, message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing notify-send:", err)
	}
}

func (d *Daemon) Stop() {
	pidFile, err := os.ReadFile(d.PIDFile)
	if err != nil {
		fmt.Println("Error reading PID file:", err)
		return
	}

	pid, err := strconv.Atoi(string(pidFile))
	if err != nil {
		fmt.Println("Error parsing PID:", err)
		return
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Error finding process:", err)
		return
	}

	if err := process.Signal(syscall.SIGTERM); err != nil {
		fmt.Println("Error stopping daemon:", err)
	}
}
