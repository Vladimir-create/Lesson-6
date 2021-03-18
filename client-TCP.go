package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12668")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	s2 := "{\"action\":\"create\",\"object\":\"Teacher\",\"data\":{\"id\":\"s001\",\"subject\":\"Math\",\"salary\":2345,\"classroom\":[\"CL-001\",\"CL-002\",\"CL-005\"],\"person\":{\"name\":\"Ivan\",\"surname\":\"Popov\",\"personalCode\":\"123422-43235\"}}}"
	conn.Write([]byte(s2))

fmt.Println("fdf")
	//fmt.Println(string(buf[:n]))
}
