package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Fields must be exported
type T struct {
	Data  string
	Data2 int
	Addr  *T
}

func makedb(name string) {

	file, _ := os.Create(name)
	defer file.Close()

	t := T{
		Data:  "hello",
		Data2: 17,
		Addr:  nil,
	}
	t2 := T{
		Data:  "world",
		Data2: 42,
		Addr:  &t,
	}
	enc := gob.NewEncoder(file)
	enc.Encode(&t2)
}

func readdb(name string) {
	file, _ := os.Open(name)
	defer file.Close()
	tt := T{}

	dec := gob.NewDecoder(file)
	file.Seek(0, 0)

	dec.Decode(&tt)

	tt2 := tt.Addr
	fmt.Println(tt.Data, tt.Data2)
	fmt.Println(tt2.Data, tt2.Data2)
}

func main() {
	makedb("mydatabase")
	readdb("mydatabase")
}

/*
import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Str struct {
	Name string
	File []byte
}

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
	buf := new(bytes.Buffer)
	var STR Str
	STR.Name = "tolstoy"
	STR.File = tolstoy

	binary.Write(buf, binary.LittleEndian, STR)

	ioutil.WriteFile("tolst", buf.Bytes(), 0777)
	var q []byte
	binary.Read(buf, binary.LittleEndian, q)
	fmt.Print(q)

}
*/
