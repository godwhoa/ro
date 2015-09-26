package main

import (
	"github.com/gojp/kana"
	"os"
)

func main() {
	if kana.IsLatin(os.Args[1]) {
		os.Stdout.Write([]byte(kana.RomajiToHiragana(os.Args[1]) + "\n"))
	}
	if kana.IsKana(os.Args[1]) {
		os.Stdout.Write([]byte(kana.KanaToRomaji(os.Args[1]) + "\n"))
	}
}
