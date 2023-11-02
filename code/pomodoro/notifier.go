package pomodoro

import (
	"fmt"
	"os"
	"os/exec"
)

type Notifier struct {
	script string
	title  string
}

func NewNotifier() *Notifier {
	return &Notifier{
		script: "notify-send",
		title:  "Pomodoro",
	}
}

func (n *Notifier) notify(order int, status string) {
	message := fmt.Sprintf("%s Time has started", status)
	cmd := exec.Command(n.script, n.title, message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing notify-send:", err)
	}
}
