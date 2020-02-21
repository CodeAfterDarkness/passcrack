package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"strings"
)

// The Plan
// Come up with a password and hash it
// Then try to find a collision that satisfies the hash
// It *might* be the original password

// What is a password
// What is a hash
// What is salt (in terms of hashing)

const target string = "b026324c6904b2a9cb4b88d6d61c81d1" // "1"

func doTheThing(init string) {

	out := []byte(init)

	length := 5

	for {
		in := out
		hashOut := md5.Sum(out) // [16]byte
		out = []byte(hex.EncodeToString(hashOut[:]))

		brokenString := strings.Split(string(out), "")
		var str string = string(out)[:length]
		for _, char := range brokenString[length:] {
			str = str + char
			if !strings.HasPrefix(target, str) {
				break
			}
		}
		if len(str) > length+1 {
			log.Printf("Hash of %s (%s) matched %d characters", string(in), string(out), len(str))

			if len(str) == 32 {
				os.Exit(0)
			}
		}

	}
}

func main() {

	go doTheThing("fgnhjd")
	go doTheThing("6sdgfnsnb")
	doTheThing("sdfhqe3")

}
