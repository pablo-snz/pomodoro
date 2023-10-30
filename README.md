# pomodoro

This project is a productivity tool designed for Linux users who work with tiling window managers such as `i3wm` and `sway`. It's intended to help you manage your work and break times using the Pomodoro technique, a time management method that breaks work into intervals of focused work followed by short breaks.

**Under Development**

## Features

- It serves as a daemon for managing your work and break sessions.
- Easily customizable work and break times.
- Designed for Linux users working with tiling window managers

## Installation

Before compiling and using Pomodoro Daemon, make sure you have the necessary dependencies and Go environment set up.

1. Build the project using the following command:

```bash
go build -o pomodoro main.go
```

You're now ready to use Pomodoro Daemon.

## Usage

### Starting Daemon

To start the daemon, use the following command:

```{bash}
./pomodoro start --work [work_minutes] --break [break_minutes]
```

Replace `[work_minutes]` with the duration of your work interval in minutes and `[break_minutes]` with the duration of your break interval in minutes.

### Stopping the Daemon

To stop the daemon, use the following command:

```{bash}
./pomodoro stop
```

### Development Status

Pomodoro Daemon is currently under development and may have some limitations or missing features. We welcome contributions 

