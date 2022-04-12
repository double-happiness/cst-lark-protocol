package utils

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"log"
	"net"
	"strconv"
	"strings"
)

func GetRegisterServerConf(serverAddrs []string) (serverConfig []constant.ServerConfig) {
	for _, serverAddr := range serverAddrs {
		ip, port, err := GetIpAndPort(serverAddr)
		if err != nil {
			continue
		}
		tmp := constant.ServerConfig{
			IpAddr: ip,
			Port:   port,
		}
		serverConfig = append(serverConfig, tmp)
	}
	return
}

func GetIpAndPort(addr string) (string, uint64, error) {
	addrs := strings.Split(addr, ":")
	if len(addrs) != 2 {
		return "", 0, errors.New("invalid addr：" + addr)
	}
	i_port, err := strconv.Atoi(addrs[1])
	if err != nil {
		log.Println("GetIpAndPort Atoi err：", err)
		return "", 0, err
	}
	return addrs[0], uint64(i_port), nil
}

func ExternalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip.String()
		}
	}

	return ""
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
