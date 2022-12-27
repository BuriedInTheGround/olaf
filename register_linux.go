package main

import "golang.design/x/hotkey"

func registerCopy(keys []hotkey.Key) ([]*hotkey.Hotkey, error) {
	return register(keys, hotkey.ModCtrl)
}

func registerPaste(keys []hotkey.Key) ([]*hotkey.Hotkey, error) {
	return register(keys, hotkey.Mod1)
}
