package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var input = "yzbqklnj"

var i = 1

func main() {
	for {
		combined := fmt.Sprintf("%s%d", input, i)
		byteInput := []byte(combined)
		md5Sum := fmt.Sprintf("%x", md5.Sum(byteInput))
		if strings.HasPrefix(md5Sum, "000000") {
			fmt.Printf("We have a winner: '%s' with '%d' as additional input\n", md5Sum, i)
			break
		}
		i++
	}
}
