package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func makeALotOfFiles(source []byte, count int) (names []string) {
	fmt.Println("\n### Generating files ###\n")
	start := time.Now()
	for i := 1; i <= count; i++ {
		ind := strconv.Itoa(i)
		file, _ := os.Create("files/file" + ind)
		file.WriteString(string(source[:rand.Intn(len(source))]))
		defer file.Close()
		fmt.Println("Making file #" + ind)
		names = append(names, "files/file"+ind)
	}
	fmt.Println(time.Since(start))
	return names
}

func makeZip(name string, files []string) {
	/*
		buf := new(bytes.Buffer)
		w := zip.NewWriter(buf)
		defer w.Close()
		for i := 1; i <= 10; i++ {
			ind := strconv.Itoa(i)
			body, _ := ioutil.ReadFile("files/file" + ind)
			f, _ := w.Create("files/file" + ind)
			f.Write([]byte(body))
			fmt.Println("Adding file # " + ind + " to archive")
		}
	*/
	fmt.Println("\n### Adding files to ZIP ###\n")
	start := time.Now()
	newfile, _ := os.Create(name)
	defer newfile.Close()
	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	for i, file := range files {
		zipfile, _ := os.Open(file)
		defer zipfile.Close()
		info, _ := zipfile.Stat()
		header, _ := zip.FileInfoHeader(info)
		header.Method = zip.Deflate
		writer, _ := zipWriter.CreateHeader(header)
		io.Copy(writer, zipfile)
		fmt.Printf("Adding file %d "+file+"\n", i+1)
	}
	fmt.Println("\n`myfile.zip` created at files/myfiles.zip")
	fmt.Println(time.Since(start))
}

func main() {
	tolstoy, _ := ioutil.ReadFile("warandpeace")
	files := makeALotOfFiles(tolstoy, 20)
	makeZip("files/myfiles.zip", files)

}
