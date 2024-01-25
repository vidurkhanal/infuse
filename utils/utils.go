package utils

import (
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/vidurkhanal/infuse/config"
	"github.com/vidurkhanal/infuse/constants"
)

func ConstructConfig(c *fiber.Ctx) (config.Target, error) {
	var config config.Target
	if err := json.Unmarshal([]byte(
		c.Get(fmt.Sprintf("x-%s-config",
			constants.POWERED_BY))), &config); err != nil {
		return config, err
	}

	if config.Provider == "" && len(config.Targets) == 0 {
		config.Provider = c.Get(fmt.Sprintf("x-%s-mode", constants.POWERED_BY))
		config.ApiKey = getBearerToken(c.Get("authorization"))
		if config.Provider == string(constants.AZURE_OPEN_AI) {
			config.ResourceName = c.Get(fmt.Sprintf("x-%s-azure-resource-name", constants.POWERED_BY))
			config.DeploymentId = c.Get(fmt.Sprintf("x-%s-azure-deployment-id", constants.POWERED_BY))
			config.ApiVersion = c.Get(fmt.Sprintf("x-%s-azure-api-version", constants.POWERED_BY))
		}
	}

	return config, nil
}

const authHeaderPrefix = "Bearer "

func getBearerToken(authHeader string) string {
	// Check if header has prefix
	if !strings.HasPrefix(authHeader, authHeaderPrefix) {
		return ""
	}

	// Extract token
	bearerToken := strings.TrimPrefix(authHeader, authHeaderPrefix)

	// Save token
	// (e.g. to context or database)

	return bearerToken
}
