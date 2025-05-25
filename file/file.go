package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// writing to a file in go
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	content := "This is actually working"
	bytee, errors := io.WriteString(file, content)
	fmt.Println("Number of bytes written:", bytee)
	if errors != nil {
		fmt.Println("Error writing in the file:", errors)
		return
	}

	fmt.Println("File created successfully")

	// reading from a file in go using buffer
	/*
		file, err = os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file:", err)
			return
		}

		buffer := make([]byte, 1024)

		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file:", err)
				return
			}
			fmt.Println(string(buffer[:n]))
		}

	*/

	// reading from a file in go using ReadFile
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	fmt.Println("Data read from file:", string(data))

}
