package main

import (
	"keiko"
)

func main() {
	srv := keiko.NewKeikoServer()
	srv.ListenAndServe()
}
