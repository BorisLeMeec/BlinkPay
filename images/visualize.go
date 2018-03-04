package images

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func visualizeStep1(grads [][]GradientVector) {
	img := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{len(grads[0]), len(grads)}})
	for y := 0; y < len(grads); y++ {
		for x := 0; x < len(grads[y]); x++ {
			YDelta := uint8(grads[y][x].Magnitude)
			var Y uint8 = 128
			if YDelta > 30 {
				Y = 255
			}
			img.SetGray(x, y, color.Gray{Y})
		}
	}
	file, err := os.Create("v_step1.jpg")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	jpeg.Encode(file, img, &jpeg.Options{100})
}

func visualizeStep1Color(grads [][]GradientVector, img_original image.Image) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{len(grads[0]), len(grads)}})
	for y := 0; y < len(grads); y++ {
		for x := 0; x < len(grads[y]); x++ {
			YDelta := uint8(grads[y][x].Magnitude)
			var c = img_original.At(x, y)
			if YDelta < 30 {
				c = color.RGBA{0, 0, 0, 0}
			}
			img.Set(x, y, color.RGBAModel.Convert(c))
		}
	}
	file, err := os.Create("v_step1_color.jpg")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	jpeg.Encode(file, img, &jpeg.Options{100})
}
