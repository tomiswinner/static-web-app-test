// ソースコード元
// https://qiita.com/ozora/items/0597e52b3f9c1759e292
// FUNCTIONS_CUSTOMHANDLER_PORT は、Azure Functions の環境変数
// https://learn.microsoft.com/ja-jp/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows#create-and-build-your-function


package main

import (
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/api/goapi1", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
    r.Run(":" + port)
}
