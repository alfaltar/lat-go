package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

var baseURL = "http://localhost:1323"

func HomeHandler(c echo.Context) error {

	var datax, err = show()
	var datap, errp = showpopuler()
	if err != nil {

	}

	if errp != nil {

	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":  "HOME",
		"msg":   "Framework Echo Worked!",
		"data":  datax,
		"datap": datap,
	})
}

func show() ([]menu, error) {
	var err error
	var client = &http.Client{}
	var data []menu

	request, err := http.NewRequest("GET", baseURL+"/baca_menu", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func showpopuler() ([]menu, error) {
	var err error
	var client = &http.Client{}
	var data []menu

	request, err := http.NewRequest("GET", baseURL+"/baca_populer", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
