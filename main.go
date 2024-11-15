package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	// Clear screen based on operating system
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clearScreen()
		fmt.Println("=================================")
		fmt.Println("         WAR CARD GAME           ")
		fmt.Println("=================================")
		fmt.Println("\n1. Start Game")
		fmt.Println("2. Exit")
		fmt.Print("\nSelect an option (1-2): ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			clearScreen()
			playWar()
			fmt.Println("\nPress Enter to return to menu...")
			fmt.Scanln()
		case "2":
			fmt.Println("\nThanks for playing!")
			return
		default:
			fmt.Println("\nInvalid choice. Press Enter to try again...")
			fmt.Scanln()
		}
	}
}
