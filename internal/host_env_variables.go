package internal

import "os"

// HostEnvVariables is a struct provides the functionality to get environment variables from the host.
type HostEnvVariables struct{}

func (host *HostEnvVariables) getAll() []string {
	return os.Environ()
}

func (host *HostEnvVariables) get(key string) string {
	return os.Getenv(key)
}
