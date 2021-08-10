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
}
