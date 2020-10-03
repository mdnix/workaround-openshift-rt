package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func scheduled() {
	for {
		<-time.After(30 * time.Minute)
		go flush()
	}
}

func flush() {
	f, err := os.OpenFile("/proc/sys/net/ipv4/route/flush", os.O_WRONLY, 0400)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString("1"); err != nil {
		log.Println("Unable to flush route table cache")
	} else {
		log.Println("Flushed ip route cache")
	}
}

func main() {
	flush()
	go scheduled()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c

	log.Println("Caught a signal. Shutting down")
}
