package middleware

import (
	"gwftemplate/data"

	"github.com/ck3g/gwf"
)

type Middleware struct {
	App    *gwf.GWF
	Models data.Models
}
