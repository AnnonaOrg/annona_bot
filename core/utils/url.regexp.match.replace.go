package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

const (
	urlReg = `((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`
)

func UrlRegMatchReplace(str string) (ret string) {
	ret = str
	compile := regexp.MustCompile(urlReg)
	urlMap := make(map[string]string)
	submatch := compile.FindAllSubmatch([]byte(ret), -1)
	for _, m := range submatch {
		url := string(m[0])
		urlMap[url] = fmt.Sprintf("<a href=\"%s\">%s</a>", url, url)
	}
	// fmt.Println(urlMap)
	for k, v := range urlMap {
		ret = strings.ReplaceAll(ret, k, v)
	}
	// fmt.Println(ret)
	return
}

func UrlRegMatchReplaceToMarkdown(str string) (ret string) {
	ret = str
	compile := regexp.MustCompile(urlReg)
	urlMap := make(map[string]string)
	submatch := compile.FindAllSubmatch([]byte(ret), -1)
	for _, m := range submatch {
		url := string(m[0])
		urlMap[url] = fmt.Sprintf("[%s](%s)", url, url)
	}
	// fmt.Println(urlMap)
	for k, v := range urlMap {
		ret = strings.ReplaceAll(ret, k, v)
	}
	// fmt.Println(ret)
	return
}

// 识别文案中链接，并转义，输出为TG支持的HTML格式文案
func UrlRegMatchReplaceToTGHTML(str string) (ret string, err error) {
	ret = str
	tplParseText := str
	compile := regexp.MustCompile(urlReg)
	urlMap := make(map[string]string)
	submatch := compile.FindAllSubmatch([]byte(ret), -1)
	for index, m := range submatch {
		url := string(m[0])
		urlKey := "url" + fmt.Sprintf("%d", index)
		// fmt.Println("index", index, "url", url, "urlKey", urlKey)
		urlMap[urlKey] = url
		tplParseText = strings.ReplaceAll(tplParseText, url, fmt.Sprintf("<a href=\"{{ .%s }}\">{{ .%s }}</a>", urlKey, urlKey))
	}
	// fmt.Println("tplParseText", tplParseText)
	var result bytes.Buffer
	tpl, err := template.New("message").Parse(tplParseText)
	if err != nil {
		return ret, fmt.Errorf("template parsing failed: %w", err)
	}
	if err := tpl.Execute(&result, urlMap); err != nil {
		return ret, fmt.Errorf("template execution failed: %w", err)
	}
	retHtml := result.String()
	return retHtml, nil
}

// hanMap := make(map[string]string, 0)
// hanMap["汉//\\字"] = "http://han.com"
// hanMap["汉.df||字2"] = "http://han2.com"
// ret, err := util.UrlMapToTGHTML(hanMap)
//
//	if err != nil {
//		fmt.Println("出错了", err)
//	} else {
//
//		fmt.Println("ret", ret)
//	}
//
// urlMap[title]url，转义，输出为TG支持的HTML格式文案
func UrlMapToTGHTML(urlMap map[string]string) (ret string, err error) {
	ret = ""
	tplParseText := ""
	urlKeyMap := make(map[string]string)
	for k, v := range urlMap {
		vc := v
		kc := k
		kHash := "k" + MD5Hash(vc)
		urlKeyMap[kHash] = kc
	}
	for k, v := range urlKeyMap {
		vc := v
		kc := k
		tplParseText = tplParseText + fmt.Sprintf("<a href=\"%s\">{{ .%s }}</a>\n", urlMap[vc], kc)
	}
	tplParseText = strings.TrimSpace(tplParseText)
	// fmt.Println("tplParseText", tplParseText)
	var result bytes.Buffer
	tpl, err := template.New("message").Parse(tplParseText)
	if err != nil {
		return ret, fmt.Errorf("template parsing failed: %w", err)
	}
	if err := tpl.Execute(&result, urlKeyMap); err != nil {
		return ret, fmt.Errorf("template execution failed: %w", err)
	}
	retHtml := result.String()
	return retHtml, nil
}
