package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/nfnt/resize"
)

// Unicode characters representing different brightness levels
var unicodeChars = []rune(" ░▒▓█")

// encodeImage encodes the input image to JPEG format and returns the encoded bytes
func encodeImage(inputPath string) ([]byte, error) {
	// Read the input image file
	imgFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	// Decode the input image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	// Encode the image to JPEG format
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// decodeImage decodes the encoded image data and returns the decoded image
func decodeImage(encodedData []byte) (image.Image, error) {
	// Create a reader for the encoded data
	reader := bytes.NewReader(encodedData)

	// Decode the image from the reader
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// displayImageWithUnicode converts the image to a Unicode representation and displays it
func displayImageWithUnicode(img image.Image) {
	// Resize the image to a smaller size for better representation
	img = resize.Resize(50, 0, img, resize.Lanczos3)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Convert each pixel to a corresponding Unicode character based on its brightness level
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the brightness level of the pixel
			r, g, b, _ := img.At(x, y).RGBA()
			avgBrightness := (float64(r) + float64(g) + float64(b)) / (3 * 0xffff) // Normalize to a value between 0 and 1

			// Determine the corresponding Unicode character based on the brightness level
			charIndex := int(float64(len(unicodeChars)-1) * avgBrightness)

			// Print the Unicode character
			fmt.Print(string(unicodeChars[charIndex]))
		}
		fmt.Println() // Move to the next line after each row
	}
}

func main() {
	// Define command-line flags
	encodeFlag := flag.Bool("encode", false, "Encode the input image to bytes")
	decodeFlag := flag.Bool("decode", false, "Decode the encoded data and display the image")
	inputFlag := flag.String("input", "", "Input file path")
	helpFlag := flag.Bool("help", false, "Show help message")

	// Parse command-line flags
	flag.Parse()

	// Show help message if requested
	if *helpFlag || flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// Perform the requested operation
	if *encodeFlag {
		// Encoding logic
		encodedData, err := encodeImage(*inputFlag)
		if err != nil {
			fmt.Println("Error encoding image:", err)
			os.Exit(1)
		}

		// Write the encoded data to standard output
		if _, err := os.Stdout.Write(encodedData); err != nil {
			fmt.Println("Error writing encoded data to standard output:", err)
			os.Exit(1)
		}
	} else if *decodeFlag {
		// Decoding logic

		// Read the encoded bytes from the file
		encodedData, err := ioutil.ReadFile(*inputFlag)
		if err != nil {
			fmt.Println("Error reading encoded data:", err)
			os.Exit(1)
		}

		// Decode the image from the bytes
		decodedImg, err := decodeImage(encodedData)
		if err != nil {
			fmt.Println("Error decoding image:", err)
			os.Exit(1)
		}

		// Display the decoded image with Unicode representation
		displayImageWithUnicode(decodedImg)
	} else {
		fmt.Println("No valid operation specified. Use -help for usage information.")
		os.Exit(1)
	}
}
