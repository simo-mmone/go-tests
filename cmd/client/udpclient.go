package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func getServerIp() string {

	//get public ip
	ip, _ := net.LookupIP("https://simo-mmone-animated-space-fiesta-wqqwqgj76j42gxvg-1234.preview.app.github.dev")
	return ip[0].String()
}

func main() {
	p := make([]byte, 2048)

	fmt.Println(getServerIp())

	conn, err := net.Dial("udp", "52.184.222.68:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	start := time.Now()
	_, err = bufio.NewReader(conn).Read(p)
	fmt.Println(time.Since(start))
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
