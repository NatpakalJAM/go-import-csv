package controllers

import (
	"bytes"
	"echo-http/database"
	"echo-http/model"
	"net/http"
	"strconv"
	"strings"

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

	buf := new(bytes.Buffer)
	buf.ReadFrom(src)
	s := buf.String()

	arr := strings.Split(s, "\n")
	user := model.User{}
	for i, row := range arr {
		if i > 0 {
			row := strings.Split(row, ",")
			if len(row) > 1 {
				if row[0] == "" {
					user = model.User{
						Firstname: row[1],
						Lastname:  row[2],
					}
					if database.Db.NewRecord(user) {
						database.Db.Create(&user)
					}
				} else {
					id, _ := strconv.Atoi(row[0])
					user = model.User{
						ID:        id,
						Firstname: row[1],
						Lastname:  row[2],
					}
					database.Db.Model(&user).Updates(&user)
				}
			}
		}
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "OK",
		"msg":  "OK",
	})
}
