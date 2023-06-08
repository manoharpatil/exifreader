---
### Problem:
Given the attached directory that contains images, and that also contains a sub-directory that contains images, create a command-line utility written in Go that reads the EXIF data from the images and writes the following attributes to a CSV file:

1. Image file path
2. GPS position latitude
3. GPS position longitude

Feel free to use go-exif, or other exif-extraction library, to read the EXIF data.

For extra credit, make it capable of writing to HTML as well.

https://go.dev

https://en.wikipedia.org/wiki/Exif

https://github.com/dsoprea/go-exif

---

---
### main.go Explanation

This program is a command-line tool that processes images in a specified directory and extracts GPS latitude and longitude data from their EXIF metadata. It then writes the extracted data to a CSV file and generates an HTML file for displaying the data in a table format.

Let's go through the code step by step:

> The program imports necessary packages: encoding/csv for working with CSV files, fmt for formatted output, log for logging errors, os for file operations, path/filepath for working with file paths, strings for string operations, and github.com/rwcarlsen/goexif/exif for extracting EXIF data from images.

> The main() function is the entry point of the program. It starts by specifying the directory path containing the images (./images) and the output file paths for the CSV and HTML files (./output.csv and ./output.html, respectively).

> The program opens the CSV file for writing using os.Create(). If an error occurs, it logs the error and terminates the program. The defer statement ensures that the file is closed at the end of the function.

> A CSV writer is created using csv.NewWriter() to write data to the CSV file. The defer statement ensures that the writer is flushed and any buffered data is written to the file at the end of the function.

> The program writes the CSV header to the file using csvWriter.Write([]string{"Image File", "GPS Latitude", "GPS Longitude"}).

> The program opens the HTML file for writing using os.Create(). If an error occurs, it logs the error and terminates the program. The defer statement ensures that the file is closed at the end of the function.

> The program writes the HTML header to the file, which contains the opening tags for an HTML table.

> The program uses filepath.Walk() to traverse the directory and its subdirectories. For each file encountered, it checks if the file is not a directory and is an image file (using the isImageFile() function).

> If the file is an image, the program reads the EXIF data from the image using the readEXIFData() function.

> The program extracts the GPS latitude and longitude from the EXIF data using the extractGPSData() function.

> The program writes the attributes (file path, latitude, and longitude) to the CSV file using csvWriter.Write([]string{path, latitude, longitude}).

> The program writes the attributes to the HTML file as a table row using htmlFile.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", path, latitude, longitude)).

> After processing all files, the program flushes the CSV writer to ensure all data is written to the file. It then checks for any errors during CSV writing using csvWriter.Error().

> The program writes the HTML footer, which contains the closing tags for the HTML table, to the HTML file.

> Finally, the program prints a message indicating the paths of the created CSV and HTML files.

>The program also includes three helper functions:
>>* isImageFile() determines if a given file is an image file based on its extension (.jpg, .jpeg, or .png).
>>* readEXIFData() reads the EXIF data from a specified file using the exif.Decode() function from the github.com/rwcarlsen/goexif/exif package. 
>>* extractGPSData() extracts the GPS latitude and longitude from the given EXIF data using the exifData.LatLong() method.

> Overall, this program demonstrates how to process images, read EXIF metadata, and extract specific information from it. It also showcases file I/O operations for writing data to CSV and HTML files.

---

---
### Output:

#### CSV Output
![CSV-Output](docs/csv-output.jpg)

#### HTML Output
![HTML-Output](docs/html-output.jpg)

#### Terminal Output
>* GOROOT=C:\Program Files\Go #gosetup
>* GOPATH=C:\Users\mspat\go #gosetup
>* "C:\Program Files\Go\bin\go.exe" build -o C:\Users\mspat\AppData\Local\Temp\GoLand\___go_build_exifreader.exe exifreader #gosetup
>* C:\Users\mspat\AppData\Local\Temp\GoLand\___go_build_exifreader.exe
>* 2023/06/08 13:07:45 Error reading EXIF data from images\exif-error.jpg: EOF 
>* 2023/06/08 13:07:45 Error reading EXIF data from images\more_images\dog.png: exif: failed to find exif intro marker 
>* 2023/06/08 13:07:45 Error reading EXIF data from images\rock.jpg: EOF 
>* CSV file created: ./output.csv 
>* HTML file created: ./output.html 
>* 
>* Process finished with the exit code 0

---

---
#### Time Complexity:
The time complexity of the given code can be analyzed as follows:

* Opening and creating files:  
  The time complexity of opening and creating files is generally considered to be constant or O(1) because it does not depend on the number of files. Therefore, the time complexity for opening the CSV and HTML files is O(1).


* Walking through the directory:   
  The time complexity of walking through the directory and its subdirectories using filepath.Walk is proportional to the number of files and directories in the directory tree. Let's denote the total number of files and directories as n. In the worst case, if there are n files and directories, the time complexity would be O(n).


* Checking if a file is an image:   
  The time complexity of checking if a file is an image by extracting its extension and comparing it with known image extensions is constant or O(1). This operation does not depend on the size of the file.


* Reading EXIF data:   
  The time complexity of reading EXIF data from an image file using exif.Decode depends on the size of the file and the complexity of the EXIF decoding algorithm. Let's denote the size of the image file as m. The time complexity of reading EXIF data can be approximated as O(m), assuming the decoding algorithm has a linear time complexity.


* Extracting GPS data:   
  The time complexity of extracting GPS data from the EXIF data using exifData.LatLong() depends on the complexity of the GPS extraction algorithm. Assuming it has a constant time complexity or O(1), the time complexity of extracting GPS data is also O(1).


* Writing to CSV and HTML files:   
  The time complexity of writing a single row of data to the CSV file and the HTML file is considered constant or O(1). The total time complexity depends on the number of images processed, which is proportional to the number of files in the directory tree, denoted as n.


> Overall, the dominant factor in the time complexity is the number of files and directories in the directory tree (n). Therefore, the overall time complexity can be approximated as O(n) in the worst case scenario.

---

---
#### Space Complexity:

Regarding space complexity, the main considerations are the memory used for storing the directory tree and the EXIF data of the images being processed.   
The space complexity can be analyzed as follows:

* Directory tree:   
  The space complexity for storing the directory tree depends on the depth and branching factor of the tree. In the worst case, if there are n files and directories, the space complexity for storing the directory tree can be approximated as O(n).


* EXIF data:  
  The space complexity for storing the EXIF data of the images being processed depends on the size of the EXIF data and the number of images. Assuming the EXIF data size is proportional to the image file size, let's denote the maximum image file size as m. The space complexity for storing the EXIF data can be approximated as O(m).


* Additional variables:   
  The space complexity for additional variables used in the code is considered constant or O(1). These variables do not depend on the input size.

> Overall, the dominant factor in the space complexity is the size of the largest image file (m). Therefore, the overall space complexity can be approximated as O(m) in the worst case scenario.
---