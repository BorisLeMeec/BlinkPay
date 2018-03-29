package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/machinebox/sdk-go/facebox"
)

func Pay(c buffalo.Context) error {
	// new User
	pictureEncoded := c.Request().FormValue("base64")

	faceboxClient := facebox.New("http://localhost:8080")
	faces, err := faceboxClient.CheckBase64(pictureEncoded)

	if err != nil {
		return c.Render(500, r.JSON(ResultPayment{false, err.Error(), 1, 0, ""}))
	}

	return c.Render(200, r.JSON(ResultPayment{true, "", 0, 1,
	faces[0].ID}))

	//resp, err := http.PostForm("192.168.0.13:8080/facebox/check",
	//	url.Values{"base64": {pictureEncoded}})
	//if err != nil {
	//	c.Render(501, r.String("internal error"))
	//	fmt.Println(err.Error())
	//}
	//res, err := ioutil.ReadAll(resp.Body)
	//if err != nil{
	//	c.Render(500, r.String("Internal Error"))
	//}

	//var rCheck ResponseCheck

	//json.Unmarshal(res, &rCheck)
	//if rCheck.Success == false || rCheck.FacesCount < 1 {
	//	c.Render(401, r.JSON(ResultPayment{false, "no one recognized on the picture",
	//			1, 0, ""}))
	//}
	//return c.Render(200, r.JSON(ResultPayment{true, "", 0, 0,
	//rCheck.Faces[0].ID}))
}

type ResultPayment struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	ErrorID int `json:"errorId"`
	TransactionID int `json:"transactionId"`
	Name string `json:"name"`

}