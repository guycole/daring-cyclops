package manager

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const Banner = "Daring Cyclops V0.0"

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var configuration Configuration

const MaxGames = 5

type Game struct {
	Active       bool
	RandomId     string
	SequentialId int
}

var allGames [MaxGames]Game

var maxSequentialId = -1

func loadConfiguration() {
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatalln("open config file failure", err)
	}

	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln("configuration parse failure", err)
	}
}

func getSequentialId() int {
	var currentId = maxSequentialId + 1
	maxSequentialId = currentId
	return maxSequentialId
}

func setupGame(ndx int) {
	log.Println("setupGame:", ndx)
	allGames[ndx].Active = false
	allGames[ndx].RandomId = uuid.NewString()
	allGames[ndx].SequentialId = getSequentialId()
}

func setup() {
	for ndx := 0; ndx < MaxGames; ndx++ {
		setupGame(ndx)
	}
}

func err(writer http.ResponseWriter, request *http.Request) {
	//vals := request.URL.Query()
	//_, err := session(writer, request)
	//if err != nil {
	//	generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	//} else {
	//	generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	//}
}

func index(writer http.ResponseWriter, request *http.Request) {
	//threads, err := data.Threads()
	//if err != nil {
	//	error_message(writer, request, "Cannot get threads")
	//} else {
	//	_, err := session(writer, request)
	//	if err != nil {
	//		generateHTML(writer, threads, "layout", "public.navbar", "index")
	//	} else {
	//		generateHTML(writer, threads, "layout", "private.navbar", "index")
	//	}
	//}
}

func httpSetup() {
	log.Println("http setup")
	// handle static assets
	//mux := http.NewServeMux()
	file_server := http.FileServer(http.Dir(configuration.Static))
	//http.Handle("/", file_server)
	http.Handle("/static/", http.StripPrefix("/static/", file_server))
	//mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	//mux.HandleFunc("/", index)
	// error
	//mux.HandleFunc("/err", err)

	http.HandleFunc("/", serveTemplate)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func init() {
	loadConfiguration()
	setup()
}

func Manager() {
	log.Println(Banner)

	// fall into event loop

	httpSetup()
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
