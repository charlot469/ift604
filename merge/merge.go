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
	var id = c.Param("id")
	req, err := http.NewRequest("GET", GitlabConstant.ApiUrl+"/projects/"+id+"/merge_requests/"+GitlabConstant.PrivateToken+"&state=opened", nil)


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
	req, err := http.NewRequest("Put", GitlabConstant.ApiUrl+"/projects/"+c.Param("id")+"/merge_requests/"+c.Param("mergeid")+"/merge"+GitlabConstant.PrivateToken, b)
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

func DeleteMerge(c echo.Context) error {
	client := &http.Client{}
	req, err := http.NewRequest("Delete", GitlabConstant.ApiUrl+"/projects/"+c.Param("id")+"/merge_requests/"+c.Param("mergeId")+GitlabConstant.PrivateToken, nil)

	response, error := client.Do(req)

	if error != nil {
		fmt.Printf("%s", err)
		return c.JSON(http.StatusBadRequest, err)
	}


	return c.JSON(response.StatusCode, err)
}
