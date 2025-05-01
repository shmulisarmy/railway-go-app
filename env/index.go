package env

import (
	"os"
	"strings"
)

func Load_env(env_file string) error {
	contents, err := os.ReadFile(env_file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		var key string
		var value string
		split_up := strings.Split(line, "")
		for i := range split_up {
			if split_up[i] == "=" {
				key = strings.TrimSpace(strings.Join(split_up[:i], ""))
				value = strings.TrimSpace(strings.Join(split_up[i+1:], ""))
				break
			}
		}
		os.Setenv(key, value)
	}
	return nil
}
