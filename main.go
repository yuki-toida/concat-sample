package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/xid"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")
	e.File("/", "index.html")

	e.POST("/upload", func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			panic(err)
		}

		uuid := xid.New().String()
		dir := "assets/video/" + uuid
		output := uuid + ".mp4"
		if err := os.MkdirAll(dir, 0777); err != nil {
			panic(err)
		}

		inputText := uuid + ".txt"
		files := []string{}
		for _, v := range form.File {
			file := v[0]
			src, err := file.Open()
			if err != nil {
				panic(err)
			}
			defer src.Close()

			path := dir + "/" + file.Filename
			dst, err := os.Create(path)
			if err != nil {
				panic(err)
			}
			defer dst.Close()

			if _, err = io.Copy(dst, src); err != nil {
				panic(err)
			}
			files = append(files, "file '"+path+"'")
		}

		txt, err := os.Create(inputText)
		if err != nil {
			panic(err)
		}
		defer txt.Close()
		txt.Write(([]byte)(strings.Join(files, "\n")))

		err = exec.Command("ffmpeg", "-f", "concat", "-i", inputText, "-c", "copy", "assets/video/"+output).Run()
		if err != nil {
			panic(err)
		}

		if err := os.Remove(inputText); err != nil {
			panic(err)
		}
		if err := os.RemoveAll(dir); err != nil {
			panic(err)
		}

		return c.String(http.StatusOK, output)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
