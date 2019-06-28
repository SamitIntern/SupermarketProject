package main

import (
	"fmt"
	"strings"
	"reflect"
	"regexp"
)

var currentId int

var todos Todos

//This function contributes to the functionality of addition of a new produce to the repository
//This function checks the repository if a produce with the same produce code is present in the repository
//This function is called by the function "RepoAddItem" of this class

func CheckForExistingProduce(todo Todo) Todo {
	for _, t := range todos {
		if strings.EqualFold(todo.ProduceCode, t.ProduceCode) {
			return t
		}
	}
	return Todo{}
}

//This function contributes to the functionality of deletion of a produce from the repository
//This function checks the repository if a produce with the same produce ID is present in the repository
//This function is called by the function "RepoDeleteItem" of this class

func CheckForNonExistingProduce(produceId int) Todo {
	for _, t := range todos {
		if t.Id == produceId {
			return t
		}
	}
	return Todo{}
}

//This function contributes to the functionality of searching for a produce in the repository
//This function checks the repository if a produce with the same produce ID is present in the repository
//This function is called by the function "TodoShow" of the "handlers.go" class

func RepoFindItem(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

//This function contributes to the functionality of creating the repository
//This function is called by the function "TodoCreate" of the "handlers.go" class

func RepoCreateRepo() Todos {

	var newtodos Todos

	var t Todo
	currentId = 0
	t = RepoCreateItem(Todo{Name: "Lettuce", ProduceCode:"A12T-4GH7-QPL9-3N4M", UnitPrice:3.46 })
	newtodos = append(newtodos, t)
	t = RepoCreateItem(Todo{Name: "peach", ProduceCode:"E5T6-9UI3-TH15-QR88", UnitPrice:2.99})
	newtodos = append(newtodos, t)
	t = RepoCreateItem(Todo{Name: "Green Pepper", ProduceCode:"YRT6-72AS-K736-L4AR", UnitPrice:0.79})
	newtodos = append(newtodos, t)
	t = RepoCreateItem(Todo{Name: "Gala Apple", ProduceCode:"TQ4C-VV6T-75ZX-1RMR", UnitPrice:3.59})
	newtodos = append(newtodos, t)
	todos = newtodos

	return todos
}

//This function contributes to the functionality of addition of a new produce to the repository and creation of repository
//This function is called by the functions "RepoAddItem" and "RepoCreateRepo" of this class

func RepoCreateItem(t Todo) Todo {

	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

//This function contributes to the functionality of addition of a new produce to the repository
//This function is called by the function "TodoAdd" of "handlers.go" class

func RepoAddItem(t Todo) Todo {

	var oldTodo = CheckForExistingProduce(t)
	if oldTodo.Id > 0 {
		return Todo{}
	}

	if !CheckValidAddParamTypes(t){
		return Todo{}
	}

	if !CheckValidAddParamNumber(t){
		return Todo{}
	}

	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

//This function contributes to the functionality of removal of a produce from the repository
//This function is called by the function "TodoDelete" of "handlers.go" class

func RepoDeleteItem( produceId int) Todos {

	var oldTodo = CheckForNonExistingProduce(produceId)
	if oldTodo.Id == 0 {
		return Todos{}
	}

	var newTodos Todos
	for _, todo := range todos {
		if todo.Id == produceId{
			continue
		}else {
			newTodos = append(newTodos, todo)
		}
	}
	todos = newTodos
	return newTodos
}

//This function contributes to the functionality of removal of a produce from the repository
//This function is called by the function "TodoShowAll" of "handlers.go" class

func GetAllItems() Todos {
	fmt.Print(todos)
	return todos
}

//This function contributes to the functionality of addition of a new produce to the repository
//This function validates the params before a new produce is added to the repository
//It checks if the "Produce Code" and "Produce Name" are strings
//It also validates if the "Produce Code" is 19 character long alphanumeric string with each 4 characters separated by
// a hyphen
//This function is called by the function "RepoAddItem" of this class

func CheckValidAddParamTypes(todo Todo) bool{
	var valid bool
	valid = true
	if reflect.TypeOf(todo.Name).String() != "string"{
		valid = false
		return valid
	}
	if reflect.TypeOf(todo.ProduceCode).String() != "string"{
		valid = false
		return valid
	}
	if len(todo.ProduceCode) != 19 {
		valid = false
		return valid
	}

	x := regexp.MustCompile(`-`)
	substrings := x.Split(todo.ProduceCode, -1)
	if len(substrings) == 4{
		for _, subs := range substrings{
			if strings.Contains(subs, "-"){
				valid = false
				return valid
			}
		}
	}else{
		valid = false
		return valid
	}

	return valid
}

//This function contributes to the functionality of addition of a new produce to the repository
//This function validates the params before a new produce is added to the repository by making sure that neither the
// "Produce Name" or the "Produce Code" fields are empty
//This function is called by the function "RepoAddItem" of this class

func CheckValidAddParamNumber(todo Todo) bool{
	var valid bool
	valid = true
	if todo.Name == ""{
		valid = false
		return valid
	}
	if todo.ProduceCode == ""{
		valid = false
		return valid
	}
	return valid
}
