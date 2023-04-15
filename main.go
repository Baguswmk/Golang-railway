package main

import (
	"DTS-Kominfo-Hactiv8/Chapter3/Challange2/database"
	"DTS-Kominfo-Hactiv8/Chapter3/Challange2/router"
)

func main(){
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}