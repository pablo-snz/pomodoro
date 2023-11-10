package pomodoro

import (
	"context"
	"sort"
	"time"
)

type PomodoroStates struct {
	Order int
	State string
	Time  float64
}

type Pomodoro struct {
	states      []PomodoroStates
	notifier    *Notifier
	status      string
	currentTime time.Time
	time        float64
}

func NewPomodoro(pomodoroStates []PomodoroStates, sound bool, message bool) *Pomodoro {

	sort.Slice(pomodoroStates, func(i, j int) bool {
		return pomodoroStates[i].Order < pomodoroStates[j].Order
	})

	return &Pomodoro{
		notifier: NewNotifier(sound, message),
		states:   pomodoroStates,
	}
}

func (p *Pomodoro) Start(ctx context.Context) {
	idx := 0

	len_states := len(p.states)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			state := p.states[idx]
			p.status = state.State
			p.time = state.Time
			go p.notifier.notify(state.Order, p.status)

			stateDuration := time.Duration(state.Time*60) * time.Second
			timer := time.NewTimer(stateDuration)
			p.currentTime = time.Now().Add(stateDuration)

			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				timer.Stop()
			}

			idx++
			if idx >= len_states {
				idx = 0
			}
		}
	}
}

func (p *Pomodoro) GetTime() (int, int) {
	duration := p.currentTime.Sub(time.Now())
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) - (minutes * 60)

	return minutes, seconds
}

func (p *Pomodoro) GetStatus() string {
	return p.status
}

func (p *Pomodoro) GetStates() []PomodoroStates {
	return p.states
}
