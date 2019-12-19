package merge

import (
	"awesomeProject/GitlabConstant"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

type merge struct
{
	id int
	merge_request_iid int
	merge_commit_message string
}

func GetMergeRequests(c echo.Context) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", GitlabConstant.Url+"/projects/"+c.QueryParam(":id")+"/merge_requests/"+GitlabConstant.AccesToken+"&state=opened", nil)


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

func AcceptMerge(c echo.Context) (err error) {
	u := new(merge)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(u); err != nil {
		 return c.JSON(http.StatusBadRequest, err)
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)

	client := &http.Client{}
	req, err := http.NewRequest("Put", GitlabConstant.Url+"/projects/"+c.QueryParam(":id")+"/merge_requests/"+c.QueryParam(":mergeid")+"/merge"+GitlabConstant.AccesToken, b)
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
