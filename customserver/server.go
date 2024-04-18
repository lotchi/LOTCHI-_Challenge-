package main

import (
	"log/slog"
	"net/http"

	"lotchiexample.com/customserver/constants"
	"lotchiexample.com/customserver/handlers"
)

/** redirect incomming request https **/
func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, constants.HTTPS+r.RequestURI,
		http.StatusMovedPermanently,
	)
}

func main() {
	slog.Info("Listening " + constants.HTTPS)
	http.HandleFunc("/", handlers.Home)
	go http.ListenAndServeTLS(
		constants.HTTPS,
		constants.CERT,
		constants.KEY,
		nil,
	)

	if err := http.ListenAndServe(constants.HTTP, http.HandlerFunc(redirectToHTTPS)); err != nil {
		slog.Error("failed to start custom server", "err", err)
	}
}
