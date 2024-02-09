//	func SpamMasker(s, spam string) string {
//		index := strings.Index(s, spam)
//		if index == -1 {
//			return s
//		}
//
//		bytes := []byte(s)
//		startIndex := index + 7
//
//		for j := startIndex; j < len(bytes); j++ {
//			if bytes[j] == ' ' {
//				break
//			} else {
//				bytes[j] = '*'
//			}
//		}
//		return string(bytes)
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func SpamMasker(s, spam string) string {
	bytes := []byte(s)
	subBytes := []byte(spam)

	for i := 0; i < len(s); i++ {
		match := true
		for j := 0; j < len(spam); j++ {
			if i+j >= len(s) || bytes[i+j] != subBytes[j] {
				match = false
			}
		}
		if match {
			for j := i + len(spam); j < len(s); j++ {
				if bytes[j] != ' ' && j < len(s) {
					bytes[j] = '*'
				} else {
					break
				}

			}
			i = i + len(spam)
		}
	}
	return string(bytes)
}

func main() {
	spam := "http://"
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	s := myscanner.Text()
	maskedString := SpamMasker(s, spam)
	fmt.Println(maskedString)
}
