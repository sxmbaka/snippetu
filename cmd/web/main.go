package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	var err error
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLogFile, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	infoLogToFle := log.New(infoLogFile, "INFO\t", log.Ldate|log.Ltime)
	infoLogToStdio := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLogFile, err := os.OpenFile("/tmp/error.log", os.O_RDWR|os.O_CREATE, 0666)
	errorLogToFile := log.New(errorLogFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogToStdio := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLogToFile,
		Handler:  mux,
	}

	infoLogToFle.Println("Staring server at http://localhost" + *addr)
	infoLogToStdio.Println("Staring server at http://localhost" + *addr)
	err = srv.ListenAndServe()
	errorLogToFile.Fatal(err)
	errorLogToStdio.Fatal(err)
}
