package Server

import (
	"first/devices/TestBench"
	"fmt"
	"log"
	"net"
)

//根据ip地址监听相关信息
func StartListen(ipAddr string) {
	log.Println("Start listening Tcp/Ip from " + ipAddr + "  ...")
	// 创建 listener
	//go func() {
	//	listener, err := net.Listen("tcp", ipAddr)
	//	if err != nil {
	//		fmt.Println("Error listening", err.Error())
	//	}
	//	// 监听并接受来自客户端的连接
	//	for {
	//		conn, err := listener.Accept()
	//		if err != nil {
	//			fmt.Println("Error accepting", err.Error())
	//		}
	//		message := doServerStuff(conn)
	//		fmt.Println(message)
	//	}
	//}()

	listener, err := net.Listen("tcp", ipAddr)
	if err != nil {
		fmt.Println("Error listening", err.Error())
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Println("Error accepting", err.Error())
		}
		message := doServerStuff(conn)
		//log.Println(message)
		ret := TestBench.TestBenchFuncManage(message)
		_, err = conn.Write([]byte(ret))
		if err != nil {
			log.Printf("写入返回值时的连接错误！")
		}
	}
}

func doServerStuff(conn net.Conn) (message string) {
	for {
		buf := make([]byte, 512)
		lenConn, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading", err.Error())
			return "" //终止程序
		}
		log.Printf("Received data: %v", string(buf[:lenConn]))
		return string(buf[:lenConn])
	}
}
