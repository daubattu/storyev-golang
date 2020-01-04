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

func (server *Server) CreateNewWord(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	newword := models.NewWord{}
	err = json.Unmarshal(body, &newword)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newwordCreated, err := newword.CreateNewWord(server.DB)

	if err != nil {

		formattedError := utils.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, newwordCreated.ID))
	responses.JSON(w, http.StatusCreated, newwordCreated)
}

func (server *Server) UpdateNewWord(w http.ResponseWriter, r *http.Request) {
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
	newword := models.NewWord{}
	err = json.Unmarshal(body, &newword)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newwordUpdated, err := newword.UpdateNewWord(server.DB, uint32(uid))

	if err != nil {

		formattedError := utils.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, newwordUpdated.ID))
	responses.JSON(w, http.StatusCreated, newwordUpdated)
}

func (server *Server) DeleteNewWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newword := models.NewWord{}

	fmt.Println(vars)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = newword.DeleteNewWord(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) GetNewWords(w http.ResponseWriter, r *http.Request) {
	partQuery := r.URL.Query().Get("part")
	storyIdQuery := r.URL.Query().Get("story_id")

	newword := models.NewWord{}

	newwords, err := newword.FindNewWords(server.DB, partQuery, storyIdQuery)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, newwords)
}
