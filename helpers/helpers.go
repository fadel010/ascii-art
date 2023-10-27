package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Execute() {
	args := os.Args[1:]
	if len(args) == 1 {
		if len(args[0]) > 0 && TextVerification(args[0]) {
			if args[0] == "\\n" {
				fmt.Println()
			} else {
				lines := strings.Split(args[0], "\\n")
				for _, line := range lines {
					DisplayChar(line)
				}
			}
		} else if len(args[0]) > 0 {
			fmt.Println("Votre texte contient des caracteres non pris en charge.")
		}
	} else {
		fmt.Println("USAGE: 'go run . text")
	}

}

func GetAllChar() map[byte][]string {

	textByte, err := os.ReadFile("standard.txt")
	if err != nil {
		println("File not found:", err)
		os.Exit(0)
	}

	char := string(textByte)
	listChar := strings.Split(char, "\n")[1:]
	bowl := []string{}
	compt := 0
	allChar := map[byte][]string{}
	asciiNums := 32

	for _, v := range listChar {
		if compt != 8 {
			bowl = append(bowl, v)
			compt++
		} else {
			allChar[byte(asciiNums)] = bowl
			compt = 0
			bowl = []string{}
			asciiNums++
		}
	}
	return allChar
}

func GetSpecificChar(s string) [][]string {
	tabChar := [][]string{}
	allChar := GetAllChar()
	for _, v := range s {
		tabChar = append(tabChar, allChar[byte(v)])
	}
	return tabChar
}

func DisplayChar(s string) {
	tabChar := GetSpecificChar(s)
	if len(tabChar) > 0 {
		for line := range tabChar[0] {
			for char := range tabChar {
				fmt.Print(tabChar[char][line])
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}

func TextVerification(s string) bool {
	re := regexp.MustCompile(`[^[:ascii:]]`)
	return len(re.FindAllString(s, -1)) == 0
}
