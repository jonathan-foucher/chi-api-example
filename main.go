package main

import (
	"fmt"
    "net/http"
    "os"
	"github.com/joho/godotenv"
	"chi-api-example/routers"
)

func main() {
	godotenv.Load()

    router := routers.InitRouter()
    HTTP_PORT := os.Getenv("HTTP_PORT")
    fmt.Println("Application is starting on port", HTTP_PORT)
    http.ListenAndServe(":" + HTTP_PORT, router)
}
