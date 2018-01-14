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
	buildPath := "build"
	configPath := "./config.json"
	messagePath := "message.txt"
	resultPath := "build/result.txt"
	buildDirMode := 0777
	resultFileMode := 0666

	if _, err = os.Stat(buildPath); os.IsNotExist(err) {
		os.Mkdir(buildPath, os.FileMode(buildDirMode))
	}

	if cfg, err = data.FromJson(configPath); err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New(messagePath).ParseFiles(messagePath))

	if f, err = os.OpenFile(resultPath, os.O_CREATE|os.O_WRONLY, os.FileMode(resultFileMode)); err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if err = tmpl.Execute(f, cfg); err != nil {
		log.Fatal(err)
	}
}
