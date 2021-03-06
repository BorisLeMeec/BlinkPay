package actions

import (
	"github.com/BorisLeMeec/BlinkPay/api/models"
	"github.com/satori/go.uuid"
	"github.com/gobuffalo/buffalo"
	"github.com/machinebox/sdk-go/facebox"
)

var db = make(map[uuid.UUID]models.User)

type UserResource struct{}

func (ur UserResource) List(c buffalo.Context) error {
	return c.Render(200, r.JSON(db))
}

func (ur UserResource) Create(c buffalo.Context) error {
	// new User
	u := &models.User{}
	u.FirstName = c.Request().FormValue("FirstName")
	u.LastName = c.Request().FormValue("LastName")
	//if err := c.Bind(u1); err != nil {
	//	return c.Render(500, r.String(err.Error()))
	//}
	id, _ := uuid.NewV4()
	u.ID = id
	db[u.ID] = *u

	return c.Render(201, r.JSON(u))
}

func (ur UserResource) Show(c buffalo.Context) error {
	// get id and format to uuid
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		// if id isnt uuid
		return c.Render(500, r.String("id is not uuid v4"))
	}

	// get user in database
	user, ok := db[id]
	if ok {
		// if exist return user
		return c.Render(200, r.JSON(user))
	}

	// if not exist return not found
	return c.Render(404, r.String("user not found"))
}

func (ur UserResource) Check(c buffalo.Context) error {
	file, _, err := c.Request().FormFile("camera_pic")
	if err != nil {
		return c.Render(500, r.String(err.Error()))
	}
	faceboxClient := facebox.New("http://localhost:8080")
	faces, err := faceboxClient.Check(file)
	if len(faces) == 1 {
		if faces[0].Matched == true {
			return c.Render(200, r.JSON(faces[0].Name))
		}
	}
	return c.Render(401, r.String("no one recognized"))
}
