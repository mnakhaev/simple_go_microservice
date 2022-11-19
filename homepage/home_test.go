package homepage

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestHandlers_Handler(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "good",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   message,
		},
	}

	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("failed to initiate the DB")
	}
	defer db.Close()

	dbMock := sqlx.NewDb(db, "sqlmock")
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			h := NewHandlers(nil, dbMock)
			h.Home(test.out, test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("expected: %d\ngot: %d\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected: %s\ngot: %s\n", test.expectedBody, body)
				t.Fail()
			}
		})
	}

}
