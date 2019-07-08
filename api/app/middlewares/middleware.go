package middlewares

import (
	"fmt"
	"justicia/api/app/models/helpers"
	"net/http"

	"go.uber.org/zap"
)

// PanicRecoveryMiddleware handles the panic in the handlers.
func PanicRecoveryMiddleware(h http.HandlerFunc, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the error
				logger.Error(fmt.Sprint(rec))

				// write the error response
				helpers.JSONResponse(w, 500, map[string]interface{}{
					"error": "Internal Error",
				})
			}
		}()

		h(w, r)
	}
}
