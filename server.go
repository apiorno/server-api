package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/apiorno/server-api/controller"
	"github.com/apiorno/server-api/model"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {
	mux := controller.Startup()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
