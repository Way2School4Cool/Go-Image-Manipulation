package imageProcessors

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func LuminanceImage(inputFilename string) image.Image {
	// Attempt to open the image
	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Decode image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	newImg := image.NewRGBA(img.Bounds())
	var newColor color.Color

	// Loop through image line by line
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			origColor := img.At(x, y)

			if IsWithinLuminanceThreshold(origColor) {
				newColor = color.RGBA{
					R: uint8(0),
					G: uint8(0),
					B: uint8(0),
					A: uint8(255),
				}
				newImg.Set(x, y, newColor)
			} else {
				newColor = color.RGBA{
					R: uint8(255),
					G: uint8(255),
					B: uint8(255),
					A: uint8(255),
				}
				newImg.Set(x, y, newColor)
			}
		}
	}

	// fmt.Println(newImg)

	return newImg
}

// func GetLuminance(color color.Color) int

func IsWithinLuminanceThreshold(color color.Color) bool {

	red, green, blue, _ := color.RGBA()

	// fmt.Printf("r: %d, g: %d, b: %d", red, green, blue)

	const RLumin = float32(.2126)
	const GLumin = float32(.7252)
	const BLumin = float32(.0722)

	luminance := ((RLumin * float32(uint8(red))) + (GLumin * float32(uint8(green))) + (BLumin * float32(uint8(blue))))

	// fmt.Println(luminance)

	if luminance >= float32(80.0) && luminance <= float32(175.0) {
		return true
	}

	return false
}

// Calculate the luminance of an RGBA color
func Luminance(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	// Convert from 0-65535 range to 0-255 range
	r8 := float64(r >> 8)
	g8 := float64(g >> 8)
	b8 := float64(b >> 8)

	// Calculate luminance using the standard formula
	return 0.299*r8 + 0.587*g8 + 0.114*b8
}
