package main

import (
	"os"
	"log"
	"html/template"
	"./topicBlock"
)

func main() {
	argsWithoutProg := os.Args[1:]
	
	pathWithCodeForDocumentation:= argsWithoutProg[0]
	if pathWithCodeForDocumentation == ""{
		pathWithCodeForDocumentation = "./"
	}

	pathForFileDocumentation := argsWithoutProg[1]
	if pathForFileDocumentation == ""{
		pathForFileDocumentation = "./documentation.html"
	}

	docNavigator := topicBlock.CodeDocumentation{}
	docNavigator.GetDocumentationFromFilesPath(pathWithCodeForDocumentation)
	
	masterTmpl, err := template.ParseFiles("./templates/master.html")
	if err != nil {
		log.Fatal(err)
	}
	
	f, err := os.Create(pathForFileDocumentation)
	defer f.Close()
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	if err := masterTmpl.Execute(f, docNavigator.DocumentationBlocks); err != nil {
		log.Fatal(err)
	}
	
}
