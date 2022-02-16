package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	Menu()
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Menu() {
	for {
		fmt.Println("New game\nExit")

		switch ReadInput() {
		case "New game":
			fmt.Println("Let the game begin")
			StartGame()
		case "Exit":
			return
		case "":
			ClearTerminal()
		default:
			fmt.Println("invalid command")
			continue
		}
	}
}

func StartGame() {
	var (
		numStr string
		guess  int
	)

	fmt.Println("1P, please, enter the number")
	fmt.Scan(&numStr)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}

	ClearTerminal()
	fmt.Println("2P, please, enter the number")

	for fmt.Scan(&guess); num != guess; fmt.Scan(&guess) {
		if guess > num {
			fmt.Println("Try less")
		} else if guess < num {
			fmt.Println("Try greater")
		}
	}
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "linux", "darwin/arm64":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("Ambiguous OS!")
	}
}
