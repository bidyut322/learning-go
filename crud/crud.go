package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequst() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code:", res.StatusCode)
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println("Response data:", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Todo item:", todo)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Bidyut Maji",
		Completed: true,
	}

	// Convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Convert json data to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response : ", string(data))
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    23974,
		Title:     "Updated Title",
		Completed: true,
	}

	// Convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	// Convert json data to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response:", string(data))
	fmt.Println("Response Status Code:", res.StatusCode)

}

func performDeleteRequest() {
	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status Code:", res.StatusCode)
}

func main() {
	//performGetRequst()
	//performPostRequest()
	//performUpdateRequest()
	performDeleteRequest()
}
