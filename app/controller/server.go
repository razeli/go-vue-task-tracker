package controller

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"

	controller "github.com/razeli/go-vue-task-tracker/controller/task"
	"github.com/razeli/go-vue-task-tracker/ui"
)

var router *mux.Router

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}

	rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}

func initHandlers() {
	//router := http.NewServeMux()

	// index page
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
	})

	// static files
	staticFS, _ := fs.Sub(ui.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	router.Handle("/static/js/", httpFS)

	// static files
	httpFS2 := http.FileServer(http.FS(staticFS))
	router.Handle("/static/css/", httpFS2)

	fmt.Println("aaaaaaaaaaaa")

	fjs, _ := fs.ReadDir(ui.StaticFiles, "dist/static/js")
	for _, file := range fjs {
		s := "/static/js/" + file.Name()
		fmt.Println(s)

		router.HandleFunc(s, func(w http.ResponseWriter, r *http.Request) {
			rawFile, _ := ui.StaticFiles.ReadFile("dist" + s)
			if filepath.Ext(s) == ".js" {
				w.Header().Set("Content-Type", "application/javascript")
			}
			w.Write(rawFile)
		})
	}

	fcs, _ := fs.ReadDir(ui.StaticFiles, "dist/static/css")
	for _, file := range fcs {
		s := "/static/css/" + file.Name()
		fmt.Println(s)

		router.HandleFunc(s, func(w http.ResponseWriter, r *http.Request) {
			rawFile, _ := ui.StaticFiles.ReadFile("dist" + s)
			if filepath.Ext(s) == ".css" {
				w.Header().Set("Content-Type", "text/css")
			}
			w.Write(rawFile)
		})
	}

	//fmt.Println(fs.ReadDir(ui.StaticFiles, "dist/js"))

	/**
		staticFS2, _ := fs.Sub(ui.StaticFiles2, "dist")
		httpFS2 := http.FileServer(http.FS(staticFS2))
		router.Handle("/static/js/", httpFS2)

		staticFS3, _ := fs.Sub(ui.StaticFiles3, "dist/static/css")
		httpFS3 := http.FileServer(http.FS(staticFS3))
		router.Handle("/static/css/", httpFS3)
	**/
	// api
	/**	router.HandleFunc("/api/v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, there!"))
	})
	**/

	// we could have /api/post/{id} for GET/PUT/DELETE
	// router.HandleFunc("/api/post/{id}", controller.{appropriateMethod}).Methods("{GET or POST or PUT")
	// instead of what we have.  So the endpoint would be the same to read, update, and delete
	// we'd just have different handlers for those actions

	//router.HandleFunc("/api/tasks", controller.GetPostPut)
	//router.HandleFunc("/api/tasks/{id}", controller.GetDel)

	router.HandleFunc("/api/tasks", controller.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", controller.GetTask).Methods("GET")

	router.HandleFunc("/api/tasks", controller.CreateTask).Methods("POST")

	router.HandleFunc("/api/tasks/{id}", controller.UpdateTask).Methods("PUT")

	router.HandleFunc("/api/tasks/{id}", controller.DeleteTask).Methods("DELETE")

	//router.Use(mux.CORSMethodMiddleware(router))
	//return router

}

func Start() {
	router = mux.NewRouter()

	//r := initHandlers()
	initHandlers()
	fmt.Printf("router initialized and listening on 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
