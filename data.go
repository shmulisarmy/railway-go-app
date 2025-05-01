package main

type Todo struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	Completed        bool   `json:"completed"`
	Time_to_complete int    `json:"time_to_complete"`
	Priority         int    `json:"priority"`
}

var todos []Todo = []Todo{
	{
		Id:               1,
		Title:            "Do laundry",
		Completed:        false,
		Time_to_complete: 10,
		Priority:         1,
	},
	{
		Id:               2,
		Title:            "Make dinner",
		Completed:        true,
		Time_to_complete: 30,
		Priority:         2,
	},
	{
		Id:               3,
		Title:            "Clean the house",
		Completed:        false,
		Time_to_complete: 20,
		Priority:         3,
	},
	{
		Id:               4,
		Title:            "Buy groceries",
		Completed:        true,
		Time_to_complete: 50,
		Priority:         4,
	},
	{
		Id:               5,
		Title:            "Finish homework",
		Completed:        false,
		Time_to_complete: 10,
		Priority:         5,
	},
	{
		Id:               6,
		Title:            "Watch a movie",
		Completed:        true,
		Time_to_complete: 50,
		Priority:         6,
	},
	{
		Id:               7,
		Title:            "Take a shower",
		Completed:        false,
		Time_to_complete: 10,
		Priority:         7,
	},
	{
		Id:               8,
		Title:            "Play video games",
		Completed:        true,
		Time_to_complete: 30,
		Priority:         8,
	},
}
