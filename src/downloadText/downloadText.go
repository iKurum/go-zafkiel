package main

import (
	"fmt"
	"io/ioutil"
	"mahonia"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

//网站 笔趣阁
var url = "http://www.biquge.com.tw"

//遮天链接
var textURL = "/0_213/"

func getText(url string, reg string) [][]string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("httpGet error:", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("body err:", err)
	}

	//编码
	dec := mahonia.NewDecoder("GB18030")
	str := dec.ConvertString(string(body))

	//替换 标签br 和 空格
	str = strings.Replace(str, "<br />", "", -1)
	str = strings.Replace(str, "&nbsp;", "", -1)
	// fmt.Println(str)

	exp := regexp.MustCompile(reg)
	r := exp.FindAllStringSubmatch(str, -1)

	return r
}

func main() {
	fmt.Println("=== 开始下载 ===")

	//运行时间
	startTime := time.Now()
	defer func() {
		diffTime := time.Since(startTime)
		fmt.Println("=== 下载完成 耗时", diffTime, "===")
	}()

	//regexp 获取章节链接
	regList := `dd><a href="(.*?)">(.*?)</a></dd>`
	//regexp 获取章节内容
	regCont := `id="content">([\s\S]*?)</div>`

	//获取 列表
	str := getText(url+textURL, regList)

	//创建 txt文件
	f, err := os.Create("遮天.txt")
	if err != nil {
		fmt.Println("create err:", err)
	}

	// 循环列表，获取正文并写入txt
	for _, value := range str {
		text := getText(url+value[1], regCont)
		t := []rune(text[0][0])
		content := string(t[13 : len(t)-6])

		fmt.Println(value[2])

		f.WriteString(value[2])
		f.WriteString("\n")
		f.WriteString(content)
		f.WriteString("\n\n")
	}
}
