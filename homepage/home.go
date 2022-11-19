package homepage

import (
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

const message = "test message"

type Handlers struct {
	logger *log.Logger
	db     *sqlx.DB
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.db.ExecContext(r.Context(), "")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

// SetupRoutes allows to set up needed routes in other package, not in main.
// It also shows example of how middleware function can be used for handler.
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}
