package keiko

import (
	"fmt"
	"net"
)

func connect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", conf.KeikoIP)
	if err != nil {
		fmt.Printf("tcp error")
		return
	}
	return
}

func execute(command string) error {
	fmt.Printf("called execute\n")
	fmt.Printf(command)
	conn, err := connect()
	defer conn.Close()
	if err != nil {
		return err
	}
	_, err1 := conn.Write([]byte(command))
	if err1 != nil {
		return err
	}
	fmt.Printf("command success")

	return nil

}
