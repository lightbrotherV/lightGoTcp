package lightGoTcp

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	once   sync.Once
	server *TcpServer
)

/**
单例模式创建服务器连接
@var host       主机地址
@var port		端口
@var protocol	协议（暂时只支持tcp）
*/
func getServerConn(host string, port int, protocol string) *TcpServer {
	once.Do(func() {
		if protocol == "tcp" {
			pipe := make(MessagePipe)
			server = &TcpServer{
				Port:      port,
				Host:      host,
				StartTime: time.Now().Unix(),
				Pipe:      pipe,
			}
			addr := host + ":" + strconv.Itoa(port)
			if listener, err := net.Listen("tcp", addr); err == nil {
				server.Listener = listener
			} else {
				log.Fatal("Create TcpSocket Failed,", err.Error())
			}
		} else {
			log.Fatal("Only can use Tcp protocol")
		}
	})
	return server
}
