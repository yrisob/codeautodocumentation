package main

import (
	"strings"
	"regexp"
)

var methodRegexp = regexp.MustCompile(`//@\b(method)+\[+([\s]*)+(\b(name=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?(([\s]*)+(\,)?([\s]*)+(type=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?\]+$`)
var routerRegexp = regexp.MustCompile(`//@\b(router)+\[\b(name=)+\"(\W*\w*)*\"\]+$`)
var requestRegexp = regexp.MustCompile(`//@\b(request)+\[+\b(params=)+\"+\[?(\{?((\w*\:[^\s]*)\,?)*\}?\,?)*\]?\"+\]+$`)
var responseRegexp = regexp.MustCompile(`//@\b(response)+\[+\b(params=)+\"+\[?(\{?((\w*\:[^\s]*)\,?)*\}?\,?)*\]?\"+\]+$`)
var descriptionRegexp = regexp.MustCompile(`//@\b(description)+\[(.*)*\]+$`)

type MethodBlock struct{
	MethodType string
	MethodName string
	RequestParams string
	ResponseParams string
	Description string
}

type TopicBlock struct{
	MethodType string
	MethodsBlocks []MethodBlock
}

// func getMethodAttributes(methodDeclaration string)(methodName string, methodType string){
// 	strWithoutSpaces = strings.Replace(methodDeclaration," ","", -1)

// } 

func (tb TopicBlock) GetBlocksFromContent(content string){
	var methodsBlocks = []MethodBlock{}
	var methodPositions = methodRegexp.FindAllStringIndex(content, -1)
	for i:=0; i<len(methodPositions); i++{
		var startPositionBlock, endPositionBlock int
		startPositionBlock = methodPositions[i][0]
		if i < len(methodPositions)-1{
			endPositionBlock = methodPositions[i+1][0]
		}else{
			endPositionBlock = len(methodPositions)
		}
		var methodTextBlock = content[startPositionBlock: endPositionBlock]
		MethodBlock

	}

}
