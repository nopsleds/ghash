package main

import (
	"fmt"
	"os"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"hash"
)

func applyHash(f *os.File, hashName string, hash hash.Hash) {
	f.Seek(0, 0)
	io.Copy(hash, f)
	fmt.Printf(" - %s : %x\n", hashName, hash.Sum(nil))
}


func processFile(filePath string) {

	// open file
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error while opening file: %v", err)
		return
	}

	defer f.Close()

	fmt.Printf("Hashes for file '%v':\n", filePath)
	applyHash(f, "MD5   ", md5.New())
	applyHash(f, "SHA1  ", sha1.New())
	applyHash(f, "SHA256", sha256.New())
	fmt.Println()
}

func main(){
	for _,f := range os.Args[1:] {
		processFile(f)
	}
}
