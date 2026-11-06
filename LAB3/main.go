package main

import (
	"bufio"
	"fmt"
	"lab3/utils/crack"
	"os"
)

func main() {
	target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

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

		fmt.Println("Trying:", password, "Sha512:", hash)

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
