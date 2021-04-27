//package main
//
//import "fmt"
//
//type TreeNode struct {
//	    Val int
//	    Left *TreeNode
//	    Right *TreeNode
//}
//var i = -1
//func createByArr(a []int)*TreeNode{
//	i += 1
//	//fmt.Println("len(a)", len(a))
//	if i >= len(a){
//		return nil
//		}
//	t := new(TreeNode)
//	if a[i] != 0{
//		t.Val = a[i]
//		t.Left = nil
//		t.Right = nil
//		t.Left = createByArr(a)
//		t.Right = createByArr(a)
//		}else{
//		return nil
//		}
//	return t
//}
//var nums []int
//func inorder(root *TreeNode){
//	if root == nil{
//		return
//	}
//
//	inorder(root.Left)
//	nums = append(nums,root.Val)
//	inorder(root.Right)
//}
//func main() {
//	a := []int{10,5,3,1,0,0,0,7,6,0,0,0,15,13,0,0,18,0,0}
//	root := createByArr(a)
//	inorder(root)
//	fmt.Println(nums)
//	count := 0
//	low := 6
//	high := 10
//	for _,value := range nums{
//		if value >= low && value <= high{
//			count += value
//		}
//	}
//	fmt.Println(count)
//}
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)
type Client struct{
	C chan string //用户发送数据的管道
	Name string  //用户名
	Addr string //网络地址
}
//	保存在线用户
var onlineMap map[string]Client

var message = make(chan string)

func Manager()  {
	onlineMap = make(map[string]Client)
	for{
		msg := <- message//没有消息就阻塞
		for _,cli:=range onlineMap{
			cli.C <-msg
		}
	}
}
func WriteMsgToClient(cli Client,conn net.Conn){
	for msg := range cli.C{
		//给当前客户端发信息
		conn.Write([]byte(msg+"\n"))
	}
}

func MakeMsg(cli Client,msg string)(buf string){
	//显示用户ip
	//clientip := strings.Split(cli.Addr,":")[0]
	//显示当前时间
	timeStr:=time.Now().Format("2006-01-02 15:04:09")
	//获取当前时间并转化为字符串类型，
	//2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	buf = "["+timeStr+"]"+cli.Name+":"+msg
	return

}
func ChangeName(cli *Client,name string,conn net.Conn){
	for _,temp := range onlineMap{
		if cli.Addr == temp.Addr && name == temp.Name{ //如果与上一个名字一样
			conn.Write([]byte("与原昵称相同，请换一个！\n"))
			return
		}else if cli.Addr != temp.Addr && name == temp.Name{  //如果与其他人昵称重复
			conn.Write([]byte("昵称已存在，请换一个！\n"))
			return
		}
	} 														//修改成功
	cli.Name = name
	onlineMap[cli.Addr] = *cli
	conn.Write([]byte("修改昵称成功！！！\n"))
	msg := MakeMsg(*cli,"")
	records(cli.Addr,msg + "修改昵称为：" + name)
}



func HandleConn(conn net.Conn){  //处理用户连接
	//获取客户端网络地址
	defer conn.Close()
	cliAddr := conn.RemoteAddr().String()
	//创建结构体
	cli := Client{make(chan string),"小废物",cliAddr}
	//把结构体添加到map
	onlineMap[cliAddr] = cli

	//新开协程，给当前客户端转发消息
	go WriteMsgToClient(cli,conn)
	//广播莫个人在线
	//message <- "["+cliAddr+"]"+cli.Name+":login"
	message <- MakeMsg(cli,"上线了～")
	records(cli.Addr,MakeMsg(cli,"上线了～"))
	//提示我是谁
	cli.C<-MakeMsg(cli,"I am here")
	logMsg := fmt.Sprintf("当前在线人数共[%d]人，用户列表：\n",len(onlineMap))
	conn.Write([]byte(logMsg))
	for _,temp:=range onlineMap{
		msg := "-------"+temp.Name+"-------"+"\n"+"\n"
		conn.Write([]byte(msg))
	}
	isQuit := make(chan bool) //对方是否是主动退出
	hasData := make(chan bool)//对方是否有数据发送

	//新建协程，接受用户发送过来的数据
	go func() {
		buf := make([]byte,2048)
		for{
			n,err:= conn.Read(buf)
			if n == 0{  //对方断开，出问题
				isQuit <- true
				fmt.Println("conn read err",err)
				return
			}
			msg := string(buf[:n])  //nc测试多一个换行

			if len(msg) == 4 && msg == "-who"{
				//遍历map，发送给当前用户
				conn.Write([]byte("user list:\n"))
				for _,temp:=range onlineMap{
					msg := temp.Name+"\n"
					conn.Write([]byte(msg))

				}
			}else if len(msg) >= 9 && msg[:7]== "-rename"{
				//-rename|ekko
				name:=strings.Split(msg,"|")[1]

				ChangeName(&cli,name,conn)  //修改昵称
			}else{
				//转发内容
				message <- MakeMsg(cli,msg)
				records(cli.Addr,MakeMsg(cli,msg))

			}
			hasData <- true
		}
	}()
	for{
		//通过select 检测channel的流动
		select {
		case <- isQuit:
			delete(onlineMap,cliAddr) //当前用户从map移除
			message <- MakeMsg(cli,"退出群聊～")//广播谁下线了
			return
		case <-hasData:

			// case <-time.After(600*time.Second):
			// 	delete(onlineMap,cliAddr)
			// 	message<-MakeMsg(cli,"time out")
			// 	return
		}
	}

}
func  records(addr ,msg string){
	f,err := os.OpenFile("records.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil{
		fmt.Println("open failed",err)
		return
	}
	fmt.Fprintln(f, addr,msg)
	f.Close()
}
func main() {
	//监听


	listenner,err := net.Listen("tcp",":8000")
	if err != nil{
		fmt.Println("listen err",err)
		return
	}
	defer listenner.Close()

	//新开协程转发消息
	go Manager()

	//主协程，循环阻塞等待用户连接
	for{
		conn,err := listenner.Accept()
		if err != nil{
			fmt.Println("listener accept err",err)
			continue
		}
		go HandleConn(conn) //处理用户连接
	}
}

