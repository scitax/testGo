package models

import "fmt"

func GetIdis() {
	value := client.Get("key")
	fmt.Print(value)
}
