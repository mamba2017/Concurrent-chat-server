package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	re1 := regexp.MustCompile(`<h1 class="post-title">(?s:(.*?))</h1>`)
	if re1 == nil{
		fmt.Printf("regexp.MustCompile err")
		return
	}
	tempTitle := re1.FindAllStringSubmatch(result, -1)
	for _,data := range tempTitle {
		title = data[1]
		break
	}
	//取内容
	re2 := regexp.MustCompile(`<code>(?s:(.*?))</code>`)
	if re2 == nil{
		fmt.Printf("regexp.MustCompile err")
		return
	}
	tempContent := re2.FindAllStringSubmatch(result, -1)
	for _,data := range tempContent {
		content = data[1]
		break
	}
	return
}
func toFile(i int ,fileTitle,fileContent []string){
	f,err := os.Create(strconv.Itoa(i)+".txt")
	if err != nil{
		fmt.Println("create file failed")
		return
	}
	defer f.Close()
	for i:=0;i < len(fileTitle);i++{
		f.WriteString(fileTitle[i]+"\n")
		f.WriteString(fileContent[i]+"\n")
		f.WriteString("==================================================================================\n")
	}
}
func SipderPage(i int){
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
	fileTitle := make([]string,0)
	fileContent := make([]string,0)
	//取网址
	for _,data := range joyUrls{
		//爬取每个段子,网址data[1]
		data[1] = "http://www.duanziwang.com/"+data[1][6:]
		//fmt.Println(data)
		title,content,err := SpiderOneJoy(data[1])
		if err != nil{
			fmt.Println("spideronejoy err",err)
			continue
		}
		content = strings.Replace(content," ","",-1)
		content = strings.Replace(content,"&quot;&quot;","",-1)
		content = strings.Replace(content,"&amp;nbsp;","",-1)
		title = strings.Replace(title,"&amp;nbsp;","",-1)

		//fmt.Println("title:",title)
		//fmt.Println("content:",content)
		fileTitle = append(fileTitle,title)  //追加内容
		fileContent = append(fileContent,content)
	}
	//写入文件
	toFile(i,fileTitle,fileContent)

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