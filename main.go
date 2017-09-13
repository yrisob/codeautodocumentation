package main

import (
	"regexp"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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

func main() {

//	var methodRegexp = regexp.MustCompile(`//@\b(method)+\[+([\s]*)+(\b(name=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?(([\s]*)+(\,)?([\s]*)+(type=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?\]+$`)
	var methodAttribute = "//@method[name=\"GetPeople\",type=\"GET\"]"

	re := regexp.MustCompile(`(//@method\[(name=)\")?(\")?(\,)?((type=)+\"+\b[0-9A-Za-z]+\")?(\])?`)
	reType := regexp.MustCompile(`(//@method\[(name=)+\"+\b[0-9A-Za-z]+\")?(\,)?(type=\")?(\")?(\])?`)
	methodName := re.ReplaceAllString(methodAttribute,"")
	methodType := reType.ReplaceAllString(methodAttribute,"")

	fmt.Println(methodAttribute)
	fmt.Println(methodName)
	fmt.Println(methodType)

	fmt.Println("start read directory")
	files := ReadAllFilesFromDir("../")
	fmt.Println(files)

}
