//go:build !windows && !cgo

package main

import "golang.design/x/hotkey"

func registerCopy(keys []hotkey.Key) ([]*hotkey.Hotkey, error) {
	panic("olaf: cannot use when CGO_ENABLED=0")
}

func registerPaste(keys []hotkey.Key) ([]*hotkey.Hotkey, error) {
	panic("olaf: cannot use when CGO_ENABLED=0")
}
