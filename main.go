package main

import (
	"fmt"
	t "time"

	bot "github.com/go-vgo/robotgo"
)

func main() {
	interval := 200 * t.Second

	for {
		fmt.Println("pressed:", t.Now().Local().Format("03:04:05 PM"))
		bot.KeyTap("shift")
		t.Sleep(interval)
	}
}
