package services

import (
	"bytes"
	"fmt"
	"fr/diginamic/go-cap/internals/config"
	"os"
	"text/template"
)

//MUTATES CONFIG
func askCustomDir(subject string, executable func (dir string)) {
    //mut
	var dir string
	fmt.Printf("Couldn't find %s directory as written in config. Please provide your custom %s directory from the './src/' directory\n",subject, subject)
    //mutation
	fmt.Scanln(&dir)
	if dir[len(dir)-1] != '/' {
		dir = dir + "/"
	}
	executable(dir)

    //Overrides the config file if you are using it
    if !config.CONFIG.Transient {
        config.WriteConfigFile()
    }
}

func createFile(content *bytes.Buffer, path string) {
    f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    f.Write(content.Bytes())
}

func getTmpleBytes(mainClass, suffix string, data any) *bytes.Buffer {
    //mut
    var tplBytes bytes.Buffer

    fileContent, err := os.ReadFile(config.RESOURCE_FOLDER + "Dto.php");
    tmpl, err := template.New(mainClass).Parse(string(fileContent))
    if err != nil { panic(err) }
    //mutable
    err = tmpl.Execute(&tplBytes, data)
    if err != nil { panic(err) }

    return &tplBytes
}
