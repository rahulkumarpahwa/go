package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// A. opening the file

	// fp, err := os.Open("../example.txt")
	// if err != nil {
	// 	fmt.Printf("ERROR : opening file:  %v\n", err)
	// }

	// // we can put the different types of method over the file pointer fp
	// fileInfo, err := fp.Stat()
	// if err != nil {
	// 	fmt.Printf("ERROR : stat: %v\n", err)
	// }

	// fmt.Println(fileInfo.Name())    // File Name
	// fmt.Println(fileInfo.Size())    // File Size in BYTES (one BYTE per character)
	// fmt.Println(fileInfo.IsDir())   // IS File A Directory?
	// fmt.Println(fileInfo.Mode())    // File permissions
	// fmt.Println(fileInfo.ModTime()) // last modification time
	// fp.Close()

	// --------------------------------------

	// B. reading the file:
	// fp, err := os.Open("../example.txt")
	// if err != nil {
	// 	fmt.Printf("Error : stat: %v\n", err)
	// }
	// defer fp.Close()

	// // we normally store the data read from the file in the buffer (temp allocation)
	// buf := make([]byte, 12)
	// d, err := fp.Read(buf)

	// // Reading the file in string from the Buffer:
	// for i := 0; i < len(buf); i++ {
	// 	fmt.Printf("%s", string(buf[i]))
	// }

	// if err != nil {
	// 	fmt.Printf("ERROR : Read File: %v\n", err)
	// }

	// fmt.Println("\n", d, "", buf)

	// ----------------------------------

	// C. Other way to read the file:
	// data, err := os.ReadFile("../example.txt") // this reads the whole data at once.
	// if err != nil {
	// 	fmt.Printf("ERROR : Read File: %v\n", err)

	// }
	// fmt.Println(string(data))
	/*
		why should n't we use the ReadFile() ?
		Because it loads the whole data from the file at once in the memory and in case of large data we will get our memory full with data, even if we need just a chunk of it.
	*/

	// so, to handle the large files, we will use the streaming fashion as in case of node.js

	// -----------------------------------

	// D. Read Folder / directory

	// dir, err := os.Open("../")
	// if err != nil {
	// 	fmt.Printf("ERROR : Open Dir: %v\n", err)
	// }

	// defer dir.Close()

	// filesInfo, err := dir.ReadDir(0) // n>0, then n files, when n <= 0, then all the files and this gives the slice of the fileInfo
	// if err != nil {
	// 	fmt.Printf("ERROR : Read Dir: %v\n", err)
	// }

	// for _, fi := range filesInfo {
	// 	fmt.Println(fi.Name(),":", fi.IsDir())
	// }

	// -----------------------------------

	// E. Create and Write a file

	// f, err := os.Create("../example2.txt")
	// if err != nil {
	// 	fmt.Printf("ERROR : Create File: %v\n", err)
	// 	return
	// }
	// // Wrting over the file
	// f.WriteString("Hi Go!")
	// f.WriteString("This is the next line!") // this will appended in file without previous being deleted.

	// f.Truncate(0) // remove all content
	// f.Seek(0, 0)  // moves the cursor to the start

	// f.WriteString("this is the new line!")

	// // writing in byte instead of string

	// f.Write([]byte("this is text converted to byte and then written to the file."))

	// ----------------------------------

	// F. Read from one File and Write in other file
	// we will use the stream and then put the data intead of putting everything in memory and then print.

	// sourceFile, err := os.Open("../example.txt")
	// if err != nil {
	// 	fmt.Printf("ERROR : Open Source File: %v\n", err)
	// 	return
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.text")
	// if err != nil {
	// 	fmt.Printf("ERROR : creating dest File: %v\n", err)
	// 	return
	// }
	// defer destFile.Close()

	// // we will stream, using the buffer inbuild named as bufio package

	// reader := bufio.NewReader(sourceFile)

	// writer := bufio.NewWriter(destFile)

	// // we will get the data in reader and write in writer.
	// // and the buffer has the size of 4096

	// // infinite loop
	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			fmt.Printf("ERROR : Read Buffer Byte: %v\n", err)
	// 			return
	// 		}
	// 		break // incase of the EOF error, otherwise we will stuct in infinite loop
	// 	}

	// 	err = writer.WriteByte(b)
	// 	if err != nil {
	// 		fmt.Printf("ERROR : Write Buffer Byte: %v\n", err)
	// 		return
	// 	}
	// }
	// writer.Flush() // to empty the writer if something is left in it, to clean the buffer.
	// this is streaming fashion to pass the data from one file to another.

	// Also, if we need just to copy one file into another then we can use the copy function on the file itself.

	// ----------------------------------------

	// G. Delete a file

	// delete file
	err := os.Remove("./deletedFile.go")
	if err != nil {
		fmt.Printf("ERROR : delete file : %v\n", err)
		return
	}

	// -----------------------------------------
}
