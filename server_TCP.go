package main

import (
	"net"
	"fmt"
	"encoding/json"
)

type (
	Action struct {
		Action string `json:"action"`
		ObjName string `json:"object"`
	}
	Teacher struct {
		Ch chan int
		ID string  `json:"id"`
		Salary float64 `json:"salary"`
		Subject string `json:"subject"`
		Classroom []string `json:"classroom"`
		Person struct {
			Name string `json:"name"`
			Surname string `json:"surname"`
			PersonalCode string `json:"personalCode"`
		} `json:"person"`
	}
	ReadTeacher struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
)

type (
	DefinedAction interface {
		GetFromJSON([]byte)
		Process()
	}
	GeneralObject interface {
		GetReadAction() DefinedAction
		Read(str string)bool
		Print()
	}
)
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}

func (action ReadTeacher) Process() {
	fmt.Println("Read teacher", action.Data.ID)
	for i:=0;i<len(arriPerson);i++{
		if arriPerson[i].Read(action.Data.ID) {
			arriPerson[i].Print()
		}
	}
}

func (i *Teacher) Print(){
	fmt.Println("Id:", i.ID)
	fmt.Println("Salary:", i.Salary)
	fmt.Println("Subject:", i.Subject)
	for j:=0; j<len(i.Classroom);j++{
		fmt.Println("Classroom",j+1, "=", i.Classroom[j])
	}
	fmt.Println("Name:", i.Person.Name)
	fmt.Println("Surname:", i.Person.Surname)
	fmt.Println("PersonalCode:", i.Person.PersonalCode)
	fmt.Println()
}


func (i *Teacher) Read(str string) bool{
	return i.ID == str
}

func (action *ReadTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}
	
var teacher Teacher 
func main() {
	l, err := net.Listen("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	teacher.Ch = make(chan int, 1)
	teacher.Ch <- 1

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go HandleConnection(conn)
	}
}

var arriPerson []GeneralObject

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return
	}
	var act Action
	var obj GeneralObject
	var toDo DefinedAction
	
	<- teacher.Ch
	
	err = json.Unmarshal(buf[:n], &act)
	if err != nil {
		fmt.Println("error")
	}
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	}
	switch act.Action {
	case "read":
		toDo = obj.GetReadAction()
	}
	toDo.GetFromJSON(buf[:n])
	toDo.Process()
	//fmt.Println(string(buf[:n]))
	//dat.N = string(buf[:n])

	data := []byte("Connection great")
	conn.Write(data)

	teacher.Ch <- 1

	conn.Close()
}
