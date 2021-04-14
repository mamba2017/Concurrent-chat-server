package main

import (
	"fmt"
	"net"
)

func Input(conn net.Conn){
	var chatContent string
	for{
		//fmt.Println("请输入聊天内容")
		fmt.Scan(&chatContent)
		conn.Write([]byte(chatContent))
	}
}

func main() {
	var addr string
	addr = "192.168.217.130"   //测试地址
	//fmt.Scan(&addr)
	addr = addr+":8000"
	conn,err1 := net.Dial("tcp",addr)
	if err1 != nil{
		fmt.Println("net.dial err1 = ",err1)
		return
	}
	defer  conn.Close()
	go func() { //接受服务器数据
		buf := make([]byte,2048)
		for{
			n,err:= conn.Read(buf)
			if n == 0{  //对方断开，出问题
				fmt.Println("conn read err",err)
				return
			}else{
				fmt.Println(string(buf[:n-1]))
			}
		}

	}()

	Input(conn)  //输入聊天内容


}
