package main

import (
	"net"
	"fmt"
	"flag"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var (
	targetip = net.ParseIP(flag.String("t", "", "Target IP address to yoink packets"))
	iface = net.InterfaceByName(flag.String("i", "eth0", "Interface to listen and send malicious packets through."))
)

func main(
	flag.Parse()


)