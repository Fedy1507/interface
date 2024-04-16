package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name       string
	Age        string
	Occupation string
}

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var user []User
	scanner := bufio.NewScanner(f)
	var human User
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ":")
		switch words[0] {
		case "Name":
			human.Name = words[1]
		case "Age":
			human.Age = words[1]
		case "Occupation":
			human.Occupation = words[1]
			user = append(user, human)
		}
	}
	for _, i := range user {
		fmt.Println(i)
	}
}