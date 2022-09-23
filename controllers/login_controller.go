package controllers

import (
	"encoding/json"
	"net/http"
	"testgo11/models"
	"testgo11/utils"
)

var LoginController = func(w http.ResponseWriter, r *http.Request) {
	usr := &models.Usr{}
	err := json.NewDecoder(r.Body).Decode(usr)

	if err != nil {
		utils.RespondError(w, utils.Message(false, err.Error()), http.StatusBadRequest)
		return
	}
	data, err := utils.CreateToken(*usr)
	if err != nil {
		utils.RespondError(w, utils.Message(false, err.Error()), http.StatusBadRequest)
		return
	}
	resp := utils.MessageData(true, "sucess", data)
	utils.Respond(w, resp)
}
