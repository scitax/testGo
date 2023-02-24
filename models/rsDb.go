package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
)

func RSearch() {
	// Create a client. By default a client is schemaless
	// unless a schema is provided when creating the index
	c := redisearch.NewClient(fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), "myIndex")

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("body")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	// Drop an existing index. If the index does not exist an error is returned
	c.Drop()

	// Create the index with the given schema
	if err := c.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}

	docMap := make([]redisearch.Document, 0)
	// Create a document with an id and given score
	doc := redisearch.NewDocument("doc1", 1.0)
	doc.Set("title", "Hello world XX!!!").
		Set("body", "foo bar").
		Set("date", time.Now().Unix())

	docMap = append(docMap, doc)
	// // Index the document. The API accepts multiple documents at a time
	// if err := c.IndexOptions(redisearch.DefaultIndexingOptions, doc); err != nil {
	// 	log.Fatal(err)
	// }

	// Create a document with an id and given score
	doc1 := redisearch.NewDocument("doc2", 1)
	doc1.Set("body", "abra cadabra 124")

	docMap = append(docMap, doc1)
	// Index the document. The API accepts multiple documents at a time
	if err := c.IndexOptions(redisearch.DefaultIndexingOptions, docMap...); err != nil {
		fmt.Println(docMap)
		log.Fatal(err)
	}
	doc2 := redisearch.NewDocument("doc3", 1)
	doc2.Set("title", "Hello world YY!!!")
	// Index the document. The API accepts multiple documents at a time
	if err := c.IndexOptions(redisearch.DefaultIndexingOptions, doc2); err != nil {
		fmt.Println(docMap)
		log.Fatal(err)
	}

	// Searching with limit and sorting
	docs, total, err := c.Search(redisearch.NewQuery("hello world"))
	fmt.Println(docs, total, err)
	fmt.Println(docs[0].Properties["body"])

}
