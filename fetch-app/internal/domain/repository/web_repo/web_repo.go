package web_repo

import "net/http"

type webrepo struct {
	httpClient  *http.Client
	resourceUrl string
}

func New(httpClient *http.Client, resourceUrl string) *webrepo {
	return &webrepo{
		httpClient:  httpClient,
		resourceUrl: resourceUrl,
	}
}
