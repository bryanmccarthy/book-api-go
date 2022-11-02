package main

import (
    "github.com/gin-gonic/gin"
    "example.com/book-api-go/pkg/books"
    "example.com/book-api-go/pkg/common/db"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    books.RegisterRoutes(r, h)
    // register more routes here

    r.Run(port)
}
