package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes = Routes{

	Route{
		"TodoIndex",
		"GET",
		"/supermarket/",
		Index,
	},
	Route{
		"TodoCreate",
		"GET",
		"/supermarket/get/createrepo/",
		TodoCreate,
	},
	Route{
		"TodoCreateRepo",
		"POST",
		"/supermarket/createrepo/",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/supermarket/{produceId}/",
		TodoShow,
	},
	Route{
		"TodoNew",
		"GET",
		"/supermarket/get/new/{produceName}/{producePrice}/{produceCode}/",
		TodoAdd,
	},
	Route{
		"TodoNewProduce",
		"POST",
		"/supermarket/new/{produceName}/{producePrice}/{produceCode}/",
		TodoAdd,
	},
	Route{
		"Delete",
		"GET",
		"/supermarket/get/delete/{produceId}/",
		TodoDelete,
	},
	Route{
		"DeleteProduce",
		"DELETE",
		"/supermarket/delete/{produceId}/",
		TodoDelete,
	},
	Route{
		"ShowAll",
		"GET",
		"/supermarket/all/showall/",
		TodoShowAll,
	},


}
