package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var thresholdDarkness int
var showVersion bool

func init() {
	flag.IntVar(&thresholdDarkness, "th", 150, "threshold of darkness out of 255; smaller is darker")
	flag.BoolVar(&showVersion, "v", false, "show version and exit")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Printf("%s %s\n", Version, GitCommit)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "input file is missing")
		os.Exit(1)
	}

	img := openImageFile(flag.Arg(0))

	dark := IsDark(img, thresholdDarkness)
	if dark {
		fmt.Println("dark")
	} else {
		fmt.Println("pale")
	}
}

func openImageFile(file string) image.Image {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot open file")
		os.Exit(1)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, "not a known image format")
		os.Exit(1)
	}

	return img
}

func IsDark(img image.Image, threshold int) bool {
	return Darkness(img) < threshold
}

func Darkness(img image.Image) int {
	bounds := img.Bounds()

	sum := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			sum += int(gray.Y)
		}
	}
	darkness := sum / bounds.Max.Y / bounds.Max.X
	return darkness
}
