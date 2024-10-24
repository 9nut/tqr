package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"rsc.io/qr"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("argument missing")
		os.Exit(1)
	}
	q, err := qr.Encode(args[0], qr.L)
	if err != nil {
		fmt.Println("QR generation failed: ", err)
		os.Exit(2)
	}
	dotsAtEncode(os.Stdout, q)
}

// borders are required
func dotsAtEncode(w io.Writer, q *qr.Code) (err error) {
	border := strings.Repeat("⬜️", q.Size+2)
	fmt.Println(border)
	defer fmt.Println(border)
	for y := 0; y < q.Size; y++ {
		_, err = w.Write([]byte("⬜️"))
		if err != nil {
			return
		}
		for x := 0; x < q.Size; x++ {
			cm := "⬜️"
			if q.Black(x, y) {
				cm = "⬛️"
			}
			_, err = w.Write([]byte(cm))
			if err != nil {
				return
			}
		}
		_, err = w.Write([]byte("⬜️\n"))
		if err != nil {
			return
		}
	}
	return
}
