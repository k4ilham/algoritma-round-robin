package main

import (
	"fmt"
	"time"

	"github.com/chi-chu/round"
)

func main() {
	r := round.NewRound()
	r.WithHeartBeatTime(10 * time.Second)
	r.WithKeepAlive(true)
	r.WithCheckAlive(func(IP string, Port int) bool {
		return Port%7 != 0
	})
	r.AddServer(&round.Server{Name: "1", Weight: 3, CurrentWeight: 0, IP: "111:111:111", Port: 7, Alive: true})
	r.AddServer(&round.Server{Name: "2", Weight: 3, CurrentWeight: 0, IP: "222:222:222", Port: 2, Alive: true})
	r.AddServer(&round.Server{Name: "3", Weight: 2, CurrentWeight: 0, IP: "333:333:333", Port: 3, Alive: true})
	r.AddServer(&round.Server{Name: "4", Weight: 1, CurrentWeight: 0, IP: "444:444:444", Port: 4, Alive: true})
	r.AddServer(&round.Server{Name: "5", Weight: 1, CurrentWeight: 0, IP: "555:555:555", Port: 5, Alive: true})
	r.Start()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		s, err := r.GetServer()
		if err != nil {
			panic(err)
		}
		fmt.Printf("_______________________\n")
		fmt.Printf("User %v served by Server: %s \n", i+1, s.Name)
		fmt.Printf("_______________________\n")

		// Menampilkan CurrentWeight dari masing-masing server
		for _, server := range r.ServerList {
			fmt.Printf("Server %s CurrentWeight: %d\n", server.Name, server.CurrentWeight)
		}
	}
}
