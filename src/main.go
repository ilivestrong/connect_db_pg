package main

import (
	"fmt"

	"github.com/ilivestrong/db_connect_pg/src/persist"
)

func main() {
	fmt.Println("Main()")
	persist.Initialize()
}
