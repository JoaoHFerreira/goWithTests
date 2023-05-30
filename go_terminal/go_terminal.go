package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const ONE_SECOND time.Duration = 1 * time.Second

func getTerminalCMD() string {
	keyWord, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return keyWord
}

func main() {
	terminalMap := map[string]string{
		"pwd": "/",
	}

	for {

		cmd := getTerminalCMD()
		time.Sleep(ONE_SECOND)
		fmt.Println("This is my terminal: ", cmd)
		fmt.Println("This is my terminal: ", terminalMap)
	}
}
