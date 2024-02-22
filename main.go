package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/debug"

	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

const usage = `Usage:
    olaf [--version]

Options:
    --version  Print the version and exit.

At startup, olaf registers eight different hotkeys. Four of these hotkeys are
used to copy the contents of the system clipboard to a virtual clipboard, and
the other four are used to copy the contents of a virtual clipboard back into
the system clipboard.

The registered hotkeys are Ctrl+<H> and Alt+<H>, where <H> is the name of one
of the virtual clipboards. There are four virtual clipboards: u, i, o, and p.

You will usually want to run olaf as a background task, especially if you are
using the CLI. If you are using a bash-like shell, see the example.

Example:
    $ olaf &`

var Version string

func main() {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.Parse()

	if versionFlag {
		if Version != "" {
			fmt.Println(Version)
			return
		}
		if buildInfo, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(buildInfo.Main.Version)
			return
		}
		fmt.Println("(unknown)")
		return
	}

	mainthread.Init(run)
}

func run() {
	const n int = 4

	err := clipboard.Init()
	if err != nil {
		errorf("failed to initialize system clipboard: %v", err)
	}

	// We clean the clipboard formatting to avoid weird behaviors.
	clipboard.Write(clipboard.FmtText, clipboard.Read(clipboard.FmtText))

	keys := []hotkey.Key{hotkey.KeyU, hotkey.KeyI, hotkey.KeyO, hotkey.KeyP}
	copyHotkeys, err := registerCopy(keys)
	if err != nil {
		errorf("failed to register copy hotkeys: %v", err)
	}
	pasteHotkeys, err := registerPaste(keys)
	if err != nil {
		errorf("failed to register paste hotkeys: %v", err)
	}

	var cases []reflect.SelectCase
	for _, hk := range copyHotkeys {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(hk.Keydown()),
		})
	}
	for _, hk := range pasteHotkeys {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(hk.Keydown()),
		})
	}

	clipboards := make([]Clipboard, n)
	for {
		chosen, _, _ := reflect.Select(cases)
		switch {
		case chosen < n:
			clipboards[chosen] = clipboard.Read(clipboard.FmtText)
		default:
			i := chosen - n
			clipboard.Write(clipboard.FmtText, clipboards[i])
		}
	}
}

type Clipboard []byte

// l is a logger with no prefixes.
var l = log.New(os.Stderr, "", 0)

func printf(format string, v ...any) {
	l.Printf("olaf: "+format, v...)
}

func errorf(format string, v ...any) {
	l.Fatalf("olaf: error: "+format, v...)
}
