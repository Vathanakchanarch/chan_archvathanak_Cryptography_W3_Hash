package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	// Our two inputs
	inputA := "hello"
	inputB := "Hello"

	a := []byte(inputA)
	b := []byte(inputB)

	fmt.Println("Hash (MD5):")
	md5A := fmt.Sprintf("%x", md5.Sum(a))
	md5B := fmt.Sprintf("%x", md5.Sum(b))
	compare(md5A, md5B)

	fmt.Println("Hash (SHA1):")
	sha1A := fmt.Sprintf("%x", sha1.Sum(a))
	sha1B := fmt.Sprintf("%x", sha1.Sum(b))
	compare(sha1A, sha1B)

	fmt.Println("Hash (SHA256):")
	sha256A := fmt.Sprintf("%x", sha256.Sum256(a))
	sha256B := fmt.Sprintf("%x", sha256.Sum256(b))
	compare(sha256A, sha256B)

	fmt.Println("Hash (SHA512):")
	sha512A := fmt.Sprintf("%x", sha512.Sum512(a))
	sha512B := fmt.Sprintf("%x", sha512.Sum512(b))
	compare(sha512A, sha512B)

	fmt.Println("Hash (SHA3-256):")
	sha3A := fmt.Sprintf("%x", sha3.Sum256(a))
	sha3B := fmt.Sprintf("%x", sha3.Sum256(b))
	compare(sha3A, sha3B)
}

// Compare the two hash values
func compare(hashA string, hashB string) {
	fmt.Println("Output A =", hashA)
	fmt.Println("Output B =", hashB)

	if hashA == hashB {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	fmt.Println()
}
