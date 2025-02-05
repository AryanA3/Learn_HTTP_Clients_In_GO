package main

func fetchTasks(baseURL, availability string) []Issue {
	var avail string

	switch availability {
	case "Low":
		avail = "1"
	case "Medium":
		avail = "3"
	case "High":
		avail = "5"
	default:
		avail = ""
	}

	fullURL := baseURL + "?sort=estimate&limit=" + avail
	return getIssues(fullURL)
}
