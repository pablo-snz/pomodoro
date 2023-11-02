package config_parser

import (
	"code/pomodoro"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Pomodoro []pomodoro.PomodoroStates `yaml:"pomodoro"`
}

func GetPomodoroStates() ([]pomodoro.PomodoroStates, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".pomodoro")
	configFilePath := filepath.Join(configDir, "config.yml")

	pomodoroConfig, err := readPomodoroConfig(configFilePath)
	if err != nil {
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			err := os.MkdirAll(configDir, os.ModePerm)
			if err != nil {
				return nil, err
			}
		}
		pomodoroConfig = []pomodoro.PomodoroStates{
			{Order: 0, State: "work", Time: 25},
			{Order: 1, State: "rest", Time: 5},
		}

		err := writePomodoroConfig(configFilePath, pomodoroConfig)
		if err != nil {
			return pomodoroConfig, err
		}
	}

	return pomodoroConfig, nil
}

func readPomodoroConfig(filename string) ([]pomodoro.PomodoroStates, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config.Pomodoro, nil
}

func writePomodoroConfig(filename string, pomodoroConfig []pomodoro.PomodoroStates) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var config map[string]interface{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	config["pomodoro"] = pomodoroConfig

	updatedData, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, updatedData, 0644); err != nil {
		return err
	}

	return nil
}

func Parse(input string) ([]pomodoro.PomodoroStates, error) {
	parts := strings.Split(input, " ")
	states := make([]pomodoro.PomodoroStates, 0)
	idx := 0

	for _, part := range parts {
		parts := strings.Split(part, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid format: %s", part)
		}

		state := parts[0]
		time, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, err
		}
		states = append(states, pomodoro.PomodoroStates{Order: idx, State: state, Time: time})
		idx++
	}

	return states, nil
}
