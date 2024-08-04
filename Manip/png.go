package manip

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// func Redshift(inputFilename string) image.Image {
//     return PerformShift(inputFilename, "r")
// }
//
// func Greenshift(inputFilename string) image.Image {
//     return PerformShift(inputFilename, "g")
// }
//
// func Blueshift(inputFilename string) image.Image {
//     return PerformShift(inputFilename, "b")
// }
//
// func Alphashift(inputFilename string) image.Image {
//     return PerformShift(inputFilename, "a")
// }

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
            
            switch (shift) {
            case "r":
                newColor = color.RGBA{
                    R: uint8(OverflowCheck(red + 50)),
                    G: uint8(green),
                    B: uint8(blue),
                    A: uint8(alpha),
                }

            case "g":
                newColor = color.RGBA{
                    R: uint8(red),
                    G: uint8(OverflowCheck(green + 50)),
                    B: uint8(blue),
                    A: uint8(alpha),
                }
                break;

            case "b":
                newColor = color.RGBA{
                    R: uint8(red),
                    G: uint8(green),
                    B: uint8(OverflowCheck(blue + 50)),
                    A: uint8(alpha),
                }
                break;

            case "a":
                newColor = color.RGBA{
                    R: uint8(red),
                    G: uint8(green),
                    B: uint8(blue),
                    A: uint8(OverflowCheck(alpha + 50)),
                }
                break;

            default:
                newColor = origColor
            }

            newImg.Set(x, y, newColor)
        }
    }

    return newImg
}

func OverflowCheck(value uint32) uint32 {
    if value < 0 {
        return 0
    } else if value > 255 {
        return 255
    }
    return value
}
