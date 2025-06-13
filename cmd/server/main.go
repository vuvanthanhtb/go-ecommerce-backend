package main

import (
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8081")
}
