package service

import {
    "net/http"
    "encoding/json"
    "github.com/HinanawiTenshi/agenda/service/entities"
    "github.com/unrolled/render"
}

func listAllMeetingHandler(formatter *render.Render) http.handlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        key := getKeyFromRequest(req)
        if verifyKey(key) {
            meetingList, err = entities.MeetingService.FindAll()
            CheckErr(err)
            formatter.JSON(w, http.StatusOK, meetingList)
      }
      else{
          w.WriteHeader(http.StatusForbidden)
      }
    }
}

func createNewMeetingHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        key := getKeyFromRequest(req)
        if verifyKey(key) {
            decoder := json.NewDecoder(req.Body)
            var meeting entities.Meeting
            CheckErr(err)
            owner, err := entities.UserService.FindBy("key",key)
            CheckErr(err)
            meeting.Owner = owner.username
            err := entities.MeetingService.Insert(&meeting)
            CheckErr(err)
            formatter.JSON(w, http.StatusCreated, meeting)
        }
        else{
            w.WriteHeader(http.StatusForbidden)
        }
    }
}

func clearMeetingHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        key := getKeyFromRequest(req)
        if verifyKey(key) {
              user, err := entities.UserService.FindBy("key",key)
              CheckErr(err)
              err := entities.MeetingService.DeleteMeetingByOwner(user)
              CheckErr(err)
              w.WriteHeader(http.StatusNoContent)
        }
        else {
            w.WriteHeader(http.StatusForbidden)
        }
    }
}
