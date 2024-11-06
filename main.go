package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	db "chi-api-example/database"
	"chi-api-example/routers"
)

func main() {
	godotenv.Load()
	
	conn := db.InitDbConn()
	defer conn.Close(context.Background())
	
	router := routers.InitRouter()
	HTTP_PORT := os.Getenv("HTTP_PORT")
	fmt.Println("Application is starting on port", HTTP_PORT)
	http.ListenAndServe(":" + HTTP_PORT, router)
}
