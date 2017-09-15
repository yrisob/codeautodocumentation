package main

import (
	"fmt"
	"./topicBlock"
)


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


func main() {
	docNavigator := topicBlock.CodeDocumentation{}
	docNavigator.GetDocumentationFromFilesPath("../go_test")
	fmt.Println(docNavigator)
}
