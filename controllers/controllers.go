package controller

import (
	action "dummige/actions"
	"dummige/structs"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo"
)

func Get(c echo.Context) error {

	i := new(structs.Images)

	if err := c.Bind(i); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Image Size
	if c.QueryParam("size") == "" {
		i.W = 140
		i.H = 140
	} else {
		ss := strings.Split(c.QueryParam("size"), "x")

		w, errW := strconv.Atoi(ss[0])
		h, errH := strconv.Atoi(ss[1])

		if errW != nil || errH != nil{
			return c.String(http.StatusBadRequest, "Error: query \"size\" is allow Number x Number only.")
		}

		i.W = w
		i.H = h
	}

	// Image Color
	if c.QueryParam("color") == "" {
		i.C = "555555"
	} else {
		cc := c.QueryParam("color")
		if utf8.RuneCountInString(cc) == 3 || utf8.RuneCountInString(cc) == 6 {
			i.C = cc
		} else {
			i.C = "555555"
		}
	}

	// Image File Name
	i.N = "image.png"

	err := action.CreateImage(i)

	if err != nil {
		return c.String(http.StatusBadRequest, "Error: Failed Create Image.")
	}

	c.Response().Header().Set(echo.HeaderContentType, "image/png")

	c.Response().After(func() {
		os.Remove(i.N)
	})

	return c.File(i.N)

}