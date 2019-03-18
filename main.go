package main

import (
	"fmt"
	"go-import-csv/controllers"
	"go-import-csv/database"
	"go-import-csv/handler"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// TemplateRegistry -> Define the template registry struct
type TemplateRegistry struct {
	templates *template.Template
}

// DevMode is false if in production
var DevMode bool

func init() {
	viper.BindEnv("ENVIRONMENT")
	devEnv := viper.GetString("Environment")
	if devEnv == "" {
		devEnv = "development"
	}

	if devEnv == "development" {
		DevMode = true
	}
	viper.AddConfigPath("conf/")
	viper.SetConfigType("yaml")
	viper.SetConfigName(devEnv)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	database.Init()
}

// Render -> Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()
	if DevMode {
		e.Debug = true
	} else {
		e.Debug = false
	}

	// Instantiate a template registry and register all html files inside the view folder
	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}

	// Route => handler
	e.GET("/", handler.HomeHandler)

	e.GET("/upload", handler.UploadCSVHandler)
	e.POST("/upload", controllers.Upload)

	// Start the Echo server
	port := viper.GetInt32("app.port")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
