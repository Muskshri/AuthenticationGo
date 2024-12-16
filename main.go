package main

import (
	"gomod/database"
	"gomod/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv" //to read data from env
)

func main(){
	//adding .env file
	err:= godotenv.Load() //to load .env file
	if err!= nil{
		log.Fatal("Error loading .env file") //the server should not start if .env not uploaded.
	}
   
   //connection to db
	db, err:= database.GetDB()
	if err!= nil{
		log.Fatal("Error getting db conneciotn")
	}


	if db!= nil{
		sqlDB,_ := db.DB()
		defer sqlDB.Close()
	}
    
	router:= mux.NewRouter()
	routes.SetRouter(router)

	port := ":8080"
	fmt.Printf("server starting at port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}