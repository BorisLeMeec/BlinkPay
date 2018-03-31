package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/machinebox/sdk-go/facebox"
)

func Pay(c buffalo.Context) error {
	pictureEncoded := c.Request().FormValue("base64")

	faceboxClient := facebox.New("http://localhost:8080")
	faces, err := faceboxClient.CheckBase64(pictureEncoded)

	if err != nil {
		return c.Render(401, r.JSON(ResultPayment{false, err.Error(), 1, 0, ""}))
	}

	return c.Render(200, r.JSON(ResultPayment{true, "", 0, 1,
	faces[0].Name}))
}

type ResultPayment struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	ErrorID int `json:"errorId"`
	TransactionID int `json:"transactionId"`
	Name string `json:"name"`

}