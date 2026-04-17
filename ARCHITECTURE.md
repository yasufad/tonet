# tonet - Architecture

## Guiding Principle

**Touch nothing below `cli.go`.** The croc core (`src/croc`, `src/tcp`, `src/comm`, etc.) is battle-tested. tonet's entire contribution is a `src/ui` package plus an updated `main.go`. The CLI remains fully functional.

---

## Entry Point

```go
// main.go
func main() {
    if len(os.Args) > 1 {
        // Delegate to original croc CLI unchanged
        if err := cli.Run(); err != nil {
            log.Fatal(err)
        }
    } else {
        // Boot Fyne GUI
        ui.Run()
    }
}
```

This means:
- `tonet` → opens window
- `tonet send file.txt` → behaves exactly like `croc send file.txt`
- All shell scripts and CI pipelines using croc syntax continue to work

---

## Package: `src/ui`

### `app.go` - Bootstrap

Initialises the Fyne application and constructs the main window with a tab container:

```
┌─────────────────────────────────────────┐
│  tonet                          [─][□][×]│
├──────┬─────────┬────────┬───────────────┤
│ Send │ Receive │  Relay │  Settings     │
├──────┴─────────┴────────┴───────────────┤
│                                         │
│           [active tab content]          │
│                                         │
└─────────────────────────────────────────┘
```

Key responsibilities:
- Create `fyne.App` and `fyne.Window`
- Load persisted settings from croc's existing config JSON files
- Pass a shared `*AppState` to each tab constructor

### `send.go` - Send Tab

UI elements:
- File/folder picker (`dialog.ShowFileOpen`, `dialog.ShowFolderOpen`)
- Drag-and-drop target (`window.SetOnDropped`)
- Optional code override field (hidden unless `--classic` mode detected)
- Hash algorithm selector (xxhash / imohash / md5)
- Options: zip folder, no-compress, git-ignore, QR code toggle
- **Send** button → calls `startSend()` in a goroutine
- Generated code display: large monospace label + copy button
- QR code canvas element (rendered via `fyne/canvas.Image`)
- Progress bar + transfer rate label
- Cancel button

Flow:
```
[User picks files] → [Clicks Send]
    → goroutine: croc.New(options) → cr.Send(fileInfos, ...)
    → progress.Bridge polls state → updates Fyne widgets via fyne.Do()
```

### `receive.go` - Receive Tab

UI elements:
- Code input field with mnemonicode hint text
- Output directory picker
- Options: overwrite, stdout redirect
- **Receive** button → calls `startReceive()` in a goroutine
- Progress bar + file list (populates as files land)
- Reveal in Finder/Explorer button on completion

Flow:
```
[User enters code] → [Clicks Receive]
    → goroutine: croc.New(options) → cr.Receive()
    → progress.Bridge → Fyne widgets
```

### `relay.go` - Relay Tab

UI elements:
- Host input (optional)
- Base port input (default 9009)
- Transfer count spinner (default 4)
- **Start Relay** / **Stop Relay** toggle button
- Connection log (scrollable `widget.List`)
- Port status indicators

The relay goroutine runs `tcp.Run(...)` for each port. A context with cancel is used to stop it cleanly.

### `settings.go` - Settings Tab

Mirrors the global flags from `cli.go`:

| Setting | Widget | Stored in |
|---|---|---|
| Relay address | Entry | `send.json` / `receive.json` |
| Relay IPv6 address | Entry | same |
| Relay password | Password entry | same |
| Encryption curve | Select (p256, p384, p521, siec) | same |
| SOCKS5 proxy | Entry | env / config |
| HTTP proxy | Entry | env / config |
| Upload throttle | Entry (e.g. `500k`) | config |
| Classic mode | Checkbox | `classic_enabled` sentinel file |

All reads/writes go through the same config file paths croc already uses (`utils.GetConfigDir()`), so CLI and GUI share settings automatically.

### `progress.go` - The Bridge

The hardest problem: croc's `cr.Send()` and `cr.Receive()` are blocking calls that write progress to the logger. We need that progress in the UI.

**Approach:** Shared state struct + polling goroutine.

```go
type TransferState struct {
    mu          sync.RWMutex
    Status      string        // "idle" | "connecting" | "transferring" | "done" | "error"
    FilesTotal  int
    FilesDone   int
    BytesTotal  int64
    BytesDone   int64
    RateBps     int64
    CurrentFile string
    Code        string        // generated or provided
    Err         error
}
```

A polling goroutine (10ms tick) reads `TransferState` under `RLock` and calls `fyne.Do(func() { /* update widgets */ })` for thread-safe UI mutation. This avoids the complexity of redirecting stdout pipes while keeping UI updates smooth.

For richer integration in future: croc's `Options` struct could be extended with optional callback hooks (`OnProgress func(done, total int64)`). That's a clean upstream contribution path.

### `theme.go` - Fyne Theme

Custom `fyne.Theme` implementation:
- Colour palette: dark background (`#0f1117`), accent teal (`#00c9a7`), text primary (`#e8eaf0`)
- Font: embedded monospace for code/transfer display, sans-serif for UI chrome
- Slightly larger padding than Fyne defaults for breathing room
- Custom icons for send, receive, relay tabs

---

## Data Flow Diagram

```
┌────────────┐    croc.Options     ┌──────────────────┐
│  UI Layer  │ ──────────────────► │  croc.New()      │
│ (src/ui)   │                     │  (src/croc)      │
│            │ ◄────────────────── │                  │
│  Fyne      │   TransferState     │  cr.Send()       │
│  Widgets   │   (shared struct,   │  cr.Receive()    │
│            │    goroutine poll)  │                  │
└────────────┘                     └──────┬───────────┘
                                          │
                                          ▼
                                   ┌──────────────┐
                                   │  src/tcp     │
                                   │  src/comm    │
                                   │  (unmodified)│
                                   └──────────────┘
```

---

## Dependency Addition

Only one new top-level dependency:

```
fyne.io/fyne/v2 v2.5.x
```

The `fyne` CLI tool (`fyne package`) is a dev-time tool only, not a module dependency.

---

## Build Matrix

| OS | Arch | Fyne backend | Notes |
|---|---|---|---|
| Linux | amd64, arm64 | OpenGL (X11/Wayland) | Needs `libgl1-mesa-dev xorg-dev` |
| macOS | amd64, arm64 | Metal | Universal binary via `fyne package -os darwin` |
| Windows | amd64 | DirectX | No extra deps |

CI should produce binaries for all three in GitHub Actions using `fyne package`.

---

## What Is Explicitly NOT Changed

- `src/croc/*` - core transfer logic
- `src/tcp/*` - TCP relay
- `src/comm/*` - connection management
- `src/mnemonicode/*` - word list
- `src/models/*` - shared types
- `src/utils/*` - utilities
- `src/cli/cli.go` - CLI (only `main.go` changes how it's invoked)

The fork is intentionally shallow. Future croc upstream changes can be merged with minimal conflict.