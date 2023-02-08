package client

import (
	"bufio"
	"hua-proxy/file"
	"hua-proxy/utils"
	"io"
	"log"
	"net"
	"strconv"
)

var (
	// 本地需要暴露的服务端口
	localIp   = file.GetEnvParam().LocalIp
	localPort = file.GetEnvParam().LocalPort
	remoteIP  = file.GetEnvParam().RemoteIp
	// 远端的服务控制通道，用来传递控制信息，如出现新连接和心跳
)

// Main 客户端主启动方法
func Main() {
	addr := " HOST: " + remoteIP + ", PORT: " + strconv.Itoa(utils.ControlPort)
	tcpConn, err := utils.Connect(remoteIP, utils.ControlPort)
	if err != nil {
		log.Println("[连接失败]" + addr + err.Error())
		return
	}
	log.Println("[已连接]" + addr)

	reader := bufio.NewReader(tcpConn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		// 当有新连接信号出现时，新建一个tcp连接
		if s == utils.NewConnection+"\n" {
			go connectLocalAndRemote()
		}
	}

	log.Println("[已断开]" + addr)
}

// 创建本地和远端链接通道
func connectLocalAndRemote() {
	local := connectLocal()
	remote := connectRemote()

	if local != nil && remote != nil {
		utils.Join2Conn(local, remote)
	} else {
		if local != nil {
			_ = local.Close()
		}
		if remote != nil {
			_ = remote.Close()
		}
	}
}

func connectLocal() *net.TCPConn {
	conn, err := utils.Connect(localIp, localPort)
	if err != nil {
		log.Println("[连接本地服务失败]" + err.Error())
	}
	return conn
}

func connectRemote() *net.TCPConn {
	conn, err := utils.Connect(remoteIP, utils.TunnelPort)
	if err != nil {
		log.Println("[连接远端服务失败]" + err.Error())
	}
	return conn
}
