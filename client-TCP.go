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

	s := "{\"action\":\"delete\",\"object\":\"Teacher\",\"data\":{\"id\":\"001\"}}"
	conn.Write([]byte(s)
	s2 := "{\"action\":\"create\",\"object\":\"Teacher\",\"data\":{\"id\":\"001\",\"subject\":\"Math\",\"salary\":2345,\"classroom\":[\"CL-001\",\"CL-002\",\"CL-005\"],\"person\":{\"name\":\"Ivan\",\"surname\":\"Popov\",\"personalCode\":\"123422-43235\"}}}"
	conn.Write([]byte(s2)
	s3 := "{\"action\":\"create\",\"object\":\"Teacher\",\"data\":{\"id\":\"0ds01\",\"subject\":\"Math\",\"salary\":2345,\"classroom\":[\"CL-001\",\"CL-002\",\"CL-005\"],\"person\":{\"name\":\"Ivan\",\"surname\":\"Popov\",\"personalCode\":\"123422-43235\"}}}"
	conn.Write([]byte(s3)
	s4 := "{\"action\":\"read\",\"object\":\"Teacher\",\"data\":{\"id\":\"001\"}}"
	conn.Write([]byte(s4)
	s5 := "{\"action\":\"update\",\"object\":\"Teacher\",\"data\":{\"id\":\"001\",\"subject\":\"Math\",\"salary\":2345,\"classroom\":[\"CL-001\",\"CL-002\",\"CL-005\"],\"person\":{\"name\":\"Ivan\",\"surname\":\"nigga\",\"personalCode\":\"123422-43235\"}}}"
	conn.Write([]byte(s5)
	s6 := "{\"action\":\"delete\",\"object\":\"Teacher\",\"data\":{\"id\":\"001\"}}"
	conn.Write([]byte(s6)
	
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}
