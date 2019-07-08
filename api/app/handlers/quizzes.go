package handlers

import (
	"justicia/api/app/models/helpers"
	"justicia/api/app/models/oposicion"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

func GetQuizListHandler(w http.ResponseWriter, r *http.Request) {
	quizzes, err := di.Get(r, "quiz-manager").(*oposicion.QuizManager).GetAll()

	if err == nil {
		helpers.JSONResponse(w, 200, quizzes)
		return
	}

	helpers.JSONResponse(w, 500, map[string]interface{}{
		"error": "Internal Error",
	})
}
func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["quizId"]

	quiz, err := di.Get(r, "quiz-manager").(*oposicion.QuizManager).Get(id)

	if err == nil {
		helpers.JSONResponse(w, 200, quiz)
		return
	}

	switch e := err.(type) {
	case *helpers.ErrNotFound:
		helpers.JSONResponse(w, 404, map[string]interface{}{
			"error": e.Error(),
		})
	default:
		helpers.JSONResponse(w, 500, map[string]interface{}{
			"error": "Internal Error",
		})
	}
}
func PutQuizHandler(w http.ResponseWriter, r *http.Request) {
	var input *oposicion.Quiz

	err := helpers.ReadJSONBody(r, &input)
	if err != nil {
		helpers.JSONResponse(w, 400, map[string]interface{}{
			"error": "Could not decode request body.",
		})
		return
	}

	id := mux.Vars(r)["quizId"]

	quiz, err := di.Get(r, "quiz-manager").(*oposicion.QuizManager).Update(id, input)

	if err == nil {
		helpers.JSONResponse(w, 200, quiz)
		return
	}

	switch e := err.(type) {
	case *helpers.ErrValidation:
		helpers.JSONResponse(w, 400, map[string]interface{}{
			"error": e.Error(),
		})
	case *helpers.ErrNotFound:
		helpers.JSONResponse(w, 404, map[string]interface{}{
			"error": e.Error(),
		})
	default:
		helpers.JSONResponse(w, 500, map[string]interface{}{
			"error": "Internal Error",
		})
	}
}
