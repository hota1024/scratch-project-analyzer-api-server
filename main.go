package handler

import (
	"encoding/json"
	"github.com/ant0ine/go-json-rest/rest"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Project struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Instructions    string `json:"instructions"`
	Visibility      string `json:"visibility"`
	Public          bool   `json:"public"`
	CommentsAllowed bool   `json:"comments_allowed"`
	IsPublished     bool   `json:"is_published"`
	Author          struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		Scratchteam bool   `json:"scratchteam"`
		History     struct {
			Joined time.Time `json:"joined"`
		} `json:"history"`
		Profile struct {
			ID     interface{} `json:"id"`
			Images struct {
				Nine0X90  string `json:"90x90"`
				Six0X60   string `json:"60x60"`
				Five5X55  string `json:"55x55"`
				Five0X50  string `json:"50x50"`
				Three2X32 string `json:"32x32"`
			} `json:"images"`
		} `json:"profile"`
	} `json:"author"`
	Image  string `json:"image"`
	Images struct {
		Two82X218 string `json:"282x218"`
		Two16X163 string `json:"216x163"`
		Two00X200 string `json:"200x200"`
		One44X108 string `json:"144x108"`
		One35X102 string `json:"135x102"`
		One00X80  string `json:"100x80"`
	} `json:"images"`
	History struct {
		Created  time.Time `json:"created"`
		Modified time.Time `json:"modified"`
		Shared   time.Time `json:"shared"`
	} `json:"history"`
	Stats struct {
		Views     int `json:"views"`
		Loves     int `json:"loves"`
		Favorites int `json:"favorites"`
		Comments  int `json:"comments"`
		Remixes   int `json:"remixes"`
	} `json:"stats"`
	Remix struct {
		Parent interface{} `json:"parent"`
		Root   interface{} `json:"root"`
	} `json:"remix"`
}

func Handler() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/projects/:id", func(w rest.ResponseWriter, req *rest.Request) {
			id := req.PathParam("id")
			url := "https://api.scratch.mit.edu/projects/" + id

			response, error := http.Get(url)

			if error != nil {
				rest.Error(w, "", response.StatusCode)
				return
			}

			body, error := ioutil.ReadAll(response.Body)

			if error != nil || response.StatusCode != http.StatusOK {
				rest.Error(w, "", response.StatusCode)
				return
			}

			project := Project{}
			error = json.Unmarshal(body, &project)

			if error != nil {
				return
			}

			w.WriteJson(project)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}