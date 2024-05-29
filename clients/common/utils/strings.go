package utils

import (
	"net/url"
	"strconv"
	"strings"
)

func Contains(stringList []string, st string) bool {
	for _, s := range stringList {
		if s == st {
			return true
		}
	}
	return false
}

func IsValidWSSUrl(u string) bool {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return false
	}

	// Check scheme
	if parsedUrl.Scheme != "wss" {
		return false
	}

	// Validate host (domain name or IP address)
	if parsedUrl.Host == "" {
		return false
	}

	// If a port is specified, it should be a valid number
	if strings.Contains(parsedUrl.Host, ":") {
		hostParts := strings.Split(parsedUrl.Host, ":")
		if len(hostParts) != 2 {
			return false
		}
		port := hostParts[1]
		if portNumber, err := strconv.Atoi(port); err != nil || portNumber <= 0 || portNumber > 65535 {
			return false
		}
	}

	// If all checks pass, the URL is valid
	return true
}
