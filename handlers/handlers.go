package handlers

import (
	"demoapp/data"
	"net/http"

	"github.com/ck3g/gwf"
)

type Handlers struct {
	App    *gwf.GWF
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
