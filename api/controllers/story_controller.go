package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"storyev/api/models"
	"storyev/api/responses"
	"storyev/api/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateStory(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	story := models.Story{}
	err = json.Unmarshal(body, &story)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	storyCreated, err := story.CreateStory(server.DB)

	if err != nil {

		formattedError := utils.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, storyCreated.ID))
	responses.JSON(w, http.StatusCreated, storyCreated)
}

func (server *Server) UpdateStory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	story := models.Story{}
	err = json.Unmarshal(body, &story)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	storyUpdated, err := story.UpdateStory(server.DB, uint32(uid))

	if err != nil {

		formattedError := utils.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, storyUpdated.ID))
	responses.JSON(w, http.StatusCreated, storyUpdated)
}

func (server *Server) DeleteStory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	story := models.Story{}

	fmt.Println(vars)

	uid, err := strconv.ParseInt(vars["id"], 10, 8)

	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = story.DeleteStory(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) GetStories(w http.ResponseWriter, r *http.Request) {
	partQuery := r.URL.Query().Get("part")

	story := models.Story{}

	stories, err := story.FindAllStories(server.DB, partQuery)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, stories)
}

func (server *Server) GetStoryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	story := models.Story{}
	storyGotten, err := story.GetStoryById(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, storyGotten)
}
