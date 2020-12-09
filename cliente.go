package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"bufio"
	"os"
)

var Puerto_1, Puerto_2 string

func cliente(s string)  {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(s)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func clientePeticion()  {
	s, err := net.Listen("tcp", ":9997")
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
		handleServidor(c)
		s.Close()
		c.Close()
		return
	}
}

func hacerPuerto(x string){
	s, err := net.Listen("tcp", x)
	Puerto_1 = x
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
		handleServidor1(c)
	}
}

func hacerPuerto2(x string){
	s, err := net.Listen("tcp", x)
	Puerto_2 = x
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
		handleServidor2(c)
	}
}

func handleServidor1(c net.Conn)  {
	var s string
	err := gob.NewDecoder(c).Decode(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func handleServidor2(c net.Conn)  {
	var arreglo []string
	err := gob.NewDecoder(c).Decode(&arreglo)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}

func handleServidor(c net.Conn)  {
	var arreglo []string
	err := gob.NewDecoder(c).Decode(&arreglo)
	if err != nil {
		fmt.Println(err)
		return
	}
	go hacerPuerto(arreglo[0])
	go hacerPuerto2(arreglo[1])
}

func main(){
	var Nickname,msj string
	var opcion int64

	fmt.Println("Nickname: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Nickname = scanner.Text()

	fmt.Println("Conectando...")
	go cliente(Nickname)
	go clientePeticion()

	for {
		fmt.Println("Menu\n1.- Enviar mensaje\n2.- Mostrar mensajes\n0.- Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var msjAServer string
			fmt.Println("Mensaje: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			msj = scanner.Text()
			c, err := net.Dial("tcp", ":9998")
			if err != nil {
				fmt.Println(err)
				return
			}
			msjAServer = Nickname + " envio: " + msj
			err = gob.NewEncoder(c).Encode(msjAServer)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			c, err := net.Dial("tcp", ":9996")
			if err != nil {
				fmt.Println(err)
				return
			}
			err = gob.NewEncoder(c).Encode(Puerto_2)
			if err != nil {
				fmt.Println(err)
			}
		case 0:
			var arreglo []string
			c, err := net.Dial("tcp", ":9995")
			if err != nil {
				fmt.Println(err)
				return
			}
			arreglo = append(arreglo,Nickname)
			arreglo = append(arreglo,Puerto_1)
			err = gob.NewEncoder(c).Encode(arreglo)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("hasta pronto "+ Nickname)
			return
		}
	}
} 