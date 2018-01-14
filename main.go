package main

import (
	"./data"
	"log"
	"os"
	"text/template"
)

type Role struct {
	name   string
	duties string
}

type User struct {
	role Role
	name string
}

type Company struct {
	name  string
	goals string
}

type Open struct {
	role Role
}

func main() {
	cfg := data.NewMock()

	tmpl := template.Must(template.New("message.txt").ParseFiles("message.txt"))

	f, err := os.OpenFile("build/result.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = tmpl.Execute(f, cfg)

	if err != nil {
		log.Fatal(err)
	}
}
