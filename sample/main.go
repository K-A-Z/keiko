package main

import (
	"github.com/K-A-Z/keiko"
)

func main() {
	srv := keiko.NewKeikoServer()
	srv.ListenAndServe()
}
