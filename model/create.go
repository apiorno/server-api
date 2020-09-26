package model

import "fmt"

//CreateTodo creates a todo in the DB
func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?,?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//DeleteTodo deletes a todo by name
func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
