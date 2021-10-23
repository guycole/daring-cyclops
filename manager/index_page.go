// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"log"
	"net/http"
)

func index(response http.ResponseWriter, request *http.Request) {
	// get active games

	log.Println("thread noted")
	log.Println(response)
	log.Println(request)

	session := sessionTest(request, response)
	log.Println(session)
	if session.Id == "" {
		log.Println("nil")
		//generateHTML(response, threads, "layout", "public.navbar", "index")
	} else {
		log.Println("not nil")
	}

	rc := freshRedisConnection()
	log.Println(rc)

	//getSession(rc, "bogus")

	/*
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
	*/

}
