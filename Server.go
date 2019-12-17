
package main

import (
	"awesomeProject/merge"
	"awesomeProject/projects"
	"github.com/labstack/echo"
)

//main function
func main() {
	// create a new echo instance
	e := echo.New()
	projects.InitUser()

	e.GET("user/:username/projects/longitude/:longitude/latitude/:latitude", projects.GetProjects)
	e.DELETE("user/:username/projects/:id", projects.DeleteProject)

	e.GET("user/:username/projects/:id/MergeRequests", merge.GetMergeRequests)
	e.PUT("user/:username/projects/:id/MergeRequests/:mergeId/merge", merge.AcceptMerge)
	e.GET("ws/hook", merge.GitlabHook)
	e.GET("ws/notification", merge.Notification)

	e.Logger.Fatal(e.Start(":8000"))
}