package main

func fetchTasks(baseURL, availability string) []Issue {
	// ?

	fullURL := baseURL + "?sort=estimate"
	if availability == "Low" {
		fullURL = fullURL + "&limit=1"
	} else if availability == "High" {
		fullURL = fullURL + "&limit=5"
	} else {
		fullURL = fullURL + "&limit=3"
	}
	return getIssues(fullURL)
}
