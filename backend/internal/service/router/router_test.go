package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestLogger struct{}

func (TestLogger) Fatal(msg string) {}
func (TestLogger) Info(msg string) {
}
func (TestLogger) Infof(template string, args ...interface{}) {
}
func (TestLogger) Infow(msg string, keysAndValues ...interface{}) {
}
func (TestLogger) RequestEnd(act string, startAt time.Time, status *int, errMsg *string) {
}
func (TestLogger) Warn(msg string) {
}
func TestManager_healthCheck(t *testing.T) {
	var mng Manager
	mng.lgr = TestLogger{}

	req, _ := http.NewRequest("GET", "/services/health", nil)
	w := httptest.NewRecorder()
	mng.healthCheck(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
