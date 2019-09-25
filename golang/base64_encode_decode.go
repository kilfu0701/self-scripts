package main

// go run base64_encode_decode.go
// go run base64_encode_decode.go help
// go run base64_encode_decode.go encode 5c9adba7d4579ef73cdc6992
// go run base64_encode_decode.go decode XJrbp9RXnvc83GmS

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strings"
)

func Hex2bin(s string) []byte {
	ret, err := hex.DecodeString(s)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return ret
}

func Bin2Hex(b []byte) string {
	return hex.EncodeToString(b)
}

func EncodeIdToBase64(input string) string {
	s := base64.StdEncoding.EncodeToString(Hex2bin(input))
	base64Str := strings.Map(func(r rune) rune {
		switch r {
		case '+':
			return '-'
		case '/':
			return '_'
		}

		return r
	}, s)

	base64Str = strings.ReplaceAll(base64Str, "=", "")
	return base64Str
}

func DecodeIdToBase64(input string) string {
	base64Str := strings.Map(func(r rune) rune {
		switch r {
		case '-':
			return '+'
		case '_':
			return '/'
		}

		return r
	}, input)

	if pad := len(base64Str) % 4; pad > 0 {
		base64Str += strings.Repeat("=", 4-pad)
	}

	b, _ := base64.StdEncoding.DecodeString(base64Str)
	return Bin2Hex(b)
}

func help() {
	fmt.Printf("==== Usage ====\n")
	fmt.Printf("go run base64_encode_decode.go\n")
	fmt.Printf("go run base64_encode_decode.go help\n")
	fmt.Printf("go run base64_encode_decode.go encode 5c9adba7d4579ef73cdc6992 ... ...\n")
	fmt.Printf("go run base64_encode_decode.go decode XJrbp9RXnvc83GmS ... ...\n\n")
}

func run_sample() {
	fmt.Printf("==== run sample ====\n")

	// here we use mongodb objectID string
	str := "5c9adba7d4579ef73cdc6992"

	e := EncodeIdToBase64(str)
	fmt.Printf("encoded = %s \n", e)

	d := DecodeIdToBase64(e)
	fmt.Printf("decoded = %s \n\n", d)
}

func main() {
	//flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		help()
		run_sample()
		return
	}

	if args[0] == "help" {
		help()
		return
	}

	if len(args) >= 2 {
		if args[0] == "encode" {
			for _, v := range args[1:] {
				fmt.Printf("[encode] input = %s, result = %s\n", v, EncodeIdToBase64(v))
			}
		} else if args[0] == "decode" {
			for _, v := range args[1:] {
				fmt.Printf("[decode] input = %s, result = %s\n", v, EncodeIdToBase64(v))
			}
		}
	}
}
