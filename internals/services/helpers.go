package services

import (
	"bytes"
	"fmt"
	"fr/diginamic/go-cap/internals/config"
	"os"
	"strings"
	"text/template"
)

const (
	DTO = iota
	JSON_CONTROLLER
	HTML_CONTROLLER
	SERVICE
	ENTITY
	REPOSITORY
	MAPPER
)

var checkDirectoryExists map[int]func() = map[int]func(){}
var CONF = config.CONFIG
var SRC = config.CONFIG.SourceDirectory
var SUFFIXES = map[string]string{
	CONF.Mapper:     "Mapper.php",
	CONF.Service:    "Service.php",
	CONF.Dto:        "Dto.php",
	CONF.Entity:     ".php",
	CONF.Repository: "Repository.php",
}

func init() {
	checkDirectoryExists[DTO] = func() {
		if _, err := os.Stat(SRC + CONF.Dto); err != nil {
			askCustomDir("Dto", func(dir string) {
				CONF.Dto = dir
			})
		}
	}
	checkDirectoryExists[MAPPER] = func() {
		if _, err := os.Stat(SRC + CONF.Mapper); err != nil {
			askCustomDir("Mapper", func(dir string) {
				CONF.Mapper = dir
			})
		}
	}
	checkDirectoryExists[ENTITY] = func() {
		if _, err := os.Stat(SRC + CONF.Repository); err != nil {
			askCustomDir("Entity", func(dir string) {
				CONF.Entity = dir
			})
		}
	}
	checkDirectoryExists[REPOSITORY] = func() {
		if _, err := os.Stat(SRC + CONF.Repository); err != nil {
			askCustomDir("Repository", func(dir string) {
				CONF.Repository = dir
			})
		}
	}
	checkDirectoryExists[SERVICE] = func() {
		if _, err := os.Stat(SRC + CONF.Service); err != nil {
			askCustomDir("Service", func(dir string) {
				CONF.Service = dir
			})
		}
	}
	checkDirectoryExists[HTML_CONTROLLER] = func() {
		if _, err := os.Stat(SRC + CONF.HtmlController); err != nil {
			askCustomDir("Service", func(dir string) {
				CONF.HtmlController = dir
			})
		}
	}
	checkDirectoryExists[JSON_CONTROLLER] = func() {
		if _, err := os.Stat(SRC + CONF.JsonController); err != nil {
			askCustomDir("Service", func(dir string) {
				CONF.JsonController = dir
			})
		}
	}
}

// MUTATES CONFIG
func askCustomDir(subject string, executable func(dir string)) {
	//mut
	var dir string
	fmt.Printf("Couldn't find %s directory as written in config. Please provide your custom %s directory from the './src/' directory\n", subject, subject)
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
	var choice string

	for FileExists(path) &&
		strings.ToLower(choice) != "y" &&
		strings.ToLower(choice) != "n" {
		fmt.Println("File already exists. Do you wish to override it ? (y/n)")
		fmt.Scanln(&choice)
	}

	if strings.ToLower(choice) == "n" {
		os.Exit(0)
	}

	f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		panic(err)
	}

    if content != nil {
        f.Write(content.Bytes())
    }
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func getTmplBytes(mainClass, tmplUrl string, data any) *bytes.Buffer {
	//mut
	var tplBytes bytes.Buffer

	fileContent, err := os.ReadFile(config.RESOURCE_FOLDER + tmplUrl)
	tmpl, err := template.New(mainClass).Parse(string(fileContent))
	if err != nil {
		panic(err)
	}
	//mutable
	err = tmpl.Execute(&tplBytes, data)
	if err != nil {
		panic(err)
	}

	return &tplBytes
}

func checkDependencies(pathsToCheck []string, entityName string) {

	for _, path := range pathsToCheck {
		_, err := os.Stat(SRC + path + entityName + SUFFIXES[path])

		if err != nil {
			fmt.Println(SRC + path + entityName + SUFFIXES[path])
			fmt.Printf("required dependency %s wasn't found. Make sure it exists before creating the service\n", entityName+path)
			os.Exit(1)
		}
	}
}
