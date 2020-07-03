package main

import (
	"fmt"
	"log"

	"github.com/caucy/batch_ping"
)

func main() {
	ipSlice := []string{}
	// ip list should not more than 65535

	ipSlice = append(ipSlice, "192.168.1.1") //support ipv6
	ipSlice = append(ipSlice, "192.168.1.2")
	ipSlice = append(ipSlice, "192.168.1.3")
	ipSlice = append(ipSlice, "192.168.1.110")
	ipSlice = append(ipSlice, "192.168.1.111")
	ipSlice = append(ipSlice, "192.168.1.112")
	ipSlice = append(ipSlice, "192.168.1.113")
	ipSlice = append(ipSlice, "192.168.1.114")
	ipSlice = append(ipSlice, "192.168.1.115")
	ipSlice = append(ipSlice, "192.168.1.116")
	ipSlice = append(ipSlice, "192.168.1.117")
	ipSlice = append(ipSlice, "192.168.1.118")
	ipSlice = append(ipSlice, "192.168.1.119")
	fmt.Println(ipSlice)
	bp, err := ping.NewBatchPinger(ipSlice, true) // true will need to be root

	if err != nil {
		log.Fatalf("new batch ping err %v", err)
	}
	bp.SetDebug(false) // debug == true will fmt debug log

	bp.SetSource("") // if hava multi source ip, can use one isp
	bp.SetCount(2)

	bp.OnFinish = func(stMap map[string]*ping.Statistics) {
		for _, st := range stMap {
			if st.PacketsRecv == 0 {
				fmt.Println(st.Addr," connect bad")
			}else{
				fmt.Println(st.Addr," connect success")

			}
		}

	}

	err = bp.Run()
	if err != nil {
		log.Printf("run err %v \n", err)
	}
	bp.OnFinish(bp.Statistics())
}
