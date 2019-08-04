package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("You can enter a web address and check if it is up or not. Press Ctrl+c to close the app.\n")
	for { //infinite for
		fmt.Print("Enter a web address: ")
		scanner.Scan()                                                          //scan for value entered by command line
		text := scanner.Text()                                                  //convert value entered by scanner to string
		ping, _ := exec.Command("ping", text, "-c 5", "-i 2", "-w 10").Output() //executing a UNIX ping command with address as input and c,i,w as switches where c=request count, i=interval between packets and w=Time to wait for a response, in seconds.
		if strings.Contains(string(ping), "64 bytes from") {
			fmt.Printf("%s is up\n", text)
		} else {
			fmt.Printf("%s is down\n", text)
		}
	}
}
