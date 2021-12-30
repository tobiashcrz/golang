package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	parallelizer "github.com/shomali11/parallelizer"
	fastping "github.com/tatsushid/go-fastping"
)

func main() {

	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		for i := 0; i < 50; i++ {
			fmt.Printf("%c ", i)

			p := fastping.NewPinger()

			ip := getIp()

			ra, err := net.ResolveIPAddr("ip", ip)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			p.AddIPAddr(ra)

			p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
				fmt.Printf("\nIP Addr: %s receive - OK", addr.String())
			}
			p.MaxRTT = time.Duration(1) * time.Millisecond
			erro := p.Run()
			if erro != nil {
				log.Printf("error occur: %v\n", erro)
			}
			p.OnIdle = func() {
				fmt.Println("\nfinish")
			}

		}
	})
	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)

}

func getIp() string {
	min := 100
	max := 200

	return fmt.Sprintf("%d%s%d%s%d%s%d", rand.Intn(max-min)+min, ".", rand.Intn(max-min)+min, ".", rand.Intn(max-min)+min, ".", rand.Intn(max-min)+min)
}
