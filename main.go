package main

import (
	manip "ImageManipulation/Manip"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

// arg 0 == the exe?
// arg 1 == Option for manip
// arg 2 == File Input name
// arg 3 == File Output name
func main() {
	start := time.Now()
	var checkpoint time.Time
	//var shiftedImage image.Image

	option := os.Args[1]
	inputFilename := os.Args[2]
	outputFilename := os.Args[3]
	var outputImage image.Image

	if inputFilename == "" || outputFilename == "" {
		log.Fatal("No filenames given")
	}

	switch option {
	case "l": // Generate Black and White Image based on luminance
		outputImage = manip.LuminanceImage(inputFilename)

	case "s": // Pixel Sorting (Based on luminance)
		luminanceImage := manip.LuminanceImage(inputFilename)

		checkpoint = time.Now()
		fmt.Printf("Generated luminance in: %v\n", checkpoint.Sub(start))

		outputImage = manip.PixelSort(inputFilename, luminanceImage)

	default: //Perform Shift based on color option
		outputImage = manip.PerformShift(inputFilename, option)
	}

	checkpoint = time.Now()
	fmt.Printf("Generated image in: %v\n", checkpoint.Sub(start))

	SaveToPDFFile(outputImage, outputFilename)

	checkpoint = time.Now()
	fmt.Printf("Saved to file in: %v\n", checkpoint.Sub(start))
}

func HelpMenu() {
	menu := []string{
		"",
		"Menu:",
		"g - Greyscale (returns a greyscaled image based on luminance)",
		"s - Pixel Sorting (returns a pixel sorted image based on luminance)",
		"r - Red Shift (returns a red shifted image)",
		"b - Blue Shift (returns a blue shifted image)",
		"a - Alpha Shift (returns a alpha shifted image)",
		"",
	}
	log.Fatal(strings.Join(menu, "\n"))
}

func SaveToPDFFile(image image.Image, outputFilename string) {
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("failed to create: %s\n%s\n", outputFilename, err)
	}

	defer outputFile.Close()

	err = png.Encode(outputFile, image)
	if err != nil {
		log.Fatalf("failed to encode: %s\n%s\n", outputFilename, err)
	}
}
