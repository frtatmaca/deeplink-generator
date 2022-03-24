package models

// UriResponse holds imported Swagger Info so clients can modify it
type UriRequest struct {
	Link string `json:"link"`
}

// UriResponse holds exported Swagger Info so clients can modify it
type UriResponse struct {
	Link     string `json:"link"`
	DeepLink string `json:"deepLink"`
}
