package main

import (
	"crypto/x509"
	"strings"
	"strconv"
)

// To verify client HostName in client cert against the valid client HostName provided
// SAN or CommonName is matched
func VerifyClientHostName(cert *x509.Certificate, hostname string) error {

	h := strings.ToLower(hostname)
	if cert.DNSNames != nil {
		for _, name := range cert.DNSNames {
			if matchClientHostnames(strings.ToLower(name), h) {
				return nil
			}
		}
		// If Subject Alt Name (SAN) is given, we ignore the common name.
	} else if matchClientHostnames(strings.ToLower(cert.Subject.CommonName), h) {
		return nil
	}

	return x509.HostnameError{cert, h}
}

func matchClientHostnames(pattern, host string) bool {
	host = strings.TrimSuffix(host, ".")
	pattern = strings.TrimSuffix(pattern, ".")

	if len(pattern) == 0 || len(host) == 0 {
		return false
	}

	patternParts := strings.Split(pattern, ".")
	hostParts := strings.Split(host, ".")

	if len(patternParts) != len(hostParts) {
		return false
	}

	for i, patternPart := range patternParts {
		if i == 0 && patternPart == "*" && !IsInt(patternPart) && !IsInt(hostParts[i]) {
			continue
		} else if i == 0 && strings.HasPrefix(hostParts[i], "*") && !IsInt(patternPart) && !IsInt(hostParts[i]) {
			// If wildcard * present in the hostName provided
			// to be matched against the one given in cert.
			// Make sure the string after * matches the cert sub domain pattern
			tmpHostPart := strings.TrimPrefix(hostParts[i], "*")
			if strings.Contains(patternPart, tmpHostPart) {
				continue
			}
		} else if i == 0 &&  strings.HasSuffix(hostParts[i], "*") && !IsInt(patternPart) && !IsInt(hostParts[i]) {
			tmpHostPart := strings.TrimSuffix(hostParts[i], "*")
			if strings.Contains(patternPart, tmpHostPart) {
				continue
			}
		}
		if patternPart != hostParts[i] {
			return false
		}
	}

	return true
}

func IsInt(str string) bool {

    /** converting the str variable into an int **/
    _, err := strconv.Atoi(str)
    if err == nil {
        return true
    } else {
        return false
    }
}
