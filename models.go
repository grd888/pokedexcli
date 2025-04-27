package main

type LocationAreaResponse struct {
	Count int               `json:"count"`
	Next string             `json:"next"`
	Previous string         `json:"previous"`
	Results []LocationArea  `json:"results"`
}

type LocationArea struct {
	Name string  `json:"name"`
	Url string   `json:"url"`
}
