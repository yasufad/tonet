# tonet 🐊

> Easily and securely send things from one computer to another - now with a native desktop UI.

**tonet** is a fork of [croc](https://github.com/schollz/croc) by schollz that adds a cross-platform native GUI built with [Fyne](https://fyne.io). Everything under the hood is unchanged - the same PAKE key exchange, the same relay infrastructure, the same multiplexed TCP transfers. tonet just gives it a proper window.

---

## Features

- All croc features: end-to-end encrypted transfers, relay-assisted or local, multiplexed ports, resumable transfers
- Native cross-platform GUI (Windows, macOS, Linux) via Fyne v2
- Drag-and-drop file/folder sending
- QR code display for mobile receive
- Built-in relay management panel
- Persistent settings (relay address, encryption curve, proxy)
- **Full CLI backward compatibility** - `tonet send file.txt` works exactly as before

---

## Installation

```bash
# From source (requires Go 1.24+)
git clone https://github.com/yasufad/tonet
cd tonet
go build -o tonet .

# Or install directly
go install github.com/yasufad/tonet@latest
```

### Platform dependencies for Fyne

| Platform | Requirement |
|---|---|
| Linux | `libgl1-mesa-dev xorg-dev` |
| macOS | Xcode Command Line Tools |
| Windows | No extra deps (uses DirectX) |

---

## Usage

### GUI mode

```bash
tonet
```

Launches the desktop UI. No arguments needed.

### CLI mode (identical to croc)

```bash
# Send
tonet send file.txt
CROC_SECRET=my-code tonet send file.txt

# Receive
tonet my-code

# Run your own relay
tonet relay --ports 9009,9010,9011,9012,9013
```

All original croc flags and environment variables are supported unchanged.

---

## Architecture

```
tonet/
├── main.go                  # Entry: UI if no args, CLI otherwise
├── src/
│   ├── cli/
│   │   └── cli.go           # Original croc CLI (unmodified)
│   ├── ui/
│   │   ├── app.go           # Fyne app bootstrap, tab container
│   │   ├── send.go          # Send tab
│   │   ├── receive.go       # Receive tab
│   │   ├── relay.go         # Relay management tab
│   │   ├── settings.go      # Settings tab (reads/writes croc config JSON)
│   │   ├── theme.go         # tonet Fyne theme
│   │   └── progress.go      # Goroutine bridge: croc state → Fyne widgets
│   ├── comm/                # (croc, unmodified)
│   ├── croc/                # (croc, unmodified)
│   ├── mnemonicode/         # (croc, unmodified)
│   ├── models/              # (croc, unmodified)
│   ├── tcp/                 # (croc, unmodified)
│   └── utils/               # (croc, unmodified)
├── assets/
│   ├── icon.png
│   └── icon.icns
├── go.mod
├── go.sum
└── README.md
```

See [ARCHITECTURE.md](./ARCHITECTURE.md) for the full design document.

---

## Building for distribution

```bash
# Install fyne CLI tool
go install fyne.io/fyne/v2/cmd/fyne@latest

# Package for current platform
fyne package -os linux -icon assets/icon.png
fyne package -os darwin -icon assets/icon.icns
fyne package -os windows -icon assets/icon.png
```

---

## Credits

- [croc](https://github.com/schollz/croc) by [@schollz](https://github.com/schollz) - all transfer logic
- [Fyne](https://fyne.io) - UI toolkit
- [pake](https://github.com/schollz/pake) - password-authenticated key exchange

## Licence

MIT - same as croc.