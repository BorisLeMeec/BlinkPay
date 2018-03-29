package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func Pay(c buffalo.Context) error {
	// new User
	pictureEncoded := c.Request().FormValue("base64")
	resp, err := http.PostForm("192.168.0.13:8080/facebox/check",
		url.Values{"base64": {pictureEncoded}})
	if err != nil {
		c.Render(501, r.String("internal error"))
		fmt.Println(err.Error())
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		c.Render(500, r.String("Internal Error"))
	}

	var rCheck ResponseCheck

	json.Unmarshal(res, &rCheck)
	if rCheck.Success == false || rCheck.FacesCount < 1 {
		c.Render(401, r.JSON(ResultPayment{false, "no one recognized on the picture",
				1, 0, ""}))
	}
	return c.Render(200, r.JSON(ResultPayment{true, "", 0, 0,
	rCheck.Faces[0].ID}))
}

type ResponseCheck struct {
	Success bool `json:"success"`
	FacesCount int `json:"facesCount"`
	Faces []Face `json:"faces"`
}

type Face struct {
	Rect Rectangle `json:"rect"`
	ID_pic string `json:"id"`
	ID string `json:"name"`
	Matched bool `json:"matched"`
	Confidence float64 `json:"confidence"`
}

type Rectangle struct {
	Top int `json:"top"`
	Left int `json:"left"`
	Width int `json:"width"`
	Height int `json:"height"`
}

type ResultPayment struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	ErrorID int `json:"errorId"`
	TransactionID int `json:"transactionId"`
	Name string `json:"name"`

}