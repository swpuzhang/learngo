package usegorout

import (
	"io"
	"log"
	"net"
	"os"
)

// Netcat  模拟nc
func Netcat(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Fatal(err)
		}
	}()
	if _, err := io.Copy(conn, os.Stdin); err != nil {
		log.Fatal(err)
	}
	return
}
