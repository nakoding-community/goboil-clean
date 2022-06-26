package response

import "github.com/nakoding-community/goboil-clean/internal/abstraction"

type Meta struct {
	Success bool                        `json:"success" default:"true"`
	Message string                      `json:"message" default:"true"`
	Info    *abstraction.PaginationInfo `json:"info"`
}
