package users

import (
	"fmt"
	"net/http"

	"github.com/prestonbourne/goserve/store"
	"github.com/prestonbourne/goserve/utils"
)

type UserController struct {
	store store.PostgresStore
}

func NewUserController(store store.PostgresStore) *UserController {
	return &UserController{
		store: store,
	}
}

func (c *UserController) Add(w http.ResponseWriter, r *http.Request) error {

	addUserReq := &CreateAccountRequest{}
	// I tried to make the second param for
	if err := utils.DecodeAndWrite(r, addUserReq); err != nil {
		return fmt.Errorf("%w", err)
	}
	newUser := NewUser(addUserReq.FirstName, addUserReq.LastName, addUserReq.UserName)

	//todo:
	c.store.AddUser(newUser.FirstName, newUser.LastName, newUser.UserName, newUser.CreatedAt.UTC())

	return utils.WriteJSON(w, http.StatusOK, newUser)
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) error {
	accounts, err := c.store.GetAllUsers()
	if err != nil {
		utils.LogError("Failed to fetch accounts from Postgres", err)
		return utils.WriteJSON(w, http.StatusInternalServerError, "An Unexpected Error Occured")
	}
	return utils.WriteJSON(w, http.StatusAccepted, accounts)
}

// func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) error {

// }
