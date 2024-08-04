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

    if (inputFilename == "" || outputFilename == "") {
        log.Fatal("No filenames given")
    }

    fmt.Println(option)

    shiftedImage := manip.PerformShift(inputFilename, option)

    checkpoint = time.Now()
    fmt.Printf("Generated image in: %v\n", checkpoint.Sub(start))

    SaveToPDFFile(shiftedImage, outputFilename)

    checkpoint = time.Now()
    fmt.Printf("Saved to file in: %v\n", checkpoint.Sub(start))
}

func HelpMenu() {
    menu := []string{
        "",
        "Menu:",
        "g - Greyscale (returns a txt of the greyscaled image)",
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
