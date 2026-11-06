package main

import (
	"bufio"
	"fmt"
	"lab1/utils/crack"
	"os"
)

func main() {
	target := "6a85dfd77d9cb35770c9dc6728d73d3f"

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

		fmt.Println("Trying:", password, "MD5:", hash)

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
