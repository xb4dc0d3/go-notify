# Go-Notify

The inotify wrapper using tmux as intermediary

*Based on: https://pkg.go.dev/github.com/fsnotify/fsnotify*


## How to

### Build

```bash
$ git clone https://github.com/xb4dc0d3/go-notify
$ go build -o go-notify .
```

### Run

1. Watch the `/tmp` directory
```bash
$ ./go-notify attach /tmp
```

2. Watch the `/tmp` directory and attach to tmux session with name `window-tmp`
```bash
$ /go-notify attach-tmux /tmp window-tmp
```