package converter

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Get("/api", Home)
	r.Get("/api/get-encode", GetHandlerEncode)
	r.Get("/api/get-marshal", GetHandlerMarshal)
	r.Post("/api/list", PostConversions)
	r.Post("/api/weight-us", PostWeightUS)
	r.Post("/api/volume-us", PostVolumeUS)
	r.Post("/api/weight-metric", PostWeightMetric)
	r.Post("/api/volume-metric", PostVolumeMetric)
	r.Post("/api/height-feet", PostHeightFeet)
	r.Post("/api/height-metric", PostHeightMetric)

	return r
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, welcome to the recipe .."))
}

func GetHandlerMarshal(w http.ResponseWriter, r *http.Request) {
	HandleGetRequestMarshal(w, r, Input, RecipeList)
}

func GetHandlerEncode(w http.ResponseWriter, r *http.Request) {
	HandleGetRequestEncode(w, r, SingleInput, RecipeList)
}

func PostConversions(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, RecipeList)
}

func PostWeightUS(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, WeightUS)
}

func PostVolumeUS(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, VolumeUS)
}

func PostWeightMetric(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, WeightMetric)
}

func PostVolumeMetric(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, VolumeMetric)
}

func PostHeightFeet(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, FromFeet)
}

func PostHeightMetric(w http.ResponseWriter, r *http.Request) {
	HandlePostRequest(w, r, FromMetric)
}
