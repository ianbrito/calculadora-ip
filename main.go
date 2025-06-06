package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func uint32ToIp(ip uint32) string {
	ipParts := make([]string, 4)

	for i := 3; i >= 0; i-- {
		ipParts[i] = strconv.Itoa(int(ip & 0xFF))
		ip >>= 8
	}

	return strings.Join(ipParts, ".")
}

func ipToUint32(ipStr string) uint32 {
	parts := strings.Split(ipStr, ".")
	var ip uint32

	for i := range 4 {
		oct, _ := strconv.Atoi(parts[i])
		ip = ip<<8 | uint32(oct)

	}

	return ip
}

func main() {

	var ipAddressAndMask string

	fmt.Scan(&ipAddressAndMask)

	parts := strings.Split(ipAddressAndMask, "/")

	ipStr := parts[0]
	prefix, _ := strconv.Atoi(parts[1])

	ip := ipToUint32(ipStr)
	broadcast := ^uint32(0) << (32 - prefix)
	network := ip & broadcast
	bits := 32 - prefix
	hosts := int(math.Pow(float64(2), float64(bits)) - 2)

	ipBroadcast := uint32ToIp(broadcast)
	ipNetwokAddress := uint32ToIp(network)

	fmt.Printf("Endereço de IP: %s\n", ipStr)
	fmt.Printf("Mascara de subrede: %d\n", prefix)
	fmt.Printf("Endereço de rede: %s\n", ipNetwokAddress)
	fmt.Printf("Endereço de broadcast: %s\n", ipBroadcast)
	fmt.Printf("Nº hosts: %d\n", hosts)
}
