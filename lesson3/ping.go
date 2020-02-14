package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"os"
)


func main(){


	if len(os.Args) == 2 {
		var dns string
		dns = os.Args[1]
		//pinger, err := ping.NewPinger("www.google.com")
		pinger, err := ping.NewPinger(dns)
		if err != nil {
			panic(err)
		}

		pinger.Count = 3
		pinger.Run() // blocks until finished
		stats := pinger.Statistics()
		fmt.Println(stats.Addr, stats.IPAddr, stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	} else {
		panic("Error")
	}
}
