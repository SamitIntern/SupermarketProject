package main

type Todo struct {
	Id        		 int       	`json:"id"`
	Name      		 string    	`json:"name"`
	ProduceCode      string     `json:"pcode"`
	UnitPrice        float64    `json:"price"`
}

type Todos []Todo
