package main

import (
	"strings"
	"regexp"
)

var methodRegexp = regexp.MustCompile(`//@\b(method)+\[+([\s]*)+(\b(name=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?(([\s]*)+(\,)?([\s]*)+(type=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?\]+$`)
var routerRegexp = regexp.MustCompile(`//@\b(router)+\[+([\s]*)+\b(name=)+([\s]*)+\"(\W*\w*)*\"\+([\s]*)+]+$`)
var requestRegexp = regexp.MustCompile(`//@\b(request)+\[+\b(params=)+\"+\[?(\{?((\w*\:[^\s]*)\,?)*\}?\,?)*\]?\"+\]+$`)
var responseRegexp = regexp.MustCompile(`//@\b(response)+\[+\b(params=)+\"+\[?(\{?((\w*\:[^\s]*)\,?)*\}?\,?)*\]?\"+\]+$`)
var descriptionRegexp = regexp.MustCompile(`//@\b(description)+\[(.*)*\]+$`)
var methodNameRegexp =   regexp.MustCompile(`(//@method\[(name=)\")?(\")?(\,)?((type=)+\"+\b[0-9A-Za-z]+\")?(\])?`)
var methodTypeRegexp = regexp.MustCompile(`(//@method\[(name=)+\"+\b[0-9A-Za-z]+\")?(\,)?(type=\")?(\")?(\])?`)
var routerName = regexp.MustCompile(`(//@\b(router)+\[+([\s]*)+\b(name=)+([\s]*)+\")?(\")?(\])?`)
var 

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



func getMethodNameByMethodAttr(attrString string)(string){
	return methodNameRegexp.ReplaceAllString(attrString,"")
}

func getMethodTypeByMethodAttr(attrString string)(string){
	return methodTypeRegexp.ReplaceAllString(attrString,"")
}

func getRouterName(attrString string)(string){
	return routerName.ReplaceAllString(attrString,"")
} 


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

	}

}
