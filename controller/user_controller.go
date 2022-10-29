package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vigour/infrastructure"
	"vigour/model"
	"vigour/service"

	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)

	Login(w http.ResponseWriter, r *http.Request)
	LoginWithToken(w http.ResponseWriter, r *http.Request)
	GetByUsername(w http.ResponseWriter, r *http.Request)
}

// GetAll gets all users currently in table "users"
// @tags user-manager-apis
// @Summary get all users
// @Description get all users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/all [get]
func (c *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.Response

	users, err := c.userService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		jsonResponse = &model.Response{
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Data:    users,
			Message: "OK",
			Success: true,
		}
	}
	render.JSON(w, r, jsonResponse)
}

// GetAll gets user currently in table "users"
// @tags user-manager-apis
// @Summary get user by id
// @Description get user by id
// @Accept json
// @Produce json
// @Param uid path integer true "User ID"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/{uid} [get]
func (c *userController) GetById(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.Response

	strId := chi.URLParam(r, "uid")
	uid, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		infrastructure.ErrLog.Println(err)
		jsonResponse = &model.Response{
			Data:    nil,
			Message: "invalid UID in URL",
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
	}
	userInfo, err := c.userService.GetById(uid)
	if err != nil {
		jsonResponse = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Data:    userInfo,
			Message: "Get user info successful!",
			Success: true,
		}
	}
	render.JSON(w, r, jsonResponse)
}

// GetByUsername gets user currently in table "users" with username
// @tags user-manager-apis
// @Summary get user with usn
// @Description input username => user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "username"
// @Success 200 {object} model.Response
// @Router /user/wname [get]
func (c *userController) GetByUsername(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.Response
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), 400)
		jsonResponse = &model.Response{
			Message: "Username is required",
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}
	user, err := c.userService.GetByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		jsonResponse = &model.Response{
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Data:    user,
			Message: "OK",
			Success: true,
		}
	}
	render.JSON(w, r, jsonResponse)
}

// CreateUser creates an user with given data
// @tags user-manager-apis
// @Summary	creates new user
// @Description creates new user
// @Accept json
// @Produce json
// @Param UserInfo body model.User true "User information"
// @Success 200 {object} model.Response
// @Router /user/create [post]
func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInfo *model.User
	var jsonResponse *model.Response

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInfo); err != nil {
		jsonResponse = &model.Response{
			Data:    nil,
			Message: "Error decoding request body:" + err.Error(),
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}

	newUser, err := c.userService.CreateUser(userInfo)

	if err != nil {
		jsonResponse = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Data:    newUser,
			Message: "User created succesfully!",
			Success: true,
		}
	}
	render.JSON(w, r, jsonResponse)

}

// UpdateUser update an userPassword with given data
// @tags user-manager-apis
// @Summary	Update new user password
// @Description Update UserPassword by userId
// @Accept json
// @Produce json
// @Param UserInfo body model.User true "User information"
// @Success 200 {object} model.Response
// @Router /user/update [put]
func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.Response
	var user model.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		jsonResponse = &model.Response{
			Message: err.Error(),
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}

	newUser, err := c.userService.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		jsonResponse = &model.Response{
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Data:    newUser,
			Message: "OK",
			Success: true,
		}
	}

	render.JSON(w, r, jsonResponse)
}

// DeleteUser deletes user with UserID
// @tags user-manager-apis
// @Summary delete user
// @Description delete user
// @Accept json
// @Produce json
// @Param uid path integer true "User ID"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/delete/{uid} [delete]
func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.Response

	strID := chi.URLParam(r, "uid")
	uid, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		jsonResponse = &model.Response{
			Message: "Error decoding request body:" + err.Error(),
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}

	if _, err := c.userService.DeleteUser(uid); err != nil {
		jsonResponse = &model.Response{
			Message: err.Error(),
			Success: false,
		}
	} else {
		jsonResponse = &model.Response{
			Message: "User deleted successfully!",
			Success: true,
		}
	}

	render.JSON(w, r, jsonResponse)
}

// Login log user in if they have valid credential
// @tags user-manager-apis
// @Summary log user in
// @Description log user in
// @Accept json
// @Produce json
// @Param LoginPayload body model.UserPayload true "username & password"
// @Success 200
// @Router /user/login [post]
func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.LoginResponse
	var loginDetail model.UserPayload

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginDetail); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		jsonResponse = &model.LoginResponse{
			Message: "Bad request: " + err.Error(),
			Code:    "400",
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}

	user, token, refreshToken, err := c.userService.LoginRequest(loginDetail.Username, loginDetail.Password)
	if err != nil {
		jsonResponse = &model.LoginResponse{
			Message: "Wrong username or password. Info:" + err.Error(),
			Code:    "400",
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}
	//user da bi xoa
	if user.DeletedAt != nil {
		jsonResponse = &model.LoginResponse{
			Token:        "user has been deleted!",
			RefreshToken: "user has been deleted!",
			Message:      "user has been deleted!",
			Code:         "406",
			Success:      true,
		}
	} else {
		jsonResponse = &model.LoginResponse{
			Token:        token,
			RefreshToken: refreshToken,
			UserId:       user.Id,
			Role:         user.Role,
			Username:     user.Username,
			Message:      "Logged in successfully as " + user.Role,
			Code:         "200",
			Success:      true,
		}
	}

	render.JSON(w, r, jsonResponse)
}

// LoginWithToken provides token each login attempt
// @tags user-manager-apis
// @Summary login user
// @Description login user, return new token string jwt
// @Accept json
// @Produce json
// @Param refToken query string true "Insert your refresh token"
// @Success 200 {object} model.LoginResponse
// @Router /user/login/jwt [post]
func (c *userController) LoginWithToken(w http.ResponseWriter, r *http.Request) {
	var jsonResponse *model.LoginResponse

	refToken := r.URL.Query().Get("refToken")
	user, accessToken, refreshToken, success, err := c.userService.LoginWithToken(refToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		jsonResponse = &model.LoginResponse{
			Message: err.Error(),
			Code:    "401",
			Success: false,
		}
		render.JSON(w, r, jsonResponse)
		return
	}
	if !success {
		jsonResponse = &model.LoginResponse{
			Token:   accessToken,
			Message: err.Error(),
			Code:    "400",
			Success: false,
		}
	} else {
		jsonResponse = &model.LoginResponse{
			Token:        accessToken,
			RefreshToken: refreshToken,
			UserId:       user.Id,
			Username:     user.Username,
			Role:         user.Role,
			Message:      "jwt login successful!",
			Code:         "200",
			Success:      true,
		}
	}
	render.JSON(w, r, jsonResponse)
}

func NewUserController() UserController {
	return &userController{
		userService: service.NewUserService(),
	}
}
