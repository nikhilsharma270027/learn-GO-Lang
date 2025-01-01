package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// logging err
	// 	// panic in go means abb luch nahi ho rsakta
	// 	panic(err)
	// }
	// // fmt.Println(f)

	// fileInfo, err := f.Stat() // return file info or return error if their is error in file path
	// if err != nil {
	// 	// logging err
	// 	// panic in go means abb luch nahi ho sakta
	// 	panic(err)
	// }
	// fmt.Println("file Name", fileInfo.Name())         // will display file name
	// fmt.Println("file or folder: ", fileInfo.IsDir()) // file or folder false
	// fmt.Println("file size: ", fileInfo.Size())       //file size 0
	// fmt.Println("file size: ", fileInfo.Mode())       //file size -rw-rw-rw-
	// fmt.Println("file size: ", fileInfo.ModTime())    //file size 2025-01-01 20:23:12.973 +0530 IST

	// // read file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// logging err
	// 	// panic in go means abb luch nahi ho sakta
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 10)

	// d, err := f.Read(buf) // im reading the file and putting it in buf/
	// if err != nil {
	// 	// logging err
	// 	// panic in go means abb luch nahi ho sakta
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// Same as above ----------------------------

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	// logging err
	// 	// panic in go means abb luch nahi ho sakta
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// this method if the file is big it can cause more resouce constraits at omnce on memory

	// //---------------- streaming method  -------------//
	// dir, err := os.Open("../")
	// if err != nil {
	// 	//logging err
	// 	// panic in go means abb luch nahi ho sakta
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1) // if 1 one file is display
	// 								//if less than -0 all file will be displayed
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())

	// }

	// os.Create("exmaple2.txt")
	// f, err := os.Create("exmaple2.txt")
	// if err != nil {
	// 	// logging err
	// 	// panic in go means abb luch nahi ho sakta
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi golang")

	// bytes := []byte("heelo golang")
	// f.Write(bytes)
	// f.Write(bytes)

	// // read and write to another file(streaming fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// distFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer distFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(distFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("Written a new file successfully")

	// delete a file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("filedeleted")

	// defer sourceFile.Close()

}
