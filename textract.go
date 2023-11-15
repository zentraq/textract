package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/otiai10/gosseract/v2"
)

const (
	VERSION_NUMBER = "1.0.0"
)

var (
	help, version bool
	img           string
)

func main() {
	flag.BoolVar(&help, "h", false, "Print the help menu")
	flag.BoolVar(&version, "v", false, "Print the version number")
	flag.StringVar(&img, "i", "", "Image to extract text from")
	flag.Parse()

	// program sanity check
	if help || version || img != "" {
		if help {
			print_usage()
			return
		}
		if version {
			fmt.Println(VERSION_NUMBER)
			return
		}

		if img == "" {
			fmt.Println("[!] Image file required")
			print_usage()
			return
		} else {
			err := validate_image(img)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				text, err := extract_text(img)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Text:\n\t")
					fmt.Println(text)
				}
			}
		}
	} else {
		print_usage()
	}
}

// print out the help/usage menu
func print_usage() {
	fmt.Println("Usage: textract -i path/to/image")
	fmt.Println("\t-h: Print the help menu")
	fmt.Println("\t-v: Print the version number")
	fmt.Println("\t-i: Image to extract text from")
}

// check whether the provided file exists and is not a directory
func validate_image(image_path string) error {
	fi, err := os.Stat(image_path)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", image_path)
	} else {
		if fi.IsDir() {
			return fmt.Errorf("[!] %s is a directory\n", image_path)
		}
		return nil
	}
}

// extract text from the provided image
func extract_text(image_path string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(image_path)

	text, err := client.Text()
	if err != nil {
		return "", err
	} else {
		return text, nil
	}
}
