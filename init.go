package groupietracker

import (
	"net/http"
)

func Web() {
	// http.HandleFunc("/", Menu)     // Menu page
	// http.HandleFunc("/home", Home) // Game page
	// http.HandleFunc("/victory", Victory)
	// http.HandleFunc("/lose", Lose)
	// http.HandleFunc("/input", Input)
	// http.HandleFunc("/reset", Reset)

	fs := http.FileServer(http.Dir("serv/"))
	http.Handle("/serv/", http.StripPrefix("/serv/", fs))

	// fmt.Println("(http://localhost:8080) - server started on port", 8080)
	http.ListenAndServe(":8080", nil)

}
