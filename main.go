package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	// Specify the directory path containing the images
	directory := "./images"

	// Specify the output file paths
	csvFilePath := "./output/output.csv"
	htmlFilePath := "./output/output.html"

	// Open the CSV file for writing
	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a CSV writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Write the CSV header
	csvWriter.Write([]string{"Image File", "GPS Latitude", "GPS Longitude"})

	// Open the HTML file for writing
	htmlFile, err := os.Create(htmlFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer htmlFile.Close()

	// Write the HTML header
	htmlFile.WriteString("<html><body><table><tr><th>Image File</th><th>GPS Latitude</th><th>GPS Longitude</th></tr>")

	// Walk through the directory and its subdirectories
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is an image
		if !info.IsDir() && isImageFile(path) {
			// Read the EXIF data from the image
			exifData, err := readEXIFData(path)
			if err != nil {
				log.Printf("Error reading EXIF data from %s: %s", path, err)
				return nil
			}

			// Extract GPS latitude and longitude
			latitude, longitude, err := extractGPSData(exifData)
			if err != nil {
				log.Printf("Error extracting GPS data from %s: %s", path, err)
				return nil
			}

			// Write the attributes to the CSV file
			csvWriter.Write([]string{path, latitude, longitude})

			// Write the attributes to the HTML file
			htmlFile.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", path, latitude, longitude))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Flush the CSV writer
	csvWriter.Flush()

	// Check for any errors during CSV writing
	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}

	// Write the HTML footer
	htmlFile.WriteString("</table></body></html>")

	fmt.Println("CSV file created:", csvFilePath)
	fmt.Println("HTML file created:", htmlFilePath)
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

func readEXIFData(filepath string) (exifData *exif.Exif, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	exifData, err = exif.Decode(file)
	if err != nil {
		return nil, err
	}

	return exifData, nil
}

func extractGPSData(exifData *exif.Exif) (latitude string, longitude string, err error) {
	lat, long, err := exifData.LatLong()
	if err != nil {
		return "", "", err
	}

	latitude = fmt.Sprintf("%f", lat)
	longitude = fmt.Sprintf("%f", long)

	return latitude, longitude, nil
}
