package main

import (
	"fmt"
	"os"
	"flag"
	"image"
	"image/color"
	"image/png"
	"strconv"
	"strings"
)

const (
	progVersion = "0.1.0"
)

func main() {
	Usage := func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <width> <height>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
	}
	opt_front := flag.String("f", "#FFFFFF", "Set color for prime number")
	opt_back := flag.String("b", "#000000", "Set color for not prime number")
	opt_help := flag.Bool("h", false, "Help message")
	opt_version := flag.Bool("v", false, "Show version")
	flag.Parse()

	if *opt_help {
		Usage()
		os.Exit(0)
	} else if *opt_version {
		fmt.Fprintf(os.Stderr, "v%s\n", progVersion)
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

	p := primes(w * h)
	q := fold(p, w)

	var colorcode string
	if strings.Index(*opt_front, "#") == 0 {
		colorcode = *opt_front
	} else {
		colorcode = colorname2code(*opt_front)
	}
	red, green, blue, e := code2rgb(colorcode)
	if e != nil {
		fmt.Fprintln(os.Stderr, "Invalide color code or name.")
		os.Exit(1)
	}
	c := color.RGBA{red, green, blue, 255}
	if strings.Index(*opt_back, "#") == 0 {
		colorcode = *opt_back
	} else {
		colorcode = colorname2code(*opt_back)
	}
	red, green, blue, e = code2rgb(colorcode)
	if e != nil {
		fmt.Fprintln(os.Stderr, "Invalide color code or name.")
		os.Exit(1)
	}
	b := color.RGBA{red, green, blue, 255}
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if q[y][x] {
				img.Set(x, y, c)
			} else {
				img.Set(x, y, b)
			}
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

	return p[1:]
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

func code2rgb(code string) (uint8, uint8, uint8, error) {
	sr := code[1:3]
	sg := code[3:5]
	sb := code[5:7]
	r, e1 := strconv.ParseUint(sr, 16, 64)
	if e1 != nil { return 0, 0, 0, e1 }
	g, e2 := strconv.ParseUint(sg, 16, 64)
	if e2 != nil { return 0, 0, 0, e2 }
	b, e3 := strconv.ParseUint(sb, 16, 64)
	if e3 != nil { return 0, 0, 0, e3 }
	return uint8(r), uint8(g), uint8(b), nil
}

type Color struct {
	code string
	name string
}

func newColor(code, name string) *Color {
	c := new(Color)
	c.code = code
	c.name = name
	return c
}

func colorname2code(name string) string {
	var webcolors = []*Color {
		newColor("#F0F8FF", "AliceBlue"),
		newColor("#9966CC", "Amethyst"),
		newColor("#FAEBD7", "AntiqueWhite"),
		newColor("#00FFFF", "Aqua"),
		newColor("#7FFFD4", "Aquamarine"),
		newColor("#F0FFFF", "Azure"),
		newColor("#F5F5DC", "Beige"),
		newColor("#FFE4C4", "Bisque"),
		newColor("#000000", "Black"),
		newColor("#FFEBCD", "BlanchedAlmond"),
		newColor("#0000FF", "Blue"),
		newColor("#8A2BE2", "BlueViolet"),
		newColor("#A52A2A", "Brown"),
		newColor("#DEB887", "BurlyWood"),
		newColor("#5F9EA0", "CadetBlue"),
		newColor("#7FFF00", "Chartreuse"),
		newColor("#D2691E", "Chocolate"),
		newColor("#FF7F50", "Coral"),
		newColor("#6495ED", "CornflowerBlue"),
		newColor("#FFF8DC", "Cornsilk"),
		newColor("#DC143C", "Crimson"),
		newColor("#00FFFF", "Cyan"),
		newColor("#00008B", "DarkBlue"),
		newColor("#008B8B", "DarkCyan"),
		newColor("#B8860B", "DarkGoldenrod"),
		newColor("#A9A9A9", "DarkGray"),
		newColor("#006400", "DarkGreen"),
		newColor("#BDB76B", "DarkKhaki"),
		newColor("#8B008B", "DarkMagenta"),
		newColor("#556B2F", "DarkOliveGreen"),
		newColor("#FF8C00", "DarkOrange"),
		newColor("#9932CC", "DarkOrchid"),
		newColor("#8B0000", "DarkRed"),
		newColor("#E9967A", "DarkSalmon"),
		newColor("#8FBC8F", "DarkSeaGreen"),
		newColor("#483D8B", "DarkSlateBlue"),
		newColor("#2F4F4F", "DarkSlateGray"),
		newColor("#00CED1", "DarkTurquoise"),
		newColor("#9400D3", "DarkViolet"),
		newColor("#FF1493", "DeepPink"),
		newColor("#00BFFF", "DeepSkyBlue"),
		newColor("#696969", "DimGray"),
		newColor("#1E90FF", "DodgerBlue"),
		newColor("#B22222", "FireBrick"),
		newColor("#FFFAF0", "FloralWhite"),
		newColor("#228B22", "ForestGreen"),
		newColor("#FF00FF", "Fuchsia"),
		newColor("#DCDCDC", "Gainsboro"),
		newColor("#F8F8FF", "GhostWhite"),
		newColor("#FFD700", "Gold"),
		newColor("#DAA520", "Goldenrod"),
		newColor("#808080", "Gray"),
		newColor("#008000", "Green"),
		newColor("#ADFF2F", "GreenYellow"),
		newColor("#F0FFF0", "Honeydew"),
		newColor("#FF69B4", "HotPink"),
		newColor("#CD5C5C", "IndianRed"),
		newColor("#4B0082", "Indigo"),
		newColor("#FFFFF0", "Ivory"),
		newColor("#F0E68C", "Khaki"),
		newColor("#E6E6FA", "Lavender"),
		newColor("#FFF0F5", "LavenderBlush"),
		newColor("#7CFC00", "LawnGreen"),
		newColor("#FFFACD", "LemonChiffon"),
		newColor("#ADD8E6", "LightBlue"),
		newColor("#F08080", "LightCoral"),
		newColor("#E0FFFF", "LightCyan"),
		newColor("#FAFAD2", "LightGoldenrodYellow"),
		newColor("#90EE90", "LightGreen"),
		newColor("#D3D3D3", "LightGrey"),
		newColor("#FFB6C1", "LightPink"),
		newColor("#FFA07A", "LightSalmon"),
		newColor("#20B2AA", "LightSeaGreen"),
		newColor("#87CEFA", "LightSkyBlue"),
		newColor("#778899", "LightSlateGray"),
		newColor("#B0C4DE", "LightSteelBlue"),
		newColor("#FFFFE0", "LightYellow"),
		newColor("#00FF00", "Lime"),
		newColor("#32CD32", "LimeGreen"),
		newColor("#FAF0E6", "Linen"),
		newColor("#FF00FF", "Magenta"),
		newColor("#800000", "Maroon"),
		newColor("#66CDAA", "MediumAquamarine"),
		newColor("#0000CD", "MediumBlue"),
		newColor("#BA55D3", "MediumOrchid"),
		newColor("#9370DB", "MediumPurple"),
		newColor("#3CB371", "MediumSeaGreen"),
		newColor("#7B68EE", "MediumSlateBlue"),
		newColor("#7B68EE", "MediumSlateBlue"),
		newColor("#00FA9A", "MediumSpringGreen"),
		newColor("#48D1CC", "MediumTurquoise"),
		newColor("#C71585", "MediumVioletRed"),
		newColor("#191970", "MidnightBlue"),
		newColor("#F5FFFA", "MintCream"),
		newColor("#FFE4E1", "MistyRose"),
		newColor("#FFE4B5", "Moccasin"),
		newColor("#FFDEAD", "NavajoWhite"),
		newColor("#000080", "Navy"),
		newColor("#FDF5E6", "OldLace"),
		newColor("#808000", "Olive"),
		newColor("#6B8E23", "OliveDrab"),
		newColor("#FFA500", "Orange"),
		newColor("#FF4500", "OrangeRed"),
		newColor("#DA70D6", "Orchid"),
		newColor("#EEE8AA", "PaleGoldenrod"),
		newColor("#98FB98", "PaleGreen"),
		newColor("#AFEEEE", "PaleTurquoise"),
		newColor("#DB7093", "PaleVioletRed"),
		newColor("#FFEFD5", "PapayaWhip"),
		newColor("#FFDAB9", "PeachPuff"),
		newColor("#CD853F", "Peru"),
		newColor("#FFC0CB", "Pink"),
		newColor("#DDA0DD", "Plum"),
		newColor("#B0E0E6", "PowderBlue"),
		newColor("#800080", "Purple"),
		newColor("#FF0000", "Red"),
		newColor("#BC8F8F", "RosyBrown"),
		newColor("#4169E1", "RoyalBlue"),
		newColor("#8B4513", "SaddleBrown"),
		newColor("#FA8072", "Salmon"),
		newColor("#F4A460", "SandyBrown"),
		newColor("#2E8B57", "SeaGreen"),
		newColor("#FFF5EE", "Seashell"),
		newColor("#A0522D", "Sienna"),
		newColor("#C0C0C0", "Silver"),
		newColor("#87CEEB", "SkyBlue"),
		newColor("#6A5ACD", "SlateBlue"),
		newColor("#708090", "SlateGray"),
		newColor("#FFFAFA", "Snow"),
		newColor("#00FF7F", "SpringGreen"),
		newColor("#4682B4", "SteelBlue"),
		newColor("#D2B48C", "Tan"),
		newColor("#008080", "Teal"),
		newColor("#D8BFD8", "Thistle"),
		newColor("#FF6347", "Tomato"),
		newColor("#40E0D0", "Turquoise"),
		newColor("#EE82EE", "Violet"),
		newColor("#F5DEB3", "Wheat"),
		newColor("#FFFFFF", "White"),
		newColor("#F5F5F5", "WhiteSmoke"),
		newColor("#FFFF00", "Yellow"),
		newColor("#9ACD32", "YellowGreen"),
	}

	code := searchColorCode(webcolors, name)
	return code
}

func searchColorCode(colors []*Color, name string) string {
	if len(colors) == 0 {
		return "unknown"
	}
	left := 0
	right := len(colors) - 1
	mid := right / 2
	if colors[mid].name == name {
		return colors[mid].code
	} else if name < colors[mid].name {
		return searchColorCode(colors[left:mid], name)
	} else {
		return searchColorCode(colors[mid+1:right+1], name)
	}
}
