package main

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	flagBind        = flag.String("bind", ":8080", "address to bind server to (eg :8080, localhost:8080, 0.0.0.0:8080)")
	flagCgminerHost = flag.String("api-host", "localhost", "Cgminer API host")
	flagCgminerPort = flag.Int("api-port", 4028, "Cgminer API port")
)

func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	r := gin.Default()

	r.LoadHTMLFiles("./template/index.gohtml")

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	r.GET("/api/health-cgminer", func(c *gin.Context) {
		routeLcdHealthCheck(c)
	})

	r.GET("/", func(c *gin.Context) {
		routeIndex(c)
	})

	return r
}

func routeIndex(c *gin.Context) () {
	lcdResult, err := cgminerLcd()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	devsResult, err := cgminerDevs()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// TODO process more btc addresses
	lcdBytes := toJson(lcdResult.Items[0])
	devsBytes := toJson(devsResult.Items)
	poolInfo := GetPoolInfoFromUser(lcdResult.Items[0].User)

	c.HTML(http.StatusOK, "index.gohtml", map[string]interface{}{
		"status":   string(lcdBytes),
		"devs":     string(devsBytes),
		"poolInfo": poolInfo,
	})
}

func toJson(value interface{}) []byte {
	if result, err := json.MarshalIndent(value, "", "  "); err != nil {
		panic(err)
	} else {
		return result
	}
}

func routeLcdHealthCheck(c *gin.Context) () {
	if _, err := cgminerLcd(); err != nil {
		log.Printf("%v\n", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "DOWN"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	}
}

func main() {
	flag.Parse()

	r := setupRouter()
	if err := r.Run(*flagBind); err != nil {
		panic(err)
	}
}
