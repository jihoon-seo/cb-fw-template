package common

import (
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type myObject struct {
	MyField string `json:"myField"`
	Message string `json:"message"`
}

func RestGetObject(c echo.Context) error {

	content := map[string]string{"message": "RestGetObject called."}

	PrintJsonPretty(content)

	return c.JSON(http.StatusOK, &content)
}

func RestGetAllObjects(c echo.Context) error {

	content := map[string]string{"message": "RestGetAllObjects called."}

	PrintJsonPretty(content)

	return c.JSON(http.StatusOK, &content)
}

func RestPostObject(c echo.Context) error {

	content := map[string]string{"message": "RestPostObject called."}

	PrintJsonPretty(content)

	return c.JSON(http.StatusOK, &content)
}

func RestDeleteObject(c echo.Context) error {

	content := map[string]string{"message": "RestDeleteObject called."}

	PrintJsonPretty(content)

	return c.JSON(http.StatusOK, &content)
}

func ApiValidation() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Printf("%v\n", "[API request!]")

			return next(c)
		}
	}
}
