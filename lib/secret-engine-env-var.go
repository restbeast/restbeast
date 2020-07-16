package lib

import "os"

func secretEngineEnvVar(paths map[string]string) (secrets map[string]string) {
	secrets = make(map[string]string)
	for key, systemKey := range paths {
		secrets[key] = os.Getenv("restbeast_var_" + systemKey)
	}

	return secrets
}
