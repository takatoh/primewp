package main

import (
	"fmt"
	"os"
	"flag"
	"image"
	"image/color"
	"image/png"
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
	q := fold(p, w)
//	for y := 0; y < h; y++ {
//		fmt.Println(q[y])
//	}

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	var c color.RGBA
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if q[y][x] {
				c = color.RGBA{255, 255, 255, 255}
			} else {
				c = color.RGBA{0, 0, 0, 255}
			}
			img.Set(x, y, c)
		}
	}

	pngFilename := "primewp.png"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func primes(n int) []bool {
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

	return p
}

func fold(p []bool, w int) [][]bool {
	r := make([][]bool, 0)
	s := make([]bool, 0)
	for i := 0; i < len(p); i++ {
		s = append(s, p[i])
		if i % w == (w - 1) {
			r = append(r, s)
			s = make([]bool, 0)
		}
	}
	return r
}
