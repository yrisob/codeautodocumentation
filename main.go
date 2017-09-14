package main

import (
	"regexp"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"./topicBlock"
)

type attribute string

// ReadAllFilesFromDir получение всех вложенных файлов директории
func ReadAllFilesFromDir(dirName string) []string {
	fmt.Println("Read dir:", dirName)
	var fileList []string
	files, error := ioutil.ReadDir(dirName)

	if error != nil {
		log.Fatal(error)
	}

	for _, file := range files {
		if file.IsDir() {
			if strings.Index(file.Name(), ".") != 0 {
				fileList = append(fileList, ReadAllFilesFromDir(fmt.Sprintf("%s/%s", dirName, file.Name()))...)
			}
		} else {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList
}

var someText=`some text
//@router[path="/people/{id}"]
//@method[name="GetPeople",type="GET"]
//@request[params= {id: 10} ]
//@response[answer={id:1, name: "someName", age:12, gender: "M"}]
//@description[где id- идентификатор персоны, name - имя персоны, age- возраст персоны, gende- гендер]
some footertext`

//@router[path="/people/{id}"]
//@method[name="GetPeople",type="GET"]
//@response[answer={id:1, name: "someName", age:12, gender: "M"}]
//@description[где id- идентификатор персоны, name - имя персоны, age- возраст персоны, gende- гендер]

var regSearchAttr = regexp.MustCompile(`(//@\b(router)(.*)(\n)*)?(//@\b(method)(.*)(\n)*)(//@\b(router)(.*)(\n)*)?(//@\b(response)(.*)(\n)*)?(//@\b(request)(.*)(\n)*)?(//@\b(response)(.*)(\n)*)?(//@\b(description)(.*)(\n)*)?`)

func main() {
	var stringSlices = regSearchAttr.FindAllStringIndex(someText,1)
	fmt.Println(stringSlices)

	tb:= topicBlock.TopicBlock{}
	tb.GetBlocksFromContent(someText,"test")

	fmt.Println(tb)

	// fmt.Println("start read directory")
	// files := ReadAllFilesFromDir("../")
	// fmt.Println(files)

}
