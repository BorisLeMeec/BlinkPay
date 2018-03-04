package images

import (
	"image"
	"image/color"
	"math"
)

type GradientVector struct {
	Magnitude float64
	Angle     uint8
}

func GrayScale(img image.Image) (imgGS *image.Gray) {
	imgGS = image.NewGray(img.Bounds())

	for i := 0; i < img.Bounds().Max.Y; i++ {
		for j := 0; j < img.Bounds().Max.X; j++ {
			pix := img.At(j, i)
			imgGS.Set(j, i, color.GrayModel.Convert(pix))
		}
	}
	return
}

func CreateHOG(img *image.Gray, img_original image.Image) {
	var grads [][]GradientVector

	grads = make([][]GradientVector, img.Bounds().Max.Y)
	for i := 0; i < img.Bounds().Max.Y; i++ {
		grads[i] = make([]GradientVector, img.Bounds().Max.X)
	}
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			grads[y][x] = getGradient(img, y, x)
		}
	}
	visualizeStep1(grads)
	visualizeStep1Color(grads, img_original)
}

func getGradient(img *image.Gray, y, x int) (grad GradientVector) {
	var gradY, gradX float64
	if y == 0 {
		gradY = 0
	} else if y == img.Bounds().Max.Y-1 {
		gradY = 0
	} else {
		gradY = math.Abs(float64(img.GrayAt(x, y-1).Y) - float64(img.GrayAt(x, y+1).Y))
	}
	if x == 0 {
		gradX = 0
	} else if x == img.Bounds().Max.X-1 {
		gradX = 0
	} else {
		gradX = math.Abs(float64(img.GrayAt(x-1, y).Y) - float64(img.GrayAt(x+1, y).Y))
	}
	grad.Magnitude = math.Sqrt(math.Pow(float64(gradX), 2) + math.Pow(float64(gradY), 2))
	// avoid division by zero
	gradY = (map[bool]float64{true: 0, false: gradY})[gradY == 0] // equivalent of ternary operator
	grad.Angle = uint8(int(math.Atan(float64(gradX)/float64(gradY))*math.Pi*180) % 180)
	return
}
