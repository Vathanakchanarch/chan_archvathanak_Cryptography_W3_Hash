package main

import (
	"bufio"
	"fmt"
	"lab2/utils/crack"
	"os"
)

func main() {
	target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()
		hash := crack.Hash(password)

		fmt.Println("Trying:", password, "Sha1:", hash)

		if hash == target {
			fmt.Println("Password found:", password)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Read error:", err)
		os.Exit(1)
	}

	fmt.Println("Password not found!")
}
