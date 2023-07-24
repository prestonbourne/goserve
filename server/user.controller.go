package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/prestonbourne/goserve/models"
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

	addUserReq := &models.AddUserRequest{}
	// I tried to make the second param for
	if err := utils.DecodeAndWrite(r, addUserReq); err != nil {
		return fmt.Errorf("%w", err)
	}
	newUser, _ := models.NewUser(
		addUserReq.FirstName,
		addUserReq.LastName,
		addUserReq.UserName,
		addUserReq.Password,
	)

	//todo:
	c.store.AddUser(r.Context(), newUser.FirstName, newUser.LastName, newUser.UserName, newUser.CreatedAt.UTC())

	return utils.WriteJSON(w, http.StatusOK, newUser)
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) error {
	users, err := c.store.GetUsers()
	if err != nil {
		utils.LogError("Failed to fetch accounts from Postgres", err)
		return utils.WriteJSON(w, http.StatusInternalServerError, "An Unexpected Error Occured")
	}
	return utils.WriteJSON(w, http.StatusAccepted, users)
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) error {

	id, err := utils.GetId(r)

	if err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, err)
	}

	user, err := c.store.GetUserByID(id)
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, err.Error())
	}
	return utils.WriteJSON(w, http.StatusAccepted, user)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := utils.GetId(r)

	if err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, err)
	}

	user, err := c.store.GetUserByID(id)

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, err)
	}

	err = c.store.DeleteUser(id)

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(w, http.StatusAccepted, fmt.Sprintf("Deleted: %+v", user))

}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) error {
	var req models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, req)
}

// JWT MIDDLEWARE, FINISH LATER - not really sure if i even need this
const jwtSecret string = "password"

func validateJWT(tokenStr string) (*jwt.Token, error) {

	//this should be an environment variable
	// os.Getenv("JWT_SECRET") -> export the secret with this name in terminal;

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		// validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing the secret, e.g []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})
}
func createJWT(user *models.User) (string, error) {

	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userName":  user.UserName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString([]byte(jwtSecret))

}

func withJWTAuth(handler http.HandlerFunc) http.HandlerFunc {
	//middleware

	return func(w http.ResponseWriter, r *http.Request) {
		//call handler

		tokenStr := r.Header.Get("x-jwt-token")

		_, err := validateJWT(tokenStr)
		// I should probably make the API error type useful
		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, "Invalid Token")
			return
		}

		handler(w, r)
	}
}
