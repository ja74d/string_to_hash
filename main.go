package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/md5"
)

func main() {
	fmt.Println("Enter a string to hash:")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("SHA256:%x\n", sha256.Sum256([]byte(input)))
	fmt.Printf("MD5:%x\n", md5.Sum([]byte(input)))
}