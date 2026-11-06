package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	candidate := flag.String("flag", "cryptoCTF{meowmeow}", "candidate CTF flag")
	leak := flag.Bool("leak", false, "display the cat clue")
	flag.Parse()
	if *leak {
		fmt.Println("Cat leak: meowmeow")
	}

	matched, err := verifyFlag(*candidate)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if matched {
		fmt.Println("Correct flag!")
		fmt.Println(*candidate)
	} else {
		fmt.Println("Wrong flag!")
		os.Exit(1)
	}
}

func verifyFlag(candidate string) (bool, error) {
	pattern := `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`
	return regexp.MatchString(pattern, candidate)
}
