package topicBlock

import "regexp"


var methodBlockSearch = regexp.MustCompile(`(//@\b(router)(.*)(\n)*)?(//@\b(method)(.*)(\n)*)(//@\b(router)(.*)(\n)*)?(//@\b(response)(.*)(\n)*)?(//@\b(request)(.*)(\n)*)?(//@\b(response)(.*)(\n)*)?(//@\b(description)(.*)(\n)*)?`)
var methodRegexp = regexp.MustCompile(`//@\b(method)+\[+([\s]*)+(\b(name=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?(([\s]*)+(\,)?([\s]*)+(type=)+([\s]*)+\"+\b[0-9A-Za-z]+\"+([\s]*))?\]`)
var routerRegexp = regexp.MustCompile(`//@\b(router)\[([\s]*)\b(path=)([\s]*)\"[^\s]*\"([\s]*)?\]`)
var requestRegexp = regexp.MustCompile(`//@\b(request)\[[\s]*\b(params=)(.*)\]`)
var responseRegexp = regexp.MustCompile(`//@\b(response)\[[\s]*\b(answer=)(.*)\]`)
var descriptionRegexp = regexp.MustCompile(`//@\b(description)\[(.*)\]`)
var methodNameRegexp =   regexp.MustCompile(`(//@method\[(name=)\")?(\")?(\,)?((type=)+\"+\b[0-9A-Za-z]+\")?(\])?`)
var methodTypeRegexp = regexp.MustCompile(`(//@method\[(name=)+\"+\b[0-9A-Za-z]+\")?(\,)?(type=\")?(\")?(\])?`)
var routerName = regexp.MustCompile(`(//@\b(router)+\[+([\s]*)+\b(path=)+([\s]*)+\")?(\")?(\])?`)
var requestParams =  regexp.MustCompile(`(//@\b(request)\[[\s]*\b(params=)[\s]*)?([\s]*\]$)?`)
var responseAnswer = regexp.MustCompile(`(//@\b(response)\[[\s]*\b(answer=)[\s]*)?([\s]*\]$)?`)
var descriptionText = regexp.MustCompile(`(//@\b(description)\[)?(\])?`)

type MethodBlock struct{
	MethodType string
	MethodName string
	Router string
	RequestParams string
	ResponseParams string
	Description string
}

func getBlockParamByContent(content string, searchString *regexp.Regexp, searchParam *regexp.Regexp) string{
	wantedString :=  searchString.FindString(content)
	return searchParam.ReplaceAllString(wantedString, "")
}

func (mb *MethodBlock) GetMethodBlockFromString(content string){
	mb.MethodName = getBlockParamByContent(content, methodRegexp, methodNameRegexp) 
	mb.MethodType = getBlockParamByContent(content, methodRegexp, methodTypeRegexp)
	mb.Router = getBlockParamByContent(content, routerRegexp, routerName)
	mb.RequestParams =  getBlockParamByContent(content, requestRegexp, requestParams)
	mb.ResponseParams = getBlockParamByContent(content, responseRegexp, responseAnswer)
	mb.Description =  getBlockParamByContent(content, descriptionRegexp, descriptionText)
}

type TopicBlock struct{
	TopicName string
	MethodsBlocks []MethodBlock
}


func (tb *TopicBlock) GetBlocksFromContent(content string, topicName string){
	var blocksIndexes = methodBlockSearch.FindAllStringIndex(content, -1)
	if len(blocksIndexes)>0{
		tb.TopicName = topicName
		tb.MethodsBlocks = []MethodBlock{}
		var docBlock string
		for _ , blockIndex := range blocksIndexes{
			docBlock = content[blockIndex[0]:blockIndex[1]]
			methodBlock:= MethodBlock{}
			methodBlock.GetMethodBlockFromString(docBlock)
			tb.MethodsBlocks = append(tb.MethodsBlocks, methodBlock)
		} 
	} 
}
