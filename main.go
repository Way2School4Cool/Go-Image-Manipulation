package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// arg 0 == the exe?
// arg 1 == Option for manip
// arg 2 == File Input name
// arg 3 == File Output name
func main() {
    println(runtime.NumCPU())
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
        grayArray := GreyscaleManipulate(inputFilename)

        checkpoint = time.Now()
        fmt.Printf("Generated Array in: %v\n", checkpoint.Sub(start))

        SaveToTextFile(grayArray, outputFilename)

        checkpoint = time.Now()
        fmt.Printf("Saved to file in: %v\n", checkpoint.Sub(start))
    
    case "r":
        redImage := Redshift(inputFilename)
       
        checkpoint = time.Now()
        fmt.Printf("Generated Array in: %v\n", checkpoint.Sub(start))

        SaveToPDFFile(redImage, outputFilename)

        checkpoint = time.Now()
        fmt.Printf("Saved to file in: %v\n", checkpoint.Sub(start))

    default:
        HelpMenu()

    }
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

    fmt.Println(img)
 
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

func Redshift(inputFilename string) image.Image {
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

    // Loop through image line by line
    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            origColor := img.At(x, y)
            red, green, blue, alpha := origColor.RGBA() 
            
            newColor := color.RGBA{
                R: uint8(OverflowCheck((red + 50), 0,  255)),
                G: uint8(green),
                B: uint8(blue),
                A: uint8(alpha),
            }

            newImg.Set(x, y, newColor)
            // rgba = color.RGBA{rgba.R, rgba.G, rgba.B, rgba.A}
            //newImg.At(x,y) = color.NRGBA(rgba)
        }
    }

    return newImg
}

func OverflowCheck(value, min, max uint32) uint32 {
    if value < min {
        return min
    } else if value > max {
        return max
    }
    return value
}

func SaveToTextFile(dataToSave []string, outputFilename string) {
    var messyString string = ""
    messyString = messyString + strings.Join(dataToSave, "")

    err := os.WriteFile(outputFilename, []byte(messyString), 0666)
    if err != nil {
        log.Fatal(err)
    }
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
