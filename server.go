package lightGoTcp

import (
	"log"
	"net"
)

type MessagePipe chan interface{}

type TcpServer struct {
	Port        int
	Host        string
	StartTime   int64
	Pipe        MessagePipe
	Listener    net.Listener
	ClientQueue []net.Conn
}

/**
获取tcp连接
如果没初始化，则通过单例，初始化
@var args[0] host
@var args[1] port
@var args[2] protocol
*/
func (ts *TcpServer) GetConn(args ...interface{}) net.Listener {
	if ts.Listener == nil {
		if len(args) <= 3 {
			log.Fatal("getConn Error")
		}
		ts = getServerConn(args[0].(string), args[1].(int), args[2].(string))
	}
	return ts.Listener
}

/**
循环处理连接请求
@var handle 处理函数
需要在handle中加锁
*/
func (ts *TcpServer) Run(handle func(net.Conn)) {
	for {
		conn, err := ts.GetConn().Accept()
		if err == nil {
			log.Fatal("Client Connect Error")
		}
		ts.ClientQueue = append(ts.ClientQueue, conn)
		go handle(conn)
	}
}

/**
循环处理客户端连接
@var handle 处理函数
需要在handle中加锁
*/
func (ts *TcpServer) HandleClient(handle func(net.Conn)) {
	for {
		for _, val := range ts.ClientQueue {
			go handle(val)
		}
	}
}
