package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/xyproto/onthefly"
)

// Create a Three.JS page
func ThreeJSPage() *onthefly.Page {
	p, t := onthefly.NewThreeJS("My first Three.js app")

	t.AddCamera()
	t.CameraPos("z", 5)

	t.AddRenderer()

	// Blue cube
	cube1 := t.AddTestCube("cube1")

	// Red cube
	cube2 := t.AddTestCube("cube2")
	cube2.Js += "cube2.rotation.x += 0.9;"

	// Render function
	r := onthefly.NewRenderFunction()

	r.AddJS(cube1.Id + ".rotation.x += 0.02;")
	r.AddJS(cube1.Id + ".rotation.y += 0.02;")

	r.AddJS(cube2.Id + ".rotation.x += 0.04;")
	r.AddJS(cube2.Id + ".rotation.y += 0.07;")

	t.AddRenderFunction(r, true)

	return p
}

// Set up the paths and handlers then start serving.
func main() {
	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	// Create the page
	page := ThreeJSPage()

	// Publish the generated page (HTML and CSS)
	page.Publish(mux, "/", "/style.css", false)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 3000
	n.Run(":3000")
}
