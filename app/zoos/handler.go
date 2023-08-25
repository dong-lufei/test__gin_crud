package shop

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Zoos struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getZoos(c *gin.Context) {
	// ps := make([]Zoos, 0)

	// data, err := ioutil.ReadFile("./db/zoos.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 这里参数要指定为变量的地址
	// err = json.Unmarshal(data, &ps)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(ps)
	// // fmt.Println(strin)
	c.JSON(http.StatusOK, gin.H{"msg": "getZoos"})
}
func addZoos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "addZoos"})
}
