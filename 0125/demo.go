package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"strings"
	"net/http/cookiejar"
)

func main() {
	cookieJar, _ := cookiejar.New(nil)

	for i := 1; i < 200; i++ {
		if 0 == i%2 {
			v := url.Values{}
			/*v.Set("username", "xxxx")
			v.Set("password", "xxxx")*/
			//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
			body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
			urlStr := "http://www.cipm-expo.com/e/cpxx_chs.php?ID=1389&W=http://www.cipm-expo.com/e/cplb_chs.php$ID=8"
			client := &http.Client{
				Jar: cookieJar,
			} //客户端,被Get,Head以及Post使用
			request, err := http.NewRequest("GET", urlStr, body)
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}
			request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
			request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
			request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
			request.Header.Set("Cache-Control", "max-age=0")
			request.Header.Set("Connection", "keep-alive")
			//request.Header.Set("Cookie", cookieStr)
			request.Header.Set("Host", "www.cipm-expo.com") //
			request.Header.Set("Referer", "http://www.cipm-expo.com/e/cpll_chs.php?ID=1389")
			request.Header.Set("Upgrade-Insecure-Requests", "1")
			request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.221 Safari/537.36 SE 2.X MetaSr 1.0") //必须设定该参数,POST参数才能正常提交

			resp, err := client.Do(request) //发送请求
			defer resp.Body.Close()         //一定要关闭resp.Body
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}

			fmt.Println(string(content), '\n', urlStr)

		} else {
			v := url.Values{}
			/*v.Set("username", "xxxx")
			v.Set("password", "xxxx")*/
			//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
			body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
			urlStr := "http://www.cipm-expo.com/e/cplb_chs.php?ID=8"
			client := &http.Client{
				Jar: cookieJar,
			} //客户端,被Get,Head以及Post使用
			request, err := http.NewRequest("GET", urlStr, body)
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}
			//给一个key设定为响应的value.
			request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
			request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
			request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
			request.Header.Set("Cache-Control", "max-age=0")
			request.Header.Set("Connection", "keep-alive")
			//request.Header.Set("Cookie", "PHPSESSID=ejpl6puv8eg10nrtqmccbt93s6")
			request.Header.Set("Host", "www.cipm-expo.com")
			request.Header.Set("Upgrade-Insecure-Requests", "1")
			request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.221 Safari/537.36 SE 2.X MetaSr 1.0") //必须设定该参数,POST参数才能正常提交

			resp, err := client.Do(request) //发送请求
			defer resp.Body.Close()         //一定要关闭resp.Body
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}


			fmt.Println(string(content), '\n', urlStr,"@")
		}

	}

}
