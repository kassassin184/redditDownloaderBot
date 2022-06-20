package config

import (
	"net/http"
	"time"
)

// Version is the version of this program :|
const Version = "2.4.0"

// GlobalHttpClient is a http client which all request must be done through it
var GlobalHttpClient = &http.Client{
	Timeout: time.Second * 10,
}
