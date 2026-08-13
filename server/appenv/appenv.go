package appenv

import (
	"os"
	"strings"
)

type Environment string

const (
	Development Environment = "development"
	Test        Environment = "test"
	Staging     Environment = "staging"
	Production  Environment = "production"
)

func Parse(value string) Environment {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case string(Development):
		return Development
	case string(Test):
		return Test
	case string(Staging):
		return Staging
	case string(Production):
		return Production
	default:
		return Development
	}
}

func Current() Environment {
	return Parse(os.Getenv("APP_ENV"))
}

func IsProductionLike(env Environment) bool {
	return env == Staging || env == Production
}

func ShouldUseReleaseMode(ginMode string, env Environment) bool {
	switch strings.ToLower(strings.TrimSpace(ginMode)) {
	case "debug", string(Development):
		return false
	case "release", string(Production):
		return true
	default:
		return IsProductionLike(env)
	}
}

func Port(env Environment) string {
	switch env {
	case Test:
		return "3003"
	case Staging:
		return "3004"
	case Production:
		return "3005"
	default:
		return "3002"
	}
}
