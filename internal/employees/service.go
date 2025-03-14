package employees

// getEmployees retrieves a paginated list of employees from the employeesList array.
// The offset parameter determines where to start retrieving employees from the array.
// The limit parameter specifies the number of employees to display per page.
func getEmployees(offset, limit int) Employees {
	// If the offset is equal to the length of the array, it means there are no more records to fetch.
	if offset >= len(employeesList) {
		return Employees{}
	}

	// Calculate the end index for the page.
	// min returns the minimum of two integers.
	end := min(offset+limit, len(employeesList))

	return employeesList[offset:end]
}
