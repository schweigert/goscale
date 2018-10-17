package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"

	"github.com/schweigert/lampper"
)

func main() {
	bind := os.Getenv("BIND")
	pullSize := os.Getenv("PULLSIZE")

	size, err := strconv.Atoi(pullSize)
	if err != nil {
		panic(err)
	}

	log.Println("starting on:", bind)
	for i := 0; i < size; i++ {
		log.Println("EP", i, "->", getAddr(i))
	}

	service, err := lampper.Listen("tcp", bind)
	if err != nil {
		panic(err)
	}

	for {
		peer := service.Accept()
		random := rand.Intn(size)
		addr := getAddr(random)
		log.Println("Routing to EP", random, "->", addr)
		go handle(addr, peer)
	}
}

func getAddr(el int) string {
	elStr := strconv.Itoa(el)
	return os.Getenv("EP_" + elStr)
}

func handle(addr string, peer1 *lampper.Peer) {
	log.Println("Dial to:", addr)
	peer2, err := lampper.Dial("tcp", addr)
	if err != nil {
		log.Println("Error to dial:", addr)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go pipe(wg, peer1, peer2)
	go pipe(wg, peer2, peer1)

	wg.Wait()
	log.Println("Handle closed to:", addr)
}

func pipe(wg *sync.WaitGroup, peer1, peer2 *lampper.Peer) {
	defer wg.Done()

	for {
		bytes, err := peer1.ReadBytes(1)
		if err != nil {
			return
		}

		err = peer2.WriteBytes(bytes)
		if err != nil {
			return
		}
	}
}
