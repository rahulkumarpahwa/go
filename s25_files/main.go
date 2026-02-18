package main

import (
	"fmt"
	"os"
)

// we will learn about the file system in golang. we will read, write and work with the files.

func main() {

	// this will return the file pointer and the error.
	fp, err := os.Open("./example.txt")

	// we will handle the error here.

	if err != nil {
		// here we can either :
		// 1. Log the Error
		fmt.Println(err.Error())
		// 2. Create the Panic over the error.
		panic(err)
	}
	// we can have the multiple methods over the file pointer.
	// similar we have the Stat() method as :
	fileInfo, err := fp.Stat() // this return the file Info or error otherwise.
	if err != nil {
		// here we can either :
		// 1. Log the Error
		fmt.Println(err.Error())
		// 2. Create the Panic over the error.
		panic(err)
	}
	fmt.Println("File Name : ", fileInfo.Name())
	fmt.Println("Is Folder : ", fileInfo.IsDir()) // is it folder?
	fmt.Println("File Size in Bytes : ", fileInfo.Size())
	fmt.Println("File Permissions : ", fileInfo.Mode())
	fmt.Println("File Last Modified at : ", fileInfo.ModTime())

	// 1. now, reading the file:
	fp1, err1 := os.Open("./example.txt")

	if err1 != nil {
		// here we can either :
		// 1. Log the Error
		fmt.Println(err1.Error())
		// 2. Create the Panic over the error.
		panic(err1)
	}

	// to close the file we use the defer.
	defer fp1.Close()

	fileInfo1, err1 := fp1.Stat()

	// we will read the file and store in buffer.
	buf := make([]byte, fileInfo1.Size())

	d, err := fp1.Read(buf) // pass the buffer here.
	if err != nil {
		panic(err)
	}
	fmt.Println("File data : ", string(buf), " \nBuffer Length : ", d)

	// other way to get the data of file, character by character:
	for i := 0; i < len(buf); i++ {
		fmt.Println(i, " : ", string(buf[i]))
	}

	// 2. other way to read the file:
	fp2, err := os.ReadFile("./example.txt")
	// why not to use this method ?
	// because  it will load the whole file at once in memory which in the case of video file may exceed the memory limit.
	if err != nil {
		panic(err)
	}
	fmt.Println("File data : ", string(fp2))
	// so, solution is to get the file in streaming fashion like in the case of node.js.

	// 3. now, first understand how to read folder.
	// dir, err := os.Open(".") // current folder path is '.'
	dir, err := os.Open("../") // one folder back path is '../'
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	// now to get the file info in the directory (folder):
	filesInfo, err := dir.ReadDir(0)
	// here we pass the 'n' which is the count of the number of files and folders you want to read in that in directory / folder.
	// It returns the slice of the file in the fileInfo.
	//If n > 0, ReadDir returns at most n DirEntry records. In this case, if ReadDir returns an empty slice, it will return an error explaining why. At the end of a directory, the error is io.EOF.
	// If n <= 0, ReadDir returns all the DirEntry records remaining in the directory. When it succeeds, it returns a nil error (not io.EOF).

	for _, file := range filesInfo {
		fmt.Print(file.Name())
		fmt.Println(file.IsDir())
	}
	// this is how we read folder.

	// 4. how to write in the file:
	// @24:39

}
