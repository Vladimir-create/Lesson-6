package main

import (
	"net"
	"fmt"
	//"io"
	"encoding/json"
)

type (
	Action struct {
		Action string `json:"action"`
		ObjName string `json:"object"`
	}
	Teacher struct {
		localMutex chan int
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
	Teachermutex struct {
		Ch chan int
	}
	UpdateTeacher struct {
		T Teacher `json:"data"`
	}
	CreateTeacher struct {
		T Teacher `json:"data"`
	}
	DeleteTeacher struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	ReadTeacher struct {
		T Teacher
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
		GetCreateAction() DefinedAction
		GetUpdateAction() DefinedAction
		GetReadAction() DefinedAction
		GetDeleteAction() DefinedAction
		Read(str string)bool
		Print()
	}
)

func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}

func (action ReadTeacher) Process() {
	fmt.Println("Read teacher", action.Data.ID)
	for i:=0;i<len(arriPerson);i++{
		if arriPerson[i].Read(action.Data.ID) {
			<- action.T.localMutex 
			arriPerson[i].Print()
		}
	}
	action.T.localMutex  <- 1
}

func (action CreateTeacher) Process(){
		fmt.Println("proceess")
	action.T.localMutex <- 1 
	fmt.Println("Create Teacher")
	arriPerson = append(arriPerson, &action.T)
	PrintAll(arriPerson)
	fmt.Println("proceess2")
}

func (action UpdateTeacher) Process() {
	fmt.Println("Update Teachers")
	for i:=0;i<len(arriPerson);i++{	
		if arriPerson[i].Read(action.T.ID) {
			<- action.T.localMutex
			arriPerson[i] = &action.T
		}
	}
	PrintAll(arriPerson)
	action.T.localMutex <- 1
}

func (action DeleteTeacher) Process() {
	fmt.Println("Teacher deleted", action.Data.ID)
	for i:=0;i<len(arriPerson);i++{
		if arriPerson[i].Read(action.Data.ID) {
			<- teachermutex.Ch
			copy(arriPerson[i:], arriPerson[i+1:])
			arriPerson[len(arriPerson)-1] = nil
			arriPerson = arriPerson[:len(arriPerson)-1]
		}
	}
	PrintAll(arriPerson)
	teachermutex.Ch <- 1
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

func (action *DeleteTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *UpdateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *CreateTeacher) GetFromJSON (rawData []byte) {
		fmt.Println("getfrjson")
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
		fmt.Println("getfrjson2")
}
	
func PrintAll(arriPerson []GeneralObject){
	for i:=0; i<len(arriPerson); i++{
		arriPerson[i].Print()
	}	
}

var teachermutex Teachermutex
func main() {
	l, err := net.Listen("tcp", "127.0.0.1:12668")
	if err != nil {
		panic(err)
	}
	defer l.Close()
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
	buf := make([]byte, 0, 10000) 
    smallBuf := make([]byte, 256)   
      
    for {
        n,_ := conn.Read(smallBuf)
        if n<256 {
			buf = append(buf, smallBuf[:n]...)
            break
        }
        buf = append(buf, smallBuf[:n]...)
    }
    fmt.Println(string(buf))
	var act Action
	var obj GeneralObject
	var toDo DefinedAction
	fmt.Println(string(buf[:len(buf)]))
	err := json.Unmarshal(buf, &act)
	fmt.Println(act)
	if err != nil {
		fmt.Println("error")
	}
	switch act.ObjName {
	case "Teacher":
		obj = &Teacher{}
	}
	switch act.Action {
		case "create":
			fmt.Println("create")
			toDo = obj.GetCreateAction()
			fmt.Println("create2")
		case "update":
			toDo = obj.GetUpdateAction()
		case "read":
			toDo = obj.GetReadAction()
		case "delete":
			toDo = obj.GetDeleteAction()
	}
	fmt.Println()
	toDo.GetFromJSON(buf)
	toDo.Process()
	data := []byte("Connection great")
	conn.Write(data)
	conn.Close()
}
