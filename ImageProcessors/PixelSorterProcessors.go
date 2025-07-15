package imageProcessors

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"sort"
)

// PixelSort sorts the pixels of an image based on luminance
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

	// Prepare a new image
	newImg := image.NewRGBA(img.Bounds())

	// Loop through image line by line
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {

		// Perform sort based on luminanceImage in the x cords
		// while within the x cords
		var luminX int = 0
		var luminanceEdges []int32
		var newXCords []color.Color

		luminanceEdges = append(luminanceEdges, 0)

		// log edges of luminance changes
		for luminX < img.Bounds().Max.X {
			if luminanceImage.At(luminX, y) != luminanceImage.At(luminX+1, y) {

				// If the next pixel is not the same luminance, add the current pixel to the luminanceEdges
				if luminX+1 != img.Bounds().Max.X {
					luminanceEdges = append(luminanceEdges, int32(luminX))
				} else { // If the next pixel is the same luminance, add the next pixel to the luminanceEdges
					luminanceEdges = append(luminanceEdges, int32(luminX+1))
				}
			}
			luminX++
		}

		// for each luminance edge
		for index, value := range luminanceEdges {

			// If the last edge, break
			if index == len(luminanceEdges)-1 {
				break
			}

			var tempXCords []color.Color

			// Loop through the pixels between each edge
			for x := int(value); x < int(luminanceEdges[index+1]); x++ {
				tempXCords = append(tempXCords, img.At(x, y))
			}

			// Sort the pixels between each edge
			sort.Slice(tempXCords, func(i, j int) bool {
				return Luminance(tempXCords[i]) < Luminance(tempXCords[j])
			})

			// Create a new slice of the sorted pixels
			newXCords = append(newXCords, tempXCords...)
		}

		// Add the sorted slice to the new image
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			newImg.Set(x, y, newXCords[x])
		}
	}

	return newImg
}

// Sort sorts the pixels of an image based on luminance
func Sort(colorArray []color.Color) []color.Color {

	// If the array is less than 2, return the array
	if len(colorArray) < 2 {
		return colorArray
	}

	// Generate a random number in the range
	randomNumber := rand.Intn(len(colorArray))

	// Set a random pivot
	pivot := Luminance(colorArray[randomNumber])

	less := []color.Color{}
	greater := []color.Color{}

	// Partition the elements into less and greater slices
	for _, value := range colorArray[1:] {
		if Luminance(value) <= pivot {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}

	// Recursively sort the less and greater slices and combine them with the pivot
	return append(append(Sort(less), colorArray[randomNumber]), Sort(greater)...)
}
