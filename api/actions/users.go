package actions

import (
	"github.com/BorisLeMeec/BlinkPay/api/models"
	"github.com/satori/go.uuid"
	"github.com/gobuffalo/buffalo"
)

var db = make(map[uuid.UUID]models.User)

type UserResource struct{}

func (ur UserResource) List(c buffalo.Context) error {
	return c.Render(200, r.JSON(db))
}

func (ur UserResource) Create(c buffalo.Context) error {
	// new User
	id, _ := uuid.NewV4()
	user := &models.User{
		// on génère un nouvel id
		ID: id,
	}
	// add in database
	db[user.ID] = *user

	return c.Render(201, r.JSON(user))
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
