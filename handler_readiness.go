package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	respondWithJSON(w, 200, struct{}{})

}

// func handlerErr(w http.ResponseWriter, r *http.Request) {
// 	// respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
// }
