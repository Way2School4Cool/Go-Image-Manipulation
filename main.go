package main

import (
    "fmt"
	"image/color"
	"image/png"
	"log"
	"os"
    "strings" 
)

// arg 0 == the exe?
// arg 1 == file input name
// arg 2 == file output name
func main() {
    var inputFilename string = os.Args[1]
    var outputFilename string = os.Args[2]

    if (inputFilename == "" || outputFilename == "") {
        log.Fatal("No filenames given")
    }

    fmt.Printf("Input: %s\nOutput: %s", inputFilename, outputFilename)

    GreyscaleManipulate(inputFilename, outputFilename)
}

func GreyscaleManipulate(inputFilename, outputFilename string) {

    // Initialize
    levels := []string{" ", "░", "▒", "▓", "█"}
    messyArray := make([]string, 0)

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
 
    // Loop through image line by line
    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            
            // Get the color as greyscale
            c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
            level := c.Y / 51 // 51 * 5 = 255
            if level == 5 {
               level--
            }
            
            // Append the relative ascii char into an array 
            messyArray = append(messyArray, levels[level])
        }
        // If it is a new line, add that to the array too
        messyArray = append(messyArray, "\n")
    }

    saveToTextFile(messyArray, outputFilename)
}

func saveToTextFile(dataToSave []string, outputFilename string) {
    var messyString string = ""
    messyString = messyString + strings.Join(dataToSave, "")

    err := os.WriteFile(outputFilename, []byte(messyString), 0666)
    if err != nil {
        log.Fatal(err)
    }
}

