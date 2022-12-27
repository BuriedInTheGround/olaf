package main

import "golang.design/x/hotkey"

func register(keys []hotkey.Key, mods ...hotkey.Modifier) ([]*hotkey.Hotkey, error) {
	var hotkeys []*hotkey.Hotkey
	for i, key := range keys {
		hotkeys = append(hotkeys, hotkey.New(mods, key))
		err := hotkeys[i].Register()
		if err != nil {
			return nil, err
		}
	}
	return hotkeys, nil
}
