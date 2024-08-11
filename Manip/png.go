package manip

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
    "sort"
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

    luminance := ((RLumin*float32(uint8(red))) + (GLumin*float32(uint8(green))) + (BLumin*float32(uint8(blue))))

    // fmt.Println(luminance)

    if (luminance >= float32(100.0) && luminance <= float32(200.0)) {
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

func PixelSort(inputFilename string, luminanceImage image.Image) image.Image {
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
    // var newColor color.Color

    // Loop through image line by line
    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {

        // Perform sort based on luminanceImage in the x cords
        // while within the x cords
        var luminX int = 0
        var luminanceEdges []int32
        var newXCords []color.Color

        luminanceEdges = append(luminanceEdges, 0)

        // log edges of luminance changes
        for (luminX < img.Bounds().Max.X) {
            if (luminanceImage.At(luminX, y) != luminanceImage.At(luminX + 1, y)) {
                if luminX + 1 != img.Bounds().Max.X {
                    luminanceEdges = append(luminanceEdges, int32(luminX))
                } else {
                    luminanceEdges = append(luminanceEdges, int32(luminX + 1))
                }
            }
            luminX++
        }

        // fmt.Println(luminanceEdges)

        // for each luminance edge
        for index, value := range luminanceEdges {
            if (index == len(luminanceEdges) - 1) {
                break
            }

            // fmt.Println(int(value))
            // fmt.Println(int(luminanceEdges[index]))

            var tempXCords []color.Color

            // TODO: This is F'ed
            // get a snip of the colors in the range from edge x to edge x+1
            for x := int(value); x < int(luminanceEdges[index + 1]); x++ {
                tempXCords = append(tempXCords, img.At(x, y))
                // fmt.Print(x)
            }

            // newXCords = SortColorsByLuminance(newXCords)

            // fmt.Println(len(newXCords))

            // fmt.Println(newXCords)

            sort.Slice(tempXCords, func(i, j int) bool {
                return Luminance(tempXCords[i]) < Luminance(tempXCords[j])
            })

            newXCords = append(newXCords, tempXCords...)

            // fmt.Println(newXCords)
        }

        fmt.Println(len(newXCords))
        // fmt.Println(img.Bounds().Max.X)

        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            // fmt.Println(newXCords[x])
            newImg.Set(x, y, newXCords[x])
        }
    }

    return newImg
}

// func Sort(colorArray []color.Color) []color.Color {
//     // var newColorArray []color.Color
//
//     if len(colorArray) < 2 {
// 		return colorArray
// 	}
//
// 	// Choose a pivot (you can choose any element, here we choose the first element)
// 	pivot := colorArray[0]
//
// 	// Slices to hold the partitioned elements
// 	less := []int{}
// 	greater := []int{}
//
// 	// Partition the elements into less and greater slices
// 	for _, value := range colorArray[1:] {
// 		if value <= pivot {
// 			less = append(less, value)
// 		} else {
// 			greater = append(greater, value)
// 		}
// 	}
//
// 	// Recursively sort the less and greater slices and combine them with the pivot
// 	return append(append(Sort(less), pivot), Sort(greater)...)
//
//     // if len(newColorArray) != len(colorArray) {
//     //     fmt.Println("array size mismatch")
//     // }
//     //
//     // return newColorArray
// }

func OverflowCheck(value uint32) uint32 {
    if value < 0 {
        return 0
    } else if value > 255 {
        return 255
    }
    return value
}
