package middleware

import (
	"fmt"

	"github.com/okaaryanata/public-api/internal/domain"
)

const (
	Teapot     = "/teapot"
	HealthPath = "/health"
)

var (
	SkipPaths = []string{
		Teapot,
		HealthPath,
	}
)

func GetListSkipLogPath() []string {
	template := "%s%s"

	var listSkip []string
	for _, r := range SkipPaths {
		route := fmt.Sprintf(template, domain.MainRoute, r)
		listSkip = append(listSkip, route)
	}

	return listSkip
}
