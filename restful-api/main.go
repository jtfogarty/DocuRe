package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

//	"database/sql"
//	_ "github.com/lib/pq"

//var db *pq.db
var env sysEnv

func init() {
	//open a db connection

	env.ESPort = os.Getenv("ES_PORT_NO")
	env.ESServer = os.Getenv("ES_SERVER_NAME")

	/* 	var err error
		db, err := sql.Open("postgres", "user=theUser dbname=theDbName sslmode=verify-full")
		if err != nil {
	  		panic(err)
		} */
}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/docure")
	{
		v1.GET("/es_info", getESinfo)
		//		v1.GET("/tree_json", getTreeJSON)
		v1.GET("/get_indexes", getIndexes)
		v1.GET("/get_book_chpt/:index/:book/:chpt", getBookChapter)
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
	esIndexList struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	esBookChapter struct {
		TookInMillis int64   `json:"tookinmills"`
		TotalHits    int64   `json:"totalhits"`
		Index        string  `json:"index"`
		Verses       []verse `json:"verses"`
	}
	verse struct {
		VerseNo   string `json:"verse_no"`
		TextEntry string `json:"text_entry"`
		LineID    string `json:"line_id"`
		Book      string `json:"book"`
		ChapterNo string `json:"chapter_no"`
		Type      string `json:"type"`
	}
	sysEnv struct {
		ESPort   string
		ESServer string
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
	info, code, err := client.Ping("http://" + env.ESServer + ":" + env.ESPort).Do(ctx)
	if (err != nil) && (code != 200) {
		// Handle error
		panic(err)
	}
	res, err := client.ClusterHealth().Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	esInfo = esInfoJSON{Version: info.Version.Number, Name: info.Name, ClusterName: info.ClusterName, Status: res.Status}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "data": esInfo})
	/* 	client.Delete().Do(ctx)
	   	res1, err := client.Delete().Do(ctx)
	   	if err != nil {
	   		// Handle error
	   		panic(err)
	   	}
	   	println(res1.Result) */
}

func getIndexes(c *gin.Context) {
	//ctx := context.Background()
	var index []esIndexList
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	names, err := client.IndexNames()
	if err != nil {
		// Handle error
		panic(err)
	}
	for _, iname := range names {
		index = append(index, esIndexList{Name: iname, Status: "OK"})
		//fmt.Printf("%s\n", name)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "data": index})
	/* 	res2, err := client.Delete().Do(ctx)
	   	if err != nil {
	   		// Handle error
	   		panic(err)
	   	}
	   	println(res2.Result) */
}

func getBookChapter(c *gin.Context) {
	var bookChapter esBookChapter
	var vrs verse
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	
	qb := elastic.NewMatchQuery("book", c.Param("book"))
	bq := elastic.NewBoolQuery().Must(qb)
	bq.Must(elastic.NewMatchQuery("chapter_no", c.Param("chpt")))

	searchResult, err := client.Search().
		Index(c.Param("index")). // search in index "twitter"
		Query(bq).               // specify the query
		Sort("line_id", true).   // sort by "user" field, ascending
		//FilterPath("type", "verse").
		From(0).Size(200). // take documents 0-9
		Pretty(true).      // pretty print request and response JSON
		Do(ctx)            // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	//fmt.Printf("Found a total of %d verses\n", searchResult.Hits.TotalHits)

	bookChapter.TookInMillis = searchResult.TookInMillis

	bookChapter.TotalHits = searchResult.Hits.TotalHits

	bookChapter.Index = c.Param("index")

	for _, hit := range searchResult.Hits.Hits {
		// hit.Index contains the name of the index
		err := json.Unmarshal(*hit.Source, &vrs)
		if err != nil {
			panic(err)
		}

		//		s := string(data)
		//		sU, _ := strconv.Unquote(s)
		bookChapter.Verses = append(bookChapter.Verses, vrs)
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "data": bookChapter})
}
