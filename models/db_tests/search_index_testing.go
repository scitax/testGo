package db_tests

import (
	"fmt"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/scitax/testGo/models"
)

func SearchIndexTest() error {
	pool := models.CreateConnectionPool()
	// unless a schema is provided when creating the index
	c := redisearch.NewClientFromPool(pool, "test index")
	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("body")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	// Drop an existing index. If the index does not exist an error is returned
	c.Drop()

	// Create the index with the given schema
	if err := c.CreateIndex(sc); err != nil {
		return err
	}

	// Create a document with an id and given score
	doc := redisearch.NewDocument("doc1", 1.0)
	doc.Set("title", "Hello world").
		Set("body", "foo bar").
		Set("date", time.Now().Unix())

	// Index the document. The API accepts multiple documents at a time
	if err := c.IndexOptions(redisearch.DefaultIndexingOptions, doc); err != nil {
		return err
	}

	// Searching with limit and sorting
	if _, _, err := c.Search(redisearch.NewQuery("hello world").Limit(0, 2).SetReturnFields("title")); err != nil {
		return err
	}

	return nil
}

func SearchIndex() error {
	pool := models.CreateConnectionPool()
	// unless a schema is provided when creating the index
	c := redisearch.NewClientFromPool(pool, "idx")
	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTagField("tags")).
		AddField(redisearch.NewTextField("title")).
		AddField(redisearch.NewNumericField("date"))

	// Drop an existing index. If the index does not exist an error is returned
	c.Drop()

	// Create the index with the given schema
	if err := c.CreateIndex(sc); err != nil {

		return err

	}

	// Create a document with an id and given score
	doc := redisearch.NewDocument("doc1", 1.0)
	doc.Set("title", "Hello world").
		Set("tags", "golang, redis, database").
		Set("date", time.Now().Unix())

	// Index the document. The API accepts multiple documents at a time
	if err := c.IndexOptions(redisearch.DefaultIndexingOptions, doc); err != nil {
		return err
	}

	// res, _, err := c.Search(redisearch.NewQuery("golang"))
	// fmt.Println(res, err)
	query := redisearch.NewQuery("golang")
	res, _, _ := c.Search(query)
	fmt.Println(res, "!!!!TAGGS SYKA RABOTAY!!!")

	return nil
}
