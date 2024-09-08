package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func AdjustingScale(img string, width uint, height uint) {

	file, err := os.Open(img)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//check file type and if the file type is jpeg or jpg > switch the decode method ..... shit..///
	// 이미지 디코딩
	imges, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	//imgresized := image.NewRGBA(image.Rect(0, 0, width, height))
	resizedimg := resize.Resize(width, height, imges, resize.Lanczos3)

	outfile, err := os.Create("resized_image.png")
	if err != nil {
		fmt.Println("error occured with", err)
		return
	}
	defer outfile.Close()
	err = png.Encode(outfile, resizedimg)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}
	fmt.Println("succeed to resize image to 300x400")
}
