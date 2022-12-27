package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() {
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

func clipboardName(hotkey *hotkey.Hotkey) string {
	registerCode, _, _ := strings.Cut(hotkey.String(), "+")
	register, _ := strconv.Atoi(registerCode)
	return fmt.Sprintf("%c", register)
}

var l = log.New(os.Stderr, "", 0)

func printf(format string, v ...any) {
	l.Printf("olaf: "+format, v...)
}

func errorf(format string, v ...any) {
	l.Printf("olaf: error: "+format, v...)
}
