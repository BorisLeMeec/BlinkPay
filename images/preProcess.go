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

func CreateHOG(img *image.Gray, img_original image.Image, visualization bool) {
	var cell_size = 8

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
	if visualization {
		go visualizeStep1(grads)
		go visualizeStep1Color(grads, img_original)
	}

	hsize := int(math.Ceil(float64(len(grads))/float64(cell_size))) + 1
	vsize := int(math.Ceil(float64(len(grads[0]))/float64(cell_size))) + 1
	var bins [][][9]float64 = make([][][9]float64, hsize)
	for i := 0; i < hsize; i++ {
		bins[i] = make([][9]float64, vsize)
	}
	i, j := 0, 0
	for y := 0; y < len(grads)-1; y += cell_size {
		j = 0
		for x := 0; x < len(grads[y])-1; x += cell_size {
			bins[i][j] = getBin(grads, x, y, cell_size)
			j++
		}
		i++
	}
	//n_bins := normalize(grads, bins)
	//fmt.Println(n_bins)
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
	if gradY == 0 {
		grad.Angle = 0
	} else {
		grad.Angle = uint8(int(math.Abs(math.Atan(float64(gradX)/float64(gradY))*math.Pi*180)) % 180)
	}
	return
}

func getBin(grads [][]GradientVector, x, y, cell_size int) (bin [9]float64) {
	//bin = make([]float64, 9)

	maxY := y + cell_size
	if maxY >= len(grads) {
		maxY = len(grads) - 1
	}
	maxX := x + cell_size
	if maxX >= len(grads[y]) {
		maxX = len(grads[y]) - 1
	}
	for i := y; i < maxY; i++ {
		for j := x; j < maxX; j++ {
			a := grads[i][j].Angle
			bin[a/20] = (100 - ((float64(a-(a/20)*20) / float64(20)) * 100)) / 100 * grads[i][j].Magnitude
			if a < 160 {
				bin[a/20+1] += (100 - ((float64(a-(a/20)*20) / float64(20)) * 100)) / 100 * grads[i][j].Magnitude
			} else {
				bin[0] += ((float64(a-(a/20)*20) / float64(20)) * 100) / 100 * grads[i][j].Magnitude
			}
		}
	}
	return
}

func normalize(grads [][]GradientVector, bins [][][9]float64) (normalizeb_bins [][][9]float64) {
	normalizeb_bins = make([][][9]float64, len(bins))
	for i := 0; i < len(bins); i++ {
		normalizeb_bins[i] = make([][9]float64, len(bins[i]))
	}
	for y := 0; y < len(bins)-1; y++ {
		for x := 0; x < len(bins[y])-1; x++ {
			val := 0
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					for a := 0; a < 9; a++ {
						val += int(math.Pow(bins[y+i][x+j][a], 2))
					}
				}
			}
			n := math.Sqrt(float64(val))
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					for a := 0; a < 9; a++ {
						normalizeb_bins[y+i][x+j][a] = bins[y+i][x+j][a] / float64(n)
					}
				}
			}
		}
	}
	return
}
