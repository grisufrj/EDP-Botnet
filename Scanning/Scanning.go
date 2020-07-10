package main

import (
	"fmt"
	"net"
)

func scan(ip string, porta int) string {
	connection, err := net.Dial("tcp", ip)
	if err != nil {
		return ""
	}

	connection.Close()
	fmt.Printf("Porta %d aberta\n", porta)
	return ip
}

// ValidIps retorna os endereços ips e mac que estão sendo usados na rede
//
// []ips, []macs
//

type interfaceInfo struct {
	networkInterface string
	macAddress       net.HardwareAddr
	mask             []net.IPMask
	validIPs         []string
}

func getValidIps() ([]interfaceInfo, error) {
	var validIps []string
	infos := []interfaceInfo{}
	interfaces, err := net.Interfaces()
	if err != nil {
		return []interfaceInfo{}, err
	}
	var validIPs []string
	for _, i := range interfaces {
		ips, err := i.Addrs()

		if err != nil {
			return []interfaceInfo{}, err
		}
		var masks []net.IPMask
		for _, ip := range ips {
			var ipType net.IP = []byte(ip.String())
			fmt.Printf("%v : %s (%s)\n", i.Name, ip, ipType.DefaultMask())
			masks = append(masks, ipType.DefaultMask())
			validIps = append(validIps, ip.String())
		}

		ncInfo := interfaceInfo{
			networkInterface: i.Name,
			macAddress:       i.HardwareAddr,
			mask:             masks,
			validIPs:         validIPs,
		}
		infos = append(infos, ncInfo)
	}
	return infos, nil
}

func main() {
	var ipValidos []string
	iporta := "127.0.0.1:3306"
	porta := 3306
	result := scan(iporta, porta)
	if result == "" {
		fmt.Println("Fudeu")
	}
	ipValidos = append(ipValidos, result)
	ips, err := getValidIps()
	if err != nil {
		panic(err)
	}
	fmt.Println(ips)
}
