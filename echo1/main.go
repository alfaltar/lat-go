package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo"

	"echo/handler"
)

//valyala Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	// Echo instance
	e := echo.New()

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.html", "view/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))
	templates["order.html"] = template.Must(template.ParseFiles("view/order.html", "view/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Route => handler
	e.GET("/", handler.HomeHandler)
	e.GET("/about", handler.AboutHandler)
	e.GET("/baca_menu", handler.BacaData)
	e.GET("/baca_populer", handler.BacaPopuler)
	e.POST("/tambah_menu", handler.AddData)
	e.PUT("/update_menu", handler.UpdateData)
	e.DELETE("/delete_menu", handler.DeleteData) //pakai methode DELETE dengan params (?Id_menu=1)
	e.GET("/order", handler.OrderHandler)
	e.POST("/tambah_order", handler.AddOrder)
	e.Static("/static", "assets")

	// Start the Echo server
	e.Logger.Fatal(e.Start(":1323"))
}
