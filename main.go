package main

import (
	"github.com/afiifatuts/bankmnc/router"
)

func main() {
	r := router.StartApp()
	r.Run("localhost:8000")
}
