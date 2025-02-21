package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"rsc.io/qr"
)

var quality map[string]qr.Level = map[string]qr.Level{"L": qr.L, "M": qr.M, "Q": qr.Q, "H": qr.H}

func main() {
	bdot := flag.String("b", "⬛️", "string to use for black dots")
	wdot := flag.String("w", "⬜️", "string to use for white dots")
	eccl := flag.String("l", "L", "error correction level: L,M,Q,H")
	flag.Parse()

	args := flag.Args()

	level, ok := quality[*eccl]
	if !ok {
		fmt.Println("illegal ECC level: ", *eccl)
		os.Exit(2)
	}

	var data string

	if len(args) < 1 {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("error reading stdin: ", err)
			os.Exit(3)
		}
		data = string(b)
	} else {
		data = args[0]
	}

	q, err := qr.Encode(data, level)
	if err != nil {
		fmt.Println("QR generation failed: ", err)
		os.Exit(3)
	}
	bintextEncode(os.Stdout, q, *bdot, *wdot)
}

// borders are required
func bintextEncode(w io.Writer, q *qr.Code, black, white string) (err error) {
	border := strings.Repeat(white, q.Size+2)
	fmt.Println(border)
	defer fmt.Println(border)
	for y := 0; y < q.Size; y++ {
		_, err = w.Write([]byte(white))
		if err != nil {
			return
		}
		for x := 0; x < q.Size; x++ {
			cm := white
			if q.Black(x, y) {
				cm = black
			}
			_, err = w.Write([]byte(cm))
			if err != nil {
				return
			}
		}
		_, err = w.Write([]byte(white + "\n"))
		if err != nil {
			return
		}
	}
	return
}
