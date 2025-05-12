# Auto Key Presser bot (Shift Key)

This Go program simulates pressing the **Shift** key at a fixed interval using the [`robotgo`](https://github.com/go-vgo/robotgo) library. It can be used to prevent system sleep, keep sessions active, or for other automation tasks.

## Features

- Presses the `Shift` key every 200 seconds
- Logs each key press with the current timestamp

## Requirements

- Go 1.24.3
- CGO support (required by `robotgo`)
- `gcc` or equivalent C compiler installed on your system

## Installation

1. **Clone the repository** (or copy the code into a `.go` file):

   ```bash
   git clone https://github.com/SaptakBN/bot.git
   cd bot

