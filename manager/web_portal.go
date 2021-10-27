// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func splashPage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "hello xxx\n")
}

func webPortal() {
	log.Println("webPortal entry")

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(configuration.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)

	// error
	mux.HandleFunc("/err", err)

	// starting up the server
	server := &http.Server{
		Addr:           configuration.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(configuration.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(configuration.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

	log.Println("webPortal exit")
}

// GET /err?msg=
func err(writer http.ResponseWriter, request *http.Request) {
	log.Println("err noted")
	/*
		vals := request.URL.Query()
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
		} else {
			generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
		}
	*/
}

/*
func index(writer http.ResponseWriter, request *http.Request) {

c getSession(rc *redis.Client, id string) (session *Session, err error) {

		threads, err := data.Threads()
		if err != nil {
			error_message(writer, request, "Cannot get threads")
		} else {
			_, err := session(writer, request)
			if err != nil {
				generateHTML(writer, threads, "layout", "public.navbar", "index")
			} else {
				generateHTML(writer, threads, "layout", "private.navbar", "index")
			}
		}

	log.Println("thread noted")
}
*/

func generateHTML(writer http.ResponseWriter, filenames ...string) {
	//func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	log.Println("generate html")

	var files []string
	files[0] = "index"
	/*
		for _, file := range filenames {
			// replace data interface
			files = append(files, fmt.Sprintf("templates/%s.html", file))
		}
	*/

	templates := template.Must(template.ParseFiles(files...))
	log.Println(templates)
	//	templates.ExecuteTemplate(writer, "layout", data)
}
