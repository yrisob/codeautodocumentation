package topicBlock

import "time"
import "fmt"
import "io/ioutil"
import "log"
import "strings"

type CodeDocumentation struct{
	DateCreate time.Time
	DocumentationBlocks []TopicBlock
}

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
				fileList = append(fileList, ReadAllFilesFromDir(fmt.Sprintf("%s/%s",dirName, file.Name()))...)
			}
		} else {
			fileList = append(fileList, fmt.Sprintf("%s/%s",dirName, file.Name()))
		}
	}

	return fileList
}

func ReadDocumentationFromFile(fileName string, tb chan *TopicBlock){
	contentbytes, error := ioutil.ReadFile(fileName)

	if error != nil{
		log.Fatal(error)
	}else{
		contentString:= string(contentbytes)
		lastIndexPoint:= strings.LastIndex(fileName, ".")
		if lastIndexPoint <= 0{
			lastIndexPoint = len(fileName)
		}
		topicBlock:= TopicBlock{}
		topicBlock.GetBlocksFromContent(contentString, fileName[strings.LastIndex(fileName,"/")+1: lastIndexPoint])
		tb <- &topicBlock
	}

}

func (cd *CodeDocumentation) GetDocumentationFromFilesPath(rootPath string){
	cd.DateCreate = time.Now()
	projectFiles := ReadAllFilesFromDir(rootPath)
	cd.DocumentationBlocks = []TopicBlock{}
	topicBlock := make(chan *TopicBlock)

	for _, fileName := range projectFiles{
		go ReadDocumentationFromFile(fileName, topicBlock)
	}

	for i:=0; i<len(projectFiles);i++{
		tb:= <-topicBlock
		if tb.TopicName != ""{
			cd.DocumentationBlocks = append(cd.DocumentationBlocks, *tb)
		}
	}

}