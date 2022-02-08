package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
			ClearTerminal()
			fmt.Println("Let the game begin")
			StartGame()
		case "Exit":
			return
		case "":
			continue
		default:
			fmt.Println("Enter a proper command")
			continue
		}
	}
}

func StartGame() {
	var user1Inp, user2Inp int

	fmt.Println("1P, please, enter the number")
	fmt.Scan(&user1Inp)

	ClearTerminal()
	fmt.Println("2P, please, enter the number")

	for fmt.Scan(&user2Inp); user1Inp != user2Inp; fmt.Scan(&user2Inp) {
		if user2Inp > user1Inp {
			fmt.Println("Greater")
			continue
		} else if user2Inp < user1Inp {
			fmt.Println("Smaller")
			continue
		}
	}

	ClearTerminal()
	fmt.Println("Bingo!")
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("Ambiguous OS")
	}
}
