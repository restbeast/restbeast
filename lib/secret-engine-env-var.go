package lib

import "os"

func secretEngineEnvVar(paths map[string]string) (secrets ParsedSecret) {
	secrets = make(ParsedSecret)
	for key, systemKey := range paths {
		secrets[key] = os.Getenv("restbeast_var_" + systemKey)
	}

	return secrets
}
