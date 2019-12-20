
package main

import (
	"awesomeProject/merge"
	"awesomeProject/projects"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//main function
func main() {
	// create a new echo instance
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("projects/longitude/:longitude/latitude/:latitude", projects.GetProjects)
	e.GET("projects/:id", projects.GetProject)
	//e.DELETE("projects/:id", projects.DeleteProject)

	e.GET("projects/:id/MergeRequests", merge.GetMergeRequests)
	e.PUT("projects/:id/MergeRequests/:mergeId/merge", merge.AcceptMerge)
	e.DELETE("/projects/:id/MergeRequests/:mergeId", merge.DeleteMerge)

	e.GET("ws/hook", merge.GitlabHook)
	e.GET("ws/notification", merge.Notification)

	e.Logger.Fatal(e.Start(":8000"))
}