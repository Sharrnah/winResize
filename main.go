package main

import (
	"fmt"
	"github.com/vaughan0/go-ini"
	"strconv"
	"strings"
	"syscall"

	"github.com/JamesHovious/w32"
	"github.com/hnakamur/w32syscall"
)

func main() {
	cfg, err := ini.LoadFile("settings.ini")
	if err != nil {
		fmt.Println(err)
	}

	err = w32syscall.EnumWindows(func(hwnd syscall.Handle, lparam uintptr) bool {
		h := w32.HWND(hwnd)
		text := w32.GetWindowText(h)

		for sectionName, section := range cfg {
			sectionX, _ := strconv.Atoi(section["x"])
			sectionY, _ := strconv.Atoi(section["y"])
			sectionWidth, _ := strconv.Atoi(section["width"])
			sectionHeight, _ := strconv.Atoi(section["height"])

			if strings.Contains(text, sectionName) {
				fmt.Println(w32.GetWindowRect(h))
				w32.MoveWindow(h, sectionX, sectionY, sectionWidth, sectionHeight, true)
			}
		}
		return true
	}, 0)
	if err != nil {
		fmt.Println(err)
	}
}
