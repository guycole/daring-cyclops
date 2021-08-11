<<<<<<< HEAD
// go mod init github.com/guycole/daring-cyclops
// go install github.com/guycole/daring-cyclops

package main

import (
	"github.com/guycole/daring-cyclops/manager"
)

func main() {
	manager.Manager()
=======
package main

import {
  "net/http"
}

func main() {
  mux := http.NewServMux()
  files := http.FileServer(http.Dir("/public"))
  mux.Handle("/static/", http.StripPrefix("/static/"), files)

  mux.HandleFunc("/", index)
  mux.HandleFunc("/err", index)

  mux.HandleFunc("/authenticate", authenticate)

  mux.HandleFunc("/login", login)
  mux.HandleFunc("/logout", logout)

  server := &http.Server {
    Addr:"0.0.0.0:8080", 
    Handler: mux,
  }

  server.ListenAndServe()
>>>>>>> bdf262559b41266a23dd9a4d7f5e0feb1758068e
}
