package main

import (
	"log"
	"net/http"

	"github.com/mykhalskyio/elifTech-testTask/internal/config"
	"github.com/mykhalskyio/elifTech-testTask/internal/db"
	"github.com/mykhalskyio/elifTech-testTask/internal/handler"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := db.NewConnect(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.MigrationInitUp()
	if err != nil {
		log.Fatalln(err)
	}
	handler := handler.NewHandler(db)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Banks)
	mux.HandleFunc("/banks", handler.Banks)
	mux.HandleFunc("/mortgages", handler.Mortgages)
	mux.HandleFunc("/deletebank", handler.DeleteBank)
	mux.HandleFunc("/editbank", handler.EditBankPage)
	mux.HandleFunc("/editbankdb", handler.EditBank)
	mux.HandleFunc("/addbank", handler.AddBank)
	mux.HandleFunc("/addbankdb", handler.AddBankdb)
	mux.HandleFunc("/deletemortage", handler.DeleteMortage)
	mux.HandleFunc("/addmortage", handler.AddMortgage)
	log.Println("Server start in localhost:4000")
	log.Println(http.ListenAndServe(":25565", mux))
}
