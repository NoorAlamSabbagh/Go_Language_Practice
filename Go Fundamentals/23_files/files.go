package main

// func main() {
// 	f, err := os.Open("example.txt") //f -->file
// 	if err != nil {
// 		// log the error
// 		panic(err)
// 	}
// 	fileInfo, err := f.Stat()
// 	if err != nil {
// 		// log the error
// 		panic(err)
// 	}
// 	fmt.Println("file name:", fileInfo.Name())
// 	fmt.Println("file or folder:", fileInfo.IsDir())
// 	fmt.Println("file size:", fileInfo.Size())
// 	fmt.Println("file permission:", fileInfo.Mode())
// 	fmt.Println("file modified at:", fileInfo.ModTime())
// }

//
// func main() {
// 	// read file
// 	f, err := os.Open("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	buf := make([]byte, 12)
// 	d, err := f.Read(buf)
// 	if err != nil && err != io.EOF {
// 		panic(err)
// 	}
// 	for i := 0; i < d; i++ {
// 		fmt.Println("data:", string(buf[i]))
// 	}
// 	fmt.Println("Total bytes read:", d)
// 	fmt.Println("Full content:", string(buf[:d]))
// }

//
// func main() {
// 	data, err := os.ReadFile("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(data))
// }

//
// func main() {
// 	// read folders
// 	// dir, err := os.Open("./")
// 	dir, err := os.Open("../")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer dir.Close()
// 	// fileInfo, err := dir.ReadDir(1)
// 	// fileInfo, err := dir.ReadDir(2)
// 	fileInfo, err := dir.ReadDir(-1) //fo all file
// 	for _, fi := range fileInfo {
// 		// fmt.Println(fi.Name())
// 		fmt.Println(fi.Name(), fi.IsDir())
// 	}
// }

//
// func main() {
// 	// create a file
// 	f, err := os.Create("example2.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	f.WriteString("hi go")
// 	f.WriteString("nice language")
// 	// read and write to another file (streaming fashion)
// }

//
// func main() {
// 	// create a file
// 	f, err := os.Create("example2.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	bytes := []byte("Hello Golang")
// 	f.Write(bytes)
// 	// read and write to another file (streaming fashion)
// }

//
// func main() {
// 	// read and write to another file (streaming fashion)
// 	sourceFile, err := os.Open("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer sourceFile.Close()

// 	destFile, err := os.Create("example2.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer destFile.Close()
// 	reader := bufio.NewReader(sourceFile)
// 	writer := bufio.NewWriter(destFile)
// 	for {
// 		b, err := reader.ReadByte()
// 		if err != nil {
// 			if err.Error() != "EOF" {
// 				panic(err)
// 			}
// 			break
// 		}
// 		e := writer.WriteByte(b)
// 		if e != nil {
// 			panic(e)
// 		}
// 	}
// 	writer.Flush()
// 	fmt.Println("written to new file succesfully")
// }

//
// func main() {
// 	// delete a file
// 	err := os.Remove("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("file deleted successfully")
// }
