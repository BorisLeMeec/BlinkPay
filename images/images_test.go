package images

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"testing"
)

func openFile(fileName string) (img image.Image) {
	file, err := os.Open(fileName)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return
}

func BenchmarkCreateHOG(b *testing.B) {
	img := openFile("ProfilPicture_car.jpg")
	b.ResetTimer()
}

func BenchmarkGrayScale(b *testing.B) {
	img := openFile("ProfilPicture_car.jpg")
	b.ResetTimer()
}
