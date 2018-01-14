package main

import (
	"./data"
	"log"
	"os"
	"text/template"
)

func main() {
	var cfg data.Config
	var f *os.File
	var err error

	if _, err := os.Stat("build"); os.IsNotExist(err) {
		os.Mkdir("build", os.FileMode(0777))
	}

	cfg, err = data.FromJson("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New("message.txt").ParseFiles("message.txt"))

	f, err = os.OpenFile("build/result.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = tmpl.Execute(f, cfg)

	if err != nil {
		log.Fatal(err)
	}
}
