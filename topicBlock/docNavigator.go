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

func (cd *CodeDocumentation) GetDocumentationFromFilesPath(rootPath string){
	cd.DateCreate = time.Now()
	projectFiles := ReadAllFilesFromDir(rootPath)
	cd.DocumentationBlocks = []TopicBlock{}

	for _, fileName := range projectFiles{
		contentbytes, error := ioutil.ReadFile(fileName)

		if error != nil{
			log.Fatal(error)
		}

		contentString := string(contentbytes)
		topicBlock:= TopicBlock{}
		lastIndexPoint:= strings.LastIndex(fileName, ".")
		if lastIndexPoint <= 0{
			lastIndexPoint = len(fileName)
		}
		topicBlock.GetBlocksFromContent(contentString, fileName[strings.LastIndex(fileName,"/")+1: lastIndexPoint])

		if topicBlock.TopicName != ""{
			cd.DocumentationBlocks = append(cd.DocumentationBlocks, topicBlock)
		}
	}

}