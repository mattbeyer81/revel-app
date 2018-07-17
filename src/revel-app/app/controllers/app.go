package controllers

import (
	"github.com/revel/revel"
	"net/http"
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type App struct {
	*revel.Controller
}

type JsonResponse struct {
	CurrentUserUrl string `json:"current_user_url"`
}

func (c App) Index() revel.Result {

	var t = new(JsonResponse)

	fmt.Println("Starting the application...")
    response, _ := http.Get("https://api.github.com/repos/octocat/Spoon-Knife/stats/contributors")

	data, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(data, &t)

	return c.RenderJSON(t)
}

func (c App) Hello() revel.Result {
	return c.Render()
}
