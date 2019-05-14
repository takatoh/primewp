package main

import (
	"fmt"
	"os"
	"flag"
	"strconv"
//	"strings"
)

func main() {
	Usage := func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <width> <height>\n", os.Args[0])
	}
	opt_help := flag.Bool("h", false, "Help message")
	flag.Parse()

	if *opt_help || flag.NArg() < 1 {
		Usage()
		os.Exit(0)
	}
	w, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid argument.")
		os.Exit(1)
	}
	h, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid argument.")
		os.Exit(1)
	}

	n := w * h
	p := primes(n)
	for _, x := range p {
		fmt.Println(x)
	}
}

func primes(n int) []int {
	p := make([]bool, n + 1)
	// スライス p はゼロ値（false）で初期化されるので、2 と 3 以上の奇数だけ true に初期化する。
	if 2 < n {
		p[2] = true
		for i := 3; i <= n; i += 2 {
			p[i] = true
		}
	}
	// 3 以上の奇数を順にふるいにかける。
	for i := 3; i * i < n; i += 2 {
		if p[i] {
			for j := i + i; j <= n; j += i {
				p[j] = false
			}
		}
	}
	// 素数だけ pn に追加。
	pn := make([]int, 0)
	for i := 2; i <= n; i++ {
		if p[i] {
			pn = append(pn, i)
		}
	}
	return pn
}
