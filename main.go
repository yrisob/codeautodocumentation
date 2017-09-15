package main

import (
	"os"
	"log"
	"html/template"
	"./topicBlock"
)

func main() {
	docNavigator := topicBlock.CodeDocumentation{}
	docNavigator.GetDocumentationFromFilesPath("../go_test")
	
	masterTmpl, err := template.ParseFiles("./templates/master.html")
	if err != nil {
		log.Fatal(err)
	}
	
	f, err := os.Create("./documentation.html")
	defer f.Close()
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	if err := masterTmpl.Execute(f, docNavigator.DocumentationBlocks); err != nil {
		log.Fatal(err)
	}
	
}
