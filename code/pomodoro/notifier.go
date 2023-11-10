package pomodoro

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

type Notifier struct {
	message bool
	sound   bool
	script  string
	title   string
}

func NewNotifier(sound bool, message bool) *Notifier {
	return &Notifier{
		message: message,
		sound:   sound,
		script:  "notify-send",
		title:   "Pomodoro",
	}
}

func (n *Notifier) notify(order int, status string) {
	if n.message {
		message := fmt.Sprintf("%s Time has started", status)
		cmd := exec.Command(n.script, n.title, message)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error executing notify-send:", err)
		}
	}
	if n.sound {
		n.playSound()
	}

}

func (n *Notifier) playSound() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(homeDir + "/.pomodoro/assets/sound.mp3")
	if err != nil {
		return
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
