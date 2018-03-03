package images

import (
	"image"
	"image/color"
	"math"
)

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

func CreateHOG(img *image.Gray) (HOG [][]uint8) {
	HOG = make([][]uint8, img.Bounds().Max.Y)
	for i := 0; i < img.Bounds().Max.Y; i++ {
		HOG[i] = make([]uint8, img.Bounds().Max.X)
	}
	for i := 0; i < img.Bounds().Max.Y; i++ {
		for j := 0; j < img.Bounds().Max.X; j++ {
			HOG[i][j] = darkestDirection(img, i, j)
		}
	}

	var BigHOG [math.Ceil(float64(img.Bounds().Max.Y) / 16)][math.Ceil(float64(img.Bounds().Max.X) / 16)]uint8

	for i := 0; i < img.Bounds().Max.Y; i += 16 {
		for j := 0; j < img.Bounds().Max.X; j += 16 {
			HOG[i][j] = darkestDirectionBig(HOG, i, j)
		}
	}
	return
}

func darkestDirection(img *image.Gray, i, j int) uint8 {
	var Y uint8

	if i > 0 && img.At(j, i-1).(color.Gray).Y > Y {
		return 1
	}
	if j < img.Bounds().Max.X-1 && img.At(j+1, i).(color.Gray).Y > Y {
		return 2
	}
	if i < img.Bounds().Max.Y-1 && img.At(j, i+1).(color.Gray).Y > Y {
		return 3
	}
	if j > 0 && img.At(j-1, i).(color.Gray).Y > Y {
		return 4
	}
	return 0
}

func darkestDirectionBig(HOG [][]uint8, i, j int) uint8 {
	//if i < len
	return 4
}
