package configs

import (
	"os"
)

var PORT = os.Getenv("GDP_API_PORT")
