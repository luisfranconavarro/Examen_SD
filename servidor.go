package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"strconv"
)

var puerto int
var arregloG []string
var msjs []string

func MandarMensaje(s string, x string){
	c, err := net.Dial("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(x)
	if err != nil {
		fmt.Println(err)
	}
}

func server4()  {
	s, err := net.Listen("tcp", ":9995")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleServidorTercero(c)
	}
}

func server1()  {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleServidorPrimero(c)
		c, err = net.Dial("tcp", ":9997")
		if err != nil {
			fmt.Println(err)
			return
		}
		var tamArreglo []string
		var p1 string
		var p2 string
		p1 = ":" + strconv.Itoa(puerto)
		p2 = ":" + strconv.Itoa(puerto-1)
		tamArreglo = append(tamArreglo,p1)
		tamArreglo = append(tamArreglo,p2)
		puerto = puerto - 2
		arregloG = append(arregloG,p1)
		err = gob.NewEncoder(c).Encode(tamArreglo)
		if err != nil {
			fmt.Println(err)
		}
		c.Close()
	}
}

func handleServidorPrimero(c net.Conn)  {
	var s string
	err := gob.NewDecoder(c).Decode(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Se conecto:",s)
}

func server2()  {
	s, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleClient(c)
	}
}

func server3() {
	s, err := net.Listen("tcp", ":9996")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleServidorSegunda(c)
	}
}

func handleServidorSegunda(c net.Conn)  {
	var s string
	err := gob.NewDecoder(c).Decode(&s)
	c, err = net.Dial("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(msjs)
	if err != nil {
		fmt.Println(err)	
	}
}

func handleServidorTercero(c net.Conn)  {
	var arreglo []string
	var auxArreglo[] string
	gob.NewDecoder(c).Decode(&arreglo)
	fmt.Println("Se desconecto",arreglo[0])
	for i := 0; i < len(arregloG); i++ {
		if (arreglo[1] != arregloG[i]){
			auxArreglo = append(auxArreglo,arregloG[i])
		}
	}
	arregloG = auxArreglo
}

func handleClient(c net.Conn)  {
	var proceso string
	err := gob.NewDecoder(c).Decode(&proceso)
	if err != nil {
		fmt.Println(err)
		return
	}
	msjs = append(msjs,proceso)
	fmt.Println(proceso)
	for i := 0; i < len(arregloG); i++ {
		MandarMensaje(arregloG[i],proceso)			
	}
}

func main(){	
	var apagar string
	puerto = 9993

	go server1()
	go server2()
	go server3()
	go server4()
	fmt.Println("Se encendio el servidor\n")

	fmt.Scanln(&apagar)
	fmt.Println("Se apago el servidor\n")
} 