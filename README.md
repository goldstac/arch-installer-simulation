# Arch Installer Simulation

A minimalist terminal-based Arch Linux installation simulator. Provides a mock shell environment to practice basic Arch-like commands in a simulated, read-only session.

## Usage

```bash
python3 src/core/main.py
```

No external dependencies required — pure Python 3 standard library.

## Commands

| Command | Description |
|---|---|
| `exit` | Exit the simulator |
| `clear` | Clear the terminal |
| `whoami` | Display the simulated username |
| `sudo pacman -Syu` | Simulate an Arch system update |
| `run-app` | Launch an external application |

## License

MIT — Copyright (c) 2026 Li Productions.
