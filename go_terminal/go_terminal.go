package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const ONE_SECOND time.Duration = 1 * time.Second

type terminal struct {
	pwd          string
	dirs         map[string][]string
	files        map[string][]string
	list_content map[string]string
}

func (t *terminal) run() {
	for {
		cmd := t.getCMD()

		switch cmd {
		case "pwd\n":
			t.current_path()
		case "ls\n":
			fmt.Println("this is a ls")
		case "touch\n":
			fmt.Println("this is a touch")
		case "echo\n":
			fmt.Println("this is echo")
		case "cd\n":
			fmt.Println("this is cd")
		case "mkdir\n":
			fmt.Println("this is echo")
		default:
			fmt.Printf("%s: command not found\n", cmd[:len(cmd)-1])
		}

		fmt.Println("This is my terminal: ", *t)
	}

}

func (t *terminal) getCMD() string {
	fmt.Printf("/")
	keyWord, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return keyWord
}

func (t *terminal) current_path() {
	fmt.Println(t.pwd)
}

func (t *terminal) list_current_dir() {
	fmt.Println(t.dirs[t.pwd])
}

func main() {
	goTerminal := terminal{
		pwd: "/",
		dirs: map[string][]string{
			"/": {"/"},
		},
	}
	goTerminal.run()
}
