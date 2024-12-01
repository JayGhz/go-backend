package app

import "net/http"

func (app *Application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) { // Cambiado de 'healthcheckHandler' a 'HealthcheckHandler'
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"alive": true}`))
}
