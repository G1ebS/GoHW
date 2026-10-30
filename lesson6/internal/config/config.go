package config

import "os"

var HTTPAddr = os.Getenv("HTTP_ADDR")
