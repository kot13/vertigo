package handlers

import (
	"net/http"
	"github.com/kot13/vertigo/app/renderer"
)

// HealthCheck return health-check info
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	result := map[string]string{
		"status": "ok",
	}

	renderer.Render(result, w)
}
