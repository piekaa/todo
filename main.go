package main

import (
	"fmt"
	"github.com/todo/config"
)

func main(){

	fmt.Println("Hello World!")
	db,err := config.OpenDb()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select * from task")

	for rows.Next() {
		var id int
		var description string
		var done int
		var name string
		var parentTask int
		err := rows.Scan(&id, &description, &done, &name, &parentTask)

		if err != nil {
			fmt.Println( err )
		}

		fmt.Println( id, description, done, name, parentTask)

	}


	if err != nil {
		fmt.Println(err)
	}


}