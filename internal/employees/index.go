package employees

import (
	"net/http"
	"strconv"

	"go.leapkit.dev/core/render"
	"go.leapkit.dev/core/server"
)

// internal/employees/index.go
// Index renders the home page of the application, displaying a paginated list of employees.
func Index(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())

	// Offset is a parameter used for pagination. It determines where to start retrieving employees from the employeesList array.
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if offset < 0 {
		offset = 0
	}

	limit := 16                  // Number of employees to display per page.
	nextOffset := offset + limit // Calculate the next offset for pagination.

	total := len(employeesList)          // Get the total number of employees.
	hasMoreRecords := total > nextOffset // Determine if more employees exist beyond the current page.

	// Fetch the employees for the current page.
	employees := employees(offset, limit)

	rw.Set("employees", employees)
	rw.Set("hasMoreRecords", hasMoreRecords)
	rw.Set("nextOffset", nextOffset)

	// Check if the request is an HTMX request (partial rendering)
	if r.Header.Get("HX-Request") != "" {
		err := rw.RenderClean("employees/list.html")
		if err != nil {
			server.Errorf(w, http.StatusInternalServerError, "error rendering list template: %s", err.Error())
			return
		}
		return
	}

	// Render the full page template for normal requests.
	err := rw.Render("employees/index.html")
	if err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error rendering template: %s", err.Error())
	}
}
