package webserver

import "net/http"

func MiWebServer() {
	//Es un manejador o enrutador de web server
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
