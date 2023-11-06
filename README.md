# pomodoro-daemon

pomodoro-daemon is a timer daemon designed for Unix-based systems, allowing you to effectively manage your time using the Pomodoro Technique. This daemon runs in the background, enabling you to concentrate on your tasks. 

## Features:
- Custom pomodoro timer: Start and manage Pomodoro timers with customized settings. 
- Sound notification: using [oto](https://github.com/ebitengine/oto) and [beep](https://github.com/gopxl/beep) you can configure sound notifications to accompany your Pomodoro sessions.
- Monitor the current Pomodoro timer status and make on-the-fly adjustments to your Pomodoro settings.

## Install:

pomodoro-daemon is tailored for Unix-based systems and relies on the presence of notify-send for notifications. Ensure you have Go version 1.21 or higher installed.

1. Clone the repository:

```{bash}
git clone https://github.com/pablo-snz/pomodoro-daemon.git
cd pomodoro-daemon
```

2. Build the executable

```{bash}
make build
```

3. (OPTIONAL) move the executable to PATH:

```{bash}
make install
make clean
```

## Uninstall:

To remove the binary from your PATH, you can execute:

```{bash}
make uninstall
```

Additionally, pomodoro-daemon creates a .pomodoro folder in your home directory. You can delete it with:

```{bash}
rm -r $HOME/.pomodoro/
```

## Usage

pomodoro-daemon offers versatile Pomodoro timer functionality. You can define as many states in the format **state:time** as you wish. You can do this in the configuration file `$HOME/.pomodoro/config.yml` as shown below:

```{yaml}
pomodoro:
- order: 0
  state: work
  time: 30
- order: 1
  state: rest
  time: 5
- order: 2
  state: Look-Away
  time: 5
# ...
```
Alternatively, you can define these states on-the-fly using the `pomodoro start` command:

```{yaml}
pomodoro start "work:25 break:5 ..."
```

- `start`: Start the Pomodoro timer with custom/default settings.
    - Flags:
        - `-s, --sound`: Play a sound when the timer starts.

- `status`: Get the Pomodoro timer's current status and time remaining.

- `stop`: Gracefully stop the Pomodoro timer.

- `set`: Change the pomodoro current status with another:

    ```{bash}
    pomodoro set Look-Away
    ```

## Limitations

- Currently, pomodoro-daemon is designed for Unix-based systems due to its use of Unix socket communication (IPC).
- The manager for notifications relies on notify-send, which should be installed and properly configured to display notifications.
- Only one pomodoro at a time. (working on it)

