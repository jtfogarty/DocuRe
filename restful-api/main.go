package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

//	"database/sql"
//	_ "github.com/lib/pq"

//var db *pq.db

/* func init() {
	//open a db connection
	var err error
	db, err := sql.Open("postgres", "user=theUser dbname=theDbName sslmode=verify-full")
	if err != nil {
  		panic(err)
	}
} */

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/docure")
	{
		v1.GET("/es_info", getESinfo)
		//		v1.GET("/tree_json", getTreeJSON)
	}
	router.Run()

}

type (
	esInfoJSON struct {
		Version     string `json:"version"`
		Name        string `json:"name"`
		ClusterName string `json:"clustername"`
		Status      string `json:"status"`
	}
)

func getESinfo(c *gin.Context) {
	ctx := context.Background()
	var esInfo esInfoJSON
	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if (err != nil) && (code != 200) {
		// Handle error
		panic(err)
	}
	res, err := client.ClusterHealth().Do(context.TODO())
	if err != nil {
		// Handle error
		panic(err)
	}

	esInfo = esInfoJSON{Version: info.Version.Number, Name: info.Name, ClusterName: info.ClusterName, Status: res.Status}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "data": esInfo})

}
