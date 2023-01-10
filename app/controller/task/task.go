package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/razeli/go-vue-task-tracker/model"
)

func GetPostPut(w http.ResponseWriter, r *http.Request) {
	i := r.Method
	switch i {
	case http.MethodGet:
		GetAllTasks(w, r)
	case http.MethodPost:
		CreateTask(w, r)
	case http.MethodPut:
		UpdateTask(w, r)
	}
}

func GetDel(w http.ResponseWriter, r *http.Request) {
	i := r.Method
	switch i {
	case http.MethodGet:
		GetTask(w, r)
	case http.MethodDelete:
		DeleteTask(w, r)
	}
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllTasks")
	/**w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	**/
	w.Header().Set("Content-Type", "application/json")

	tasks, err := model.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(tasks)
	}
}

// /api/task/{id}
func GetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetTask")
	/**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	**/
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	task, err := model.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(task)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateTask")
	/**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	**/
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var task model.Tasks
	err := decoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var ret model.Tasks
	ret, err = model.CreateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ret)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateTask")
	/**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	**/
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var task model.Tasks
	err := decoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(task)
	err = model.UpdateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteTask")
	/**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	**/
	w.Header().Set("Content-Type", "application/json")
	//fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh" + r.URL.String())
	param := mux.Vars(r)
	//fmt.Println(param)
	idStr := param["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = model.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

/**
func CORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}
**/
