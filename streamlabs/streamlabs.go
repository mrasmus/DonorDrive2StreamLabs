package streamlabs

import (
	"golang.org/x/oauth2"
)

// Endpoint is StreamLab's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://streamlabs.com/api/v1.0/authorize",
	TokenURL: "https://streamlabs.com/api/v1.0/token",
}
