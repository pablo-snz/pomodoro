package daemon

import (
	"bytes"
	"code/pomodoro"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type PomodoroStatusResponse struct {
	Current string
	States  []pomodoro.PomodoroStates
	Minutes int
	Seconds int
}

type PomodoroIPCServer struct {
	message  bool
	sound    bool
	pomodoro *pomodoro.Pomodoro
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewPomodoroIPCServer(pomodoroStates []pomodoro.PomodoroStates, sound bool, message bool) *PomodoroIPCServer {
	ctx, cancel := context.WithCancel(context.Background())
	pom := pomodoro.NewPomodoro(pomodoroStates, sound, message)
	return &PomodoroIPCServer{
		message:  message,
		sound:    sound,
		pomodoro: pom,
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (s *PomodoroIPCServer) startPomodoro() {
	go s.pomodoro.Start(s.ctx)
}

func (s *PomodoroIPCServer) Start() {

	listenPath := "/tmp/pomodoro.sock"
	os.Remove(listenPath)

	l, err := net.Listen("unix", listenPath)
	if err != nil {
		fmt.Printf("Error al escuchar en el socket Unix: %v\n", err)
		return
	}
	defer l.Close()

	s.startPomodoro()

	fmt.Println("Servidor IPC Pomodoro iniciado")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		s.cancel()
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("Error al aceptar la conexión: %v\n", err)
			return
		}
		go s.handleConnection(conn)
	}
}

func (s *PomodoroIPCServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Error al leer desde la conexión: %v\n", err)
		return
	}

	request := strings.TrimSpace(string(bytes.Trim(buf, "\x00")))

	switch request {
	case "status":
		status := s.pomodoro.GetStatus()
		states := s.pomodoro.GetStates()
		minutes, seconds := s.pomodoro.GetTime()

		response := PomodoroStatusResponse{
			Current: status,
			States:  states,
			Minutes: minutes,
			Seconds: seconds,
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("Error al serializar la respuesta: %v\n", err)
			return
		}

		conn.Write(responseJSON)

	case "stop":
		conn.Write([]byte("Deteniendo el servidor"))
		conn.Close()
		os.Exit(0)
	default:
		states := s.pomodoro.GetStates()

		order, found := getOrderForState(states, request)
		if !found {
			conn.Write([]byte("Comando no válido"))
			return
		}

		newStates := make([]pomodoro.PomodoroStates, len(states))
		newStates[0] = pomodoro.PomodoroStates{
			State: request,
			Time:  states[order].Time,
			Order: 0,
		}

		for i := 1; i < len(states); i++ {
			newOrder := (i + order) % len(states)
			newStates[i] = pomodoro.PomodoroStates{
				State: states[newOrder].State,
				Time:  states[newOrder].Time,
				Order: i,
			}
		}

		s.cancel()
		ctx, cancel := context.WithCancel(context.Background())
		pom := pomodoro.NewPomodoro(newStates, s.sound, s.message)
		s.cancel = cancel
		s.ctx = ctx
		s.pomodoro = pom
		s.startPomodoro()

		conn.Write([]byte("Pomodoro reiniciado, estado actual: " + request))
	}
}

func getOrderForState(states []pomodoro.PomodoroStates, targetState string) (int, bool) {
	for _, state := range states {
		if state.State == targetState {
			return state.Order, true
		}
	}
	return 0, false
}
