package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	secretsContent, err := os.ReadFile(filepath.Join(home, "/.mfa.json"))
	if err != nil {
		log.Fatal(err)
	}

	secrets := make(map[string]string)
	if err := json.Unmarshal(secretsContent, &secrets); err != nil {
		log.Fatal(err)
	}

	var outputs []string
	for account, secret := range secrets {
		token, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			outputs = append(outputs, fmt.Sprintf("%-18s failed(%s)", account, err))
			continue
		}
		outputs = append(outputs, fmt.Sprintf("%-18s %s", account, token))
	}
	sort.Strings(outputs)
	fmt.Println(strings.Join(outputs, "\n"))
}
