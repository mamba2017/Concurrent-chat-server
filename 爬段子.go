package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func HttpGet(url string)(result string,err error)  {
	resp,err1 := http.Get(url)
	if err !=nil{
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte,1024*4)
	for{
		n,err := resp.Body.Read(buf)
		if n==0{   //读取结束或出问题
			fmt.Println("read ere=",err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func SpiderOneJoy(url string)(title ,content string,err error){
	result,err1 := HttpGet(url)
	if err1 !=nil{
		fmt.Println("httpget err",err)
		err = err1
		return
	}
	//取标题
	re1 := regexp.MustCompile(`<h1 class="post-title"><a href="(?s:(.*?))"`)
	if re1 == nil{
		fmt.Printf("regexp.MustCompile err")
		return
	}
}

func SipderPage(i int){
	fmt.Println("11111")
	url := "http://www.duanziwang.com/page/"+ strconv.Itoa(i) +"/index.html"
	//开始爬取内容
	result,err := HttpGet(url)
	if err !=nil{
		fmt.Println("httpget err",err)
		return
	}
	//取内容 <h1 class="post-title"><a href="     "
	//解释表达式
	re := regexp.MustCompile(`<h1 class="post-title"><a href="(?s:(.*?))"`)
	if re == nil{
		fmt.Printf("regexp.MustCompile err")
		return
	}
	joyUrls := re.FindAllStringSubmatch(result, -1)
	//取网址
	for _,data := range joyUrls{
		//爬取每个段子,网址data[1]
		data[1] = "http://www.duanziwang.com/"+data[1][6:]
		//fmt.Println(data[1])
		title,content,err := SpiderOneJoy()
		if err != nil{
			fmt.Println("spideronejoy err",err)
			continue
		}
		fmt.Println("title:",title)
		fmt.Println("content:",content)
	}

}
func DoWork(start,end int){
	fmt.Printf("正在爬取 %d 到 %d 的页面\n",start,end)
	for i:=start;i <=end;i++{
        //定义函数爬主页面
		SipderPage(i)
	}
}
func main() {
	var start,end int
	fmt.Scan(&start)
	fmt.Scan(&end)
	DoWork(start,end)
}
