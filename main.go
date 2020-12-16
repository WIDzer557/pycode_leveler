package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	var iPath, oPath string

	fmt.Printf("File: ")

	if len(os.Args) > 1 {
		iPath = os.Args[1]
		fmt.Printf("%s\n", iPath)
	} else {
		fmt.Scanf("%s", &iPath)
	}

	in, err := ioutil.ReadFile(iPath)
	if err != nil {
		log.Fatal(err)
	}

	var out string

	for _, line := range strings.FieldsFunc(string(in), func(r rune) bool {
		return r == '\n' || r == '\r' || r == ';'
	}) {
		c, j := tabCounter(line)
		out += strings.Repeat("\t", c) + line[j:] + "\n"
	}

	if len(os.Args) > 2 {
		oPath = strings.Join(os.Args[2:], " ")
	} else {
		oPath = iPath[:strings.LastIndex(iPath, ".")] + "_MM_aligned.py"
	}

	ioutil.WriteFile(oPath, []byte(out), 0777)
}

func tabCounter(str string) (c, i int) {
	for i, el := range str {
		switch el {
		case ' ':
			c++
		case '\t':
			c += 4
		default:
			return int(math.Ceil(float64(c)/4)), i
		}
	}

	return int(math.Ceil(float64(c)/4)), len(str)-1
}