package projects

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"awesomeProject/GitlabConstant"
)

var UserWithRight []user
var officeLat = 32.32
var officeLong = 43.43
var allowedDistance = 100.0

type user struct
{
	name string `json:"name"`
}

func InitUser() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", GitlabConstant.Url+"/users"+GitlabConstant.AccesToken, nil)


	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)

		}

		var results []user
		json.Unmarshal(contents, &results)
		UserWithRight = results
	}
}

func GetProjects(c echo.Context) error {
	if !ensureUser(c.QueryParam("username")){
		return c.JSON(http.StatusForbidden, "user has no right")
	}

	if !ensureEmplacement(parseFloat(c.QueryParam("longitude")), parseFloat(c.QueryParam("latitude"))){
		return c.JSON(http.StatusForbidden, "user is too far from office")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", GitlabConstant.Url+"/projects"+GitlabConstant.AccesToken, nil)


	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		return c.JSON(http.StatusBadRequest, err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("%s", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		var results []map[string]interface{}
		json.Unmarshal(contents, &results)
		return c.JSON(response.StatusCode, &results)
	}
}

func DeleteProject(c echo.Context) error {
	id := c.Param("id")
	client := &http.Client{}
	req, err := http.NewRequest("Delete", GitlabConstant.Url + "/projects/"+ id + GitlabConstant.AccesToken, nil)
	_, err = client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func ensureUser(username string) bool{
	for _, a := range UserWithRight {
		if a.name == username {
			return true
		}
	}
	return false
}

func ensureEmplacement(long float64, lat float64) bool{
	return math.Sqrt(math.Pow(long-officeLong,2)+math.Pow(lat-officeLat,2)) < allowedDistance
}


func parseFloat(f string) float64{
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		return s
	}

	return -1
}