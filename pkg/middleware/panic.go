package middleware

import (
	"github.com/devpies/employee-service/pkg/web"
	"go.uber.org/zap"

	"fmt"
	"net/http"
)

// Panic middleware recovers from panics and converts the panic to an error so it is
// reported in Metrics and handled in Errors.
func Panic(log *zap.Logger) web.Middleware {
	// This is the actual middleware function to be executed.
	f := func(after web.Handler) web.Handler {
		h := func(w http.ResponseWriter, r *http.Request) (err error) {
			// Defer a function to recover from a panic and set the err return
			// variable after the fact.
			defer func() {
				if r := recover(); r != nil {
					// Log the Go stack trace for this panic'd goroutine.
					log.Error("", zap.Error(fmt.Errorf("panic - %v", r)))
				}
			}()

			// Call the next Handler and set its return value in the err variable.
			return after(w, r)
		}

		return h
	}

	return f
}
