package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/pierods/mgade/decrypt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"syscall"
)

var helpFlag = flag.Bool("help", false, "")

func main() {

	flag.Parse()
	if *helpFlag {
		fmt.Println("Usage: mgade encryptedfile")
		fmt.Println("or  mgade encryptedfile clearfile")
		return
	}

	inFile := ""
	outFile := ""

	switch len(os.Args) {
	case 1:
		fmt.Print("Enter encrypted file name/path: ")
		consoleReader := bufio.NewScanner(os.Stdin)
		consoleReader.Scan()
		inFile = consoleReader.Text()

	case 2:
		inFile = os.Args[1]
		fmt.Println("Using encrypted file " + inFile)
	case 3:
		inFile = os.Args[1]
		outFile = os.Args[2]
		fmt.Println("Writing to clear file " + outFile)
	}

	password := getPassword()
	inData, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	clearData, err := decrypt.Open(inData, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	if outFile != "" {
		err = ioutil.WriteFile(outFile, clearData, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
	} else {
		fmt.Println(string(clearData))
	}
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if len(password) == 0 {
		fmt.Println("Empty password. Exiting.")
		os.Exit(2)
	}
	fmt.Print("\nConfirm password: ")
	passwordConfirm, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if bytes.Compare(password, passwordConfirm) != 0 {
		fmt.Println("\nPasswords don't match. Exiting")
		os.Exit(3)
	}
	fmt.Println()
	return password
}
