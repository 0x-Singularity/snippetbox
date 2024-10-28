package main

import (
  "log"
  "net/http"
  "strconv"
  "fmt"
)

  // Define a home handler function which writes a byte slice containing
  // "Hello from Snippetbox" as the response body.
  func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
  }

  func snippetView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil || id < 1{
      http.NotFound(w,r)
      return
    }

    msg := fmt.Sprintf("Display a specific snippet with ID %d..", id)
    w.Write([]byte(msg))
  }

  func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet"))
  }



  func main() {
    // Use the http.NewServeMux() function to initialize a new servemux, then
    // register the home function as the handler for the "/" URL pattern.
    mux := http.NewServeMux()
    mux.HandleFunc("/{$}",home)
    mux.HandleFunc("/snippet/view/{id}",snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    // Print a log message to say that the server is starting
    log.Print("starting server on port 4000")

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
  }
