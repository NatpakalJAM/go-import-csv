package controllers

import (
	"echo-http/database"
	"echo-http/model"
	"encoding/csv"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo"
)

func Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	mimeType := file.Header.Get("Content-Type")
	if mimeType != "application/octet-stream" {
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"name": "Fail",
			"msg":  "The format file is not valid.",
		})
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	user := model.User{}
	reader := csv.NewReader(src)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.Render(http.StatusOK, "home.html", map[string]interface{}{
				"name": "Fail",
				"msg":  "Something is broken.",
			})
		}
		if !reflect.DeepEqual(record, []string{"id", "firstname", "lastname"}) {
			if record[0] != "" {
				id, _ := strconv.Atoi(record[0])
				user = model.User{
					ID:        id,
					Firstname: record[1],
					Lastname:  record[2],
				}
				database.Db.Model(&user).Updates(&user)
			} else {
				user = model.User{
					Firstname: record[1],
					Lastname:  record[2],
				}
				if database.Db.NewRecord(user) {
					database.Db.Create(&user).AutoMigrate()
				}
			}
		}
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "OK",
		"msg":  "OK",
	})
}
