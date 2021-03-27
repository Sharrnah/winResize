package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/JamesHovious/w32"
	"github.com/hnakamur/w32syscall"
)

var update = false

func main() {

	argsWithoutProg := os.Args[1:]
	for _, arg := range argsWithoutProg {
		if arg == "update" {
			update = true
		}
	}

	cfg, err := ini.Load("settings.ini")
	if err != nil {
		fmt.Println(err)
	}
	secs := cfg.Sections()

	err = w32syscall.EnumWindows(func(hwnd syscall.Handle, lparam uintptr) bool {
		h := w32.HWND(hwnd)
		text := w32.GetWindowText(h)

		for _, section := range secs {
			sectionName := section.Name()
			if sectionName != "" {
				if strings.Contains(text, sectionName) {
					if !update {
						sectionX, _ := section.Key("x").Int()
						sectionY, _ := section.Key("y").Int()
						sectionWidth, _ := section.Key("width").Int()
						sectionHeight, _ := section.Key("height").Int()

						w32.MoveWindow(h, sectionX, sectionY, sectionWidth, sectionHeight, true)
						fmt.Println("Restored Position for \"" + text + "\"")
					} else {
						winRect := w32.GetWindowRect(h)
						sectionX := int(winRect.Left)
						sectionY := int(winRect.Top)
						sectionWidth := int(winRect.Right - winRect.Left)
						sectionHeight := int(winRect.Bottom - winRect.Top)

						cfg.Section(sectionName).Key("x").SetValue(strconv.Itoa(sectionX))
						cfg.Section(sectionName).Key("y").SetValue(strconv.Itoa(sectionY))
						cfg.Section(sectionName).Key("width").SetValue(strconv.Itoa(sectionWidth))
						cfg.Section(sectionName).Key("height").SetValue(strconv.Itoa(sectionHeight))

						fmt.Println("Found Position for \"" + text + "\"")
						fmt.Println("X="+ strconv.Itoa(sectionX) + " Y=" + strconv.Itoa(sectionY) + " Width=" + strconv.Itoa(sectionWidth) + " Height=" + strconv.Itoa(sectionHeight))
					}
				}
			}
		}
		return true
	}, 0)
	if err != nil {
		fmt.Println(err)
	}

	if update {
		err = cfg.SaveTo("settings.ini")
		if err != nil {
			fmt.Println(err)
			w32.MessageBox(0, "WinResize Save failed", "Window Positions saving failed.", 0)
		} else {
			fmt.Println("Saved Positions")
		}
	}
}
