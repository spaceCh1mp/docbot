package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const tokenPath string = "./token.json"

type Token struct {
	AccessToken string `json:"name"`
	AccessPath  string `json:"password"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var dir string
	fileDir := flag.String("dir", "./log/", "Program defaults to ./log/ when no file path is specified")
	fileName := flag.String("name", "", "No value is set for the name flag and must be specified")
	rel := flag.Bool("rel", true, "file path is relative by default, set this flag to false to set a path from your home directory")
	flag.Parse()
	if !(*rel) {
		dir = validate(*fileDir) + *fileName
	} else {
		dir = *fileDir + *fileName
	}
	filePath := Path(dir, *rel)
	//open log file
	file, err := ioutil.ReadFile(filePath)
	check(err)
	//read file content to variable in format
	msg := formatText(string(file))
	fmt.Println(len(msg))
	/* Day {day}
	- Accomplished task
	- Another Accomplished task
	*/
	//get twitter access token from json file
	//parse json file and assign values to the Token struct
	tokenBytes, err := ioutil.ReadFile(tokenPath)
	check(err)
	var t Token
	json.Unmarshal(tokenBytes, &t)
	//pass values to http request and hit api endpoint

}

func Path(dir string, dirType bool) string {
	if !dirType {
		home, _ := os.UserHomeDir()
		return home + dir
	}
	return dir
}

func validate(dir string) string {
	pref, suff := strings.HasPrefix(dir, "/"), strings.HasSuffix(dir, "/")
	if !(pref) || !(suff) {
		if !(pref) {
			dir = "/" + dir
		}
		if !(suff) {
			dir += "/"
		}
	}
	return dir
}

func formatText(msg string) string {
	return msg
}
