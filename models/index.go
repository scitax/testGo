package models

import (
	"fmt"
)

func CreateTAGIndex(field string, idx string) {

	_, err := client.Do("FT.CREATE", idx, "ON", "HASH", "PREFIX", "1", "fw_rules:fw-", "SCHEMA", field, "TAG").Result()
	fmt.Println(err)

}

func DropIndex(idx string) {
	res, err := client.Do("FT.DROPINDEX", idx, "DD").Result()
	fmt.Println(res, err)
}

func SearshTag(idx string, field string, tags string) {
	query := fmt.Sprintf("@%s: {%s}", field, tags)
	_, err := client.Do("FT.SEARH", idx, query).Result()
	fmt.Println(err)
}

func Addrule() {
	pipe := client.Pipeline()
	pipe.HSet("fw_rules:fw-1", "fw_id", "fw-1")
	pipe.HSet("fw_rules:fw-1", "source_tag", "tag1")
	pipe.HSet("fw_rules:fw-1", "dest_tag", "tag2")
	pipe.HSet("fw_rules:fw-2", "fw_id", "fw-2")
	pipe.HSet("fw_rules:fw-2", "source_tag", "tag1")
	pipe.HSet("fw_rules:fw-2", "dest_tag", "tag2")
	pipe.HSet("fw_rules:fw-21", "fw_id", "fw-21")
	pipe.HSet("fw_rules:fw-21", "source_tag", "tag1")
	pipe.HSet("fw_rules:fw-21", "dest_tag", "tag2")
	_, err := pipe.Exec()
	fmt.Println(err)
}
func Search() {
	res, err := client.Do("FT.SEARCH", "tag-index", "@dest_tag:{ tag1 }").Result()
	fmt.Println("Results!!!!")
	fmt.Println(res, err)
}
func Getvalues() {
	// if err := client.Set("key", "value12222", 0).Err(); err != nil {
	// 	fmt.Println(err)
	// }

	res, err := client.HGet("doc2", "title").Result()
	fmt.Println(res, err, "!!!!!!")
}
