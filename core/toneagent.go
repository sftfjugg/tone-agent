package core

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"tone-agent/comm"
)

func GetToneAgentPath() (string, error) {
	toneAgentPath := os.Getenv("TONE_AGENT_PATH")
	if toneAgentPath != "" {
		absPath, err := filepath.Abs(toneAgentPath)
		if err != nil {
			toneAgentPath = ""
		} else {
			toneAgentPath = absPath
		}
	}
	if toneAgentPath != "" {
		return toneAgentPath, nil
	}
	user, err := comm.Home()
	if err != nil {
		log.Printf("Fetch directory of system user failed: %v", err)
		return "", err
	}
	return path.Join(user, "ToneAgent"), nil
}
