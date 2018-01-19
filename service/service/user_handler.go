package service

import {
    "net/http"
    "encoding/json"
    "github.com/HinanawiTenshi/agenda/service/entities"
    "github.com/unrolled/render"
}

func listAllUsersHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        key := getKeyFromRequest(req)
        if verifyKey(key) {
            userList, err = entities.UserService.FindAll()
            for i := range userList {
                userList[i].Key = "*****"
                userList[i].Password = "*****"
            }
            CheckErr(err)
            formatter.JSON(w, http.StatusOK, userList)
        }
        else{
            w.WriteHeader(http.StatusForbidden)
        }
    }
}

func deleteCurrentUserHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        key := getKeyFromRequest(req)
        if verifyKey(key) {
            err = entities.UserService.DeleteByKey(key)
            CheckErr(err)
            w.WriteHeader(http.StatusNoContent)
        }
        else{
            w.WriteHeader(http.StatusForbidden)
        }
    }
}

func getUserKeyHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseFrom()
        user, err := entities.UserService.FindBy("username", req.FormValue("username"))
        CheckErr(err)
        if user.Password == req.FromValue("password"){
            formatter.JSON(w, http.StatusOK, user)
        }
        else{
            w.WriteHeader(http.StatusForbidden)
        }
    }
}

func creatNewUserHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        decoder := json.NewDecoder(req.Body)
        var user entities.User
        err := decoder.Decode(&user)
        CheckErr(err)
        check, err := entities.UserService.FindBy("username", user.Username)
        if check != (entities.User{}){
            w.WriteHeader(http.StatusForbidden)
            return
        }
        CheckErr(err)
        err := entities.UserService.Insert(&user)
        CheckErr(err)
        formatter.JSON(w, http.StatusCreated, user)
    }
}
