package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  //"time"
)

const VERSION = "v1"
const API_NAME = "lb"

type JsonData struct{
  Created int64
  IP string `json:"ip" binding:"required"`
  Server string `json:"server" binding:"required"`
}


func main(){
  api := "/" + VERSION + "/" + API_NAME + "/"
  //task := "/" + VERSION + "/task/"

  router := gin.Default()

  //Task router
  //router.GET(task,)
  //router.GET(task/:task_id,)

  //LoadBalance router
  router.GET(api, func(cxt *gin.Context) {
    cxt.String(http.StatusOK, "Hello World!")
  })

  router.POST(api, func(c *gin.Context) {
    var json JsonData
    c.BindJSON(&json)

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in", "IP": json.IP})
  })



  //Run at 8080
  router.Run(":8080")
}
