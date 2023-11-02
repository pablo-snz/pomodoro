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
	states   []PomodoroStates
	notifier *Notifier
	status   string
}

func NewPomodoro(pomodoroStates []PomodoroStates) *Pomodoro {

	sort.Slice(pomodoroStates, func(i, j int) bool {
		return pomodoroStates[i].Order < pomodoroStates[j].Order
	})

	return &Pomodoro{
		notifier: NewNotifier(),
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
			p.notifier.notify(state.Order, p.status)

			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(state.Time*60) * time.Second):
			}

			idx++
			if idx >= len_states {
				idx = 0
			}
		}
	}
}

func (p *Pomodoro) GetStatus() string {
	return p.status
}

func (p *Pomodoro) GetStates() []PomodoroStates {
	return p.states
}
