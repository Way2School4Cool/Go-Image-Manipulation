package main

import (
	"fmt"
	"image/color"
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

    option := os.Args[1]
    inputFilename := os.Args[2]
    outputFilename := os.Args[3]

    if (inputFilename == "" || outputFilename == "") {
        log.Fatal("No filenames given")
    }

    switch option {
    case "g":
        //fmt.Printf("Input: %s\nOutput: %s\n", inputFilename, outputFilename)
        grayArray := GreyscaleManipulate(inputFilename)

        checkpoint = time.Now()
        fmt.Printf("Generated Array in: %v\n", checkpoint.Sub(start))

        saveToTextFile(grayArray, outputFilename)

        checkpoint = time.Now()
        fmt.Printf("Saved to file in: %v\n", checkpoint.Sub(start))
    
    // case "?":
    //     redshift(inputFilename, outputFilename)
    
    default:
        helpMenu()

    }
}

func helpMenu() {
    menu := []string{
        "",
        "Menu:",
        "g - Greyscale (returns a txt of the greyscaled image)",
        "",
    }
    log.Fatal(strings.Join(menu, "\n"))
}

func GreyscaleManipulate(inputFilename string) []string {
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

    return messyArray
}

// func redshift(inputFilename, outputFilename string) {
//     // Attempt to open the image
//     file, err := os.Open(inputFilename)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     defer file.Close()
//
//     // Decode image
//     img, err := png.Decode(file)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     newImg := img
//
//     // Loop through image line by line
//     for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
//         for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
//             rgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
//             rgba = color.RGBA{rgba.R, rgba.G, rgba.B, rgba.A}
//             newImg.At(x,y) = color.NRGBA(rgba)
//         }
//     }
// }

func saveToTextFile(dataToSave []string, outputFilename string) {
    var messyString string = ""
    messyString = messyString + strings.Join(dataToSave, "")

    err := os.WriteFile(outputFilename, []byte(messyString), 0666)
    if err != nil {
        log.Fatal(err)
    }
}

