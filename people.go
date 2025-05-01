package main

type Person struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Parent_id     int    `json:"parent_id"`
	Image         string `json:"image"`
	Gender        string `json:"gender"`
	Is_descendant bool   `json:"is_descendant"`
	Spouse_id     int    `json:"spouse_id"`
}
