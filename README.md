# 2048-tui

A terminal based implementation of the classic **2048** puzzle game, written in Go.

Move the tiles, combine them to create larger numbers, and try to reach 2048, all within the command line!

## Demo
<img width="2821" height="1696" alt="Screenshot_20260721_180847" src="https://github.com/user-attachments/assets/1a86c1eb-b41f-4ffe-baae-43df9e5738db" />

## Installation

Grab a pre-built binary in the [releases page](https://github.com/pinwheeeel/2048-tui/releases)

If you are on Windows, just run the `.exe` file directly in Command Prompt or PowerShell

Otherwise, make sure to give the binary execution permissions:

```bash
mv 2048-tui-linux-amd64 2048-tui
chmod +x 2048-tui
```

You will need to add the binary to your PATH to run it from any directory in the command line.

```bash
sudo mv 2048-tui /usr/bin/2048-tui
```

### Building from source

Make sure you have [Go 1.26](https://go.dev/doc/install) installed.

```bash
git clone https://github.com/pinwheeeel/2048-tui.git
cd 2048-tui
go install  # alternatively, use `go build` to build without installing
```
