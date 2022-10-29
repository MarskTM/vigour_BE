package controller

import (
	"vigour/model"
	"vigour/service"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ProfileController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByUserId(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Upsert(w http.ResponseWriter, r *http.Request)
}

type profileController struct {
	profileService service.ProfileService
}

// Get profile by id godoc
// @tags profile-manager-apis
// @Summary Get profile by id
// @Description input: profile's id => output: struct profile
// @Accept json
// @Produce json
// @Param id path integer true "profile's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /profile/{id} [get]
func (c *profileController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get profile failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.profileService.GetById(id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get profile successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get all profiles godoc
// @tags profile-manager-apis
// @Summary Get profiles
// @Description output: struct profiles
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /profile/all [get]
func (c *profileController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	tmp, err := c.profileService.GetAll()
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get profiles successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get profile by userId godoc
// @tags profile-manager-apis
// @Summary Get profile by userId
// @Description input: user_id => output: struct profile
// @Accept json
// @Produce json
// @Param user_id path integer true "true id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /profile/user/{user_id} [get]
func (c *profileController) GetByUserId(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "user_id")
	user_id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get profile failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.profileService.GetByUserId(user_id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get profile by userId successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Create profile godoc
// @tags profile-manager-apis
// @Summary Create profiles
// @Description input: profile model.profile => output: status
// @Accept json
// @Produce json
// @param profile body model.Profile true "fill profile"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /profile/create [post]
func (c *profileController) Create(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var profile model.Profile

	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "create profile failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}
	tmp, err := c.profileService.Create(&profile)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "create profiles successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Update profile godoc
// @tags profile-manager-apis
// @Summary Update profiles
// @Description input: profile model.profile => output: status
// @Accept json
// @Produce json
// @param profile body model.Profile true "change profile"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /profile/update [put]
func (c *profileController) Update(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var profile model.Profile

	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "update profile failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.profileService.Update(profile.Id, profile)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "update profiles successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Delete profile by id godoc
// @tags profile-manager-apis
// @Summary Delete profile by id
// @Description input: profile's id => output: status
// @Accept json
// @Produce json
// @Param id path integer true "profile's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /profile/delete/{id} [delete]
func (c *profileController) Delete(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "delete profile failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	if err := c.profileService.Delete(id); err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    nil,
			Message: "delete profile successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Upsert profile godoc
// @tags profile-manager-apis
// @Summary Upsert profiles
// @Description Upsert profile: if not exist => create
// @Accept json
// @Produce json
// @param profile body model.Profile true "info profile"
// @Success 200 {object} model.Response
// @Router /profile/upsert [put]
func (c *profileController) Upsert(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var profile model.Profile

	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "upsert profile failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.profileService.Upsert(&profile)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "upsert profiles successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

func NewProfileController() ProfileController {
	profileService := service.NewProfileService()
	return &profileController{
		profileService: profileService,
	}
}
