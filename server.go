package keiko

import (
	"fmt"
	"net/http"
)

func NewKeikoServer() *http.Server {
	readConfig()
	fmt.Println(conf.KeikoIP)
	mux := newKeikoMux()
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	return srv
}

func newKeikoMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/alloff/", alloff)
	mux.HandleFunc("/lamp/red/on/", redOn)
	mux.HandleFunc("/lamp/red/off/", redOff)
	mux.HandleFunc("/lamp/red/blink/", redBlink)
	mux.HandleFunc("/lamp/red/quickblink/", redQuickBlink)
	mux.HandleFunc("/lamp/yellow/on/", yellowOn)
	mux.HandleFunc("/lamp/yellow/off/", yellowOff)
	mux.HandleFunc("/lamp/yellow/blink/", yellowBlink)
	mux.HandleFunc("/lamp/yellow/quickblink/", yellowQuickBlink)
	mux.HandleFunc("/lamp/green/on/", greenOn)
	mux.HandleFunc("/lamp/green/off/", greenOff)
	mux.HandleFunc("/lamp/green/blink/", greenBlink)
	mux.HandleFunc("/lamp/green/quickblink/", greenQuickBlink)

	mux.HandleFunc("/buzzer/off/", buzzerOff)
	mux.HandleFunc("/buzzer/quick/", buzzerQuick)
	mux.HandleFunc("/buzzer/slow/", buzzerSlow)

	//	mux.HandleFunc("/lamp/status/", checkStatus)
	return mux
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func alloff(w http.ResponseWriter, r *http.Request) {
	err := Alloff()
	checkResult("ALL OFF", err, w)
}

func redOn(w http.ResponseWriter, r *http.Request) {
	err := RedOn()
	checkResult("RED ON", err, w)
}
func redOff(w http.ResponseWriter, r *http.Request) {
	err := RedOff()
	checkResult("RED OFF", err, w)
}
func redBlink(w http.ResponseWriter, r *http.Request) {
	err := RedBlink()
	checkResult("RED BLINK", err, w)
}
func redQuickBlink(w http.ResponseWriter, r *http.Request) {
	err := RedQuickBlink()
	checkResult("RED QUICKBLINK", err, w)
}

func yellowOn(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("called yellow On\n")
	err := YellowOn()
	checkResult("YELLOW ON", err, w)
}
func yellowOff(w http.ResponseWriter, r *http.Request) {
	err := YellowOff()
	checkResult("YELLOW OFF", err, w)
}
func yellowBlink(w http.ResponseWriter, r *http.Request) {
	err := YellowBlink()
	checkResult("YELLOW BLINK", err, w)
}
func yellowQuickBlink(w http.ResponseWriter, r *http.Request) {
	err := YellowQuickBlink()
	checkResult("YELLOW QUICK BLINK", err, w)
}

func greenOn(w http.ResponseWriter, r *http.Request) {
	err := GreenOn()
	checkResult("GREEN ON", err, w)
}
func greenOff(w http.ResponseWriter, r *http.Request) {
	err := GreenOff()
	checkResult("GREEN OFF", err, w)
}
func greenBlink(w http.ResponseWriter, r *http.Request) {
	err := GreenBlink()
	checkResult("GREEN BLINK", err, w)
}
func greenQuickBlink(w http.ResponseWriter, r *http.Request) {
	err := GreenQuickBlink()
	checkResult("GREEN QUICK BLINK", err, w)
}

func buzzerOff(w http.ResponseWriter, r *http.Request) {
	err := BuzzerOff()
	checkResult("BUZZER OFF", err, w)
}
func buzzerQuick(w http.ResponseWriter, r *http.Request) {
	err := BuzzerQuick()
	checkResult("BUZZER QUICK", err, w)
}
func buzzerSlow(w http.ResponseWriter, r *http.Request) {
	err := BuzzerSlow()
	checkResult("BUZZER SLOW", err, w)
}

func checkResult(com string, err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Fprintf(w, "ERROR: "+com+" is FAILED")
	} else {
		fmt.Fprintf(w, com+" is SUCCESS")
	}
}
