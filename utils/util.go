package utils

import (
	"io"
	"log"
	"net"
	"strconv"
)

const (
	KeepAlive     = "KEEP_ALIVE"
	NewConnection = "NEW_CONNECTION"
	ControlPort   = 1099
	TunnelPort    = 1098
	VisitPort     = 9999
)

// Listen 监听本地方法
func Listen(port int) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:" + strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListener, nil
}

// Connect 链接远端网络服务
func Connect(host string, port int) (*net.TCPConn, error) {
	addr := host + ":" + strconv.Itoa(port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListener, nil
}

// Join2Conn 复制两个通道的内容
func Join2Conn(local *net.TCPConn, remote *net.TCPConn) {
	go joinConn(local, remote)
	go joinConn(remote, local)
}

func joinConn(local *net.TCPConn, remote *net.TCPConn) {
	defer local.Close()
	defer remote.Close()
	_, err := io.Copy(local, remote)
	if err != nil {
		log.Println("copy failed ", err.Error())
		return
	}
}
