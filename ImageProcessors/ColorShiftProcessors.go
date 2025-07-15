package imageProcessors

import (
	models "ImageManipulation/Models"
	utilities "ImageManipulation/Utilities"

	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func PerformShift(inputFilename, shift string) image.Image {
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
			red, green, blue, alpha := origColor.RGBA()

			switch shift {
			case models.MenuOptions.RedShift:
				newColor = color.RGBA{
					R: uint8(utilities.OverflowCheck(red + 50)),
					G: uint8(green),
					B: uint8(blue),
					A: uint8(alpha),
				}

			case models.MenuOptions.GreenShift:
				newColor = color.RGBA{
					R: uint8(red),
					G: uint8(utilities.OverflowCheck(green + 50)),
					B: uint8(blue),
					A: uint8(alpha),
				}

			case models.MenuOptions.BlueShift:
				newColor = color.RGBA{
					R: uint8(red),
					G: uint8(green),
					B: uint8(utilities.OverflowCheck(blue + 50)),
					A: uint8(alpha),
				}

			case models.MenuOptions.AlphaShift:
				newColor = color.RGBA{
					R: uint8(red),
					G: uint8(green),
					B: uint8(blue),
					A: uint8(utilities.OverflowCheck(alpha + 50)),
				}

			default:
				newColor = origColor
			}

			newImg.Set(x, y, newColor)
		}
	}

	return newImg
}
