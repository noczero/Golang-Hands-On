package main

import (
	"Golang-Hands-On/Package/database"
	_ "Golang-Hands-On/Package/helper" // this imported package will not error cause added of '_'
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
