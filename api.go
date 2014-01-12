package keiko

import (
	//"bytes"
	//"fmt"
	"log"
)

const (
	ALLOFF           = "ALOF \r"
	STATE            = "ACOP \r"
	REDOFF           = "ACOP 0xxxxxxx\r"
	REDON            = "ACOP 1xxxxxxx\r"
	REDBLINK         = "ACOP 2xxxxxxx\r"
	REDQUICKBLINK    = "ACOP 3xxxxxxx\r"
	YELLOWOFF        = "ACOP x0xxxxxx\r"
	YELLOWON         = "ACOP x1xxxxxx\r"
	YELLOWBLINK      = "ACOP x2xxxxxx\r"
	YELLOWQUICKBLINK = "ACOP x3xxxxxx\r"
	GREENOFF         = "ACOP xx0xxxxx\r"
	GREENON          = "ACOP xx1xxxxx\r"
	GREENBLINK       = "ACOP xx2xxxxx\r"
	GREENQUICKBLINK  = "ACOP xx3xxxxx\r"

	BUZZEROFF   = "ACOP xxx00xxx\r"
	BUZZERQUICK = "ACOP xxx10xxx\r"
	BUZZERSLOW  = "ACOP xxx01xxx\r"
)

//--------------LAMP-------------
func Alloff() error {
	err := execute(ALLOFF)
	log.Printf("ALL OFF")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

func RedOn() error {
	err := execute(REDON)
	log.Printf("RED ON")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

func RedOff() error {
	err := execute(REDOFF)
	log.Printf("RED OFF")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func RedBlink() error {
	err := execute(REDBLINK)
	log.Printf("RED SLOW BLINK")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func RedQuickBlink() error {
	err := execute(REDQUICKBLINK)
	log.Printf("RED QUICK BLINK")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

func YellowOn() error {
	err := execute(YELLOWON)
	log.Printf("YELLOW ON")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func YellowOff() error {
	err := execute(YELLOWOFF)
	log.Printf("YELLOW OFF")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func YellowBlink() error {
	err := execute(YELLOWBLINK)
	log.Printf("YELLOW SLOW BLINK")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func YellowQuickBlink() error {
	err := execute(YELLOWQUICKBLINK)
	log.Printf("YELLOW QUICK BLINK")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

func GreenOn() error {
	err := execute(GREENON)
	log.Printf("GREEN ON")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func GreenOff() error {
	err := execute(GREENOFF)
	log.Printf("GREEN OFF")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func GreenBlink() error {
	err := execute(GREENBLINK)
	log.Printf("GREEN SLOW BLINK")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func GreenQuickBlink() error {
	err := execute(GREENQUICKBLINK)
	log.Printf("GREEN QUICK BLINK")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

//--------------BUZZER-------------
func BuzzerOff() error {
	err := execute(BUZZEROFF)
	log.Printf("BUZZER OFF")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func BuzzerQuick() error {
	err := execute(BUZZERQUICK)
	log.Printf("BUZZER QUICK")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
func BuzzerSlow() error {
	err := execute(BUZZERSLOW)
	log.Printf("BUZZER SLOW")
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}

//--------------STATUS-------------

//func checkStatus() error {
//	conn, err := connect()

//	_, err1 := conn.Write([]byte("ACOP \r"))
//	//_, err := conn.Write([]byte("RDMN -u 2 \r")) //機器情報
//	if err1 != nil {
//		log.Printf(err1.Error(), 500)
//	}

//	reply := make([]byte, 128)

//	_, err = conn.Read(reply)
//	if err != nil {
//		fmt.Println("Write to server failed:", err.Error())
//	}

//	n := bytes.Index(reply, []byte{0})
//	s := string(reply[:n])
//	status := statusParse(s)
//	fmt.Println(s)
//	log.Printf(status)
//	conn.Close()
//}

//------------util----------------
