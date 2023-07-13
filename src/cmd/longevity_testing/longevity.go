package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/alcionai/corso/src/pkg/logger"
)

type backup struct {
	ID    string `json:"id"`
	Stats struct {
		ID      string    `json:"id"`
		EndedAt time.Time `json:"endedAt"`
	} `json:"stats"`
}

func main() {
	var backups []backup

	ctx := context.Background()
	service := os.Getenv("SERVICE")
	cmd := exec.Command(
		"./corso", "backup", "list", service, "--json", "--hide-progress")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fatal(ctx, "could not find backup", err)
	}

	jsonStart := bytes.Index(output, []byte("["))
	jsonEnd := bytes.LastIndex(output, []byte("]"))

	if jsonStart == -1 || jsonEnd == -1 || jsonEnd < jsonStart {
		fatal(ctx, "No valid JSON found in the output", nil)
	}

	jsonData := output[jsonStart : jsonEnd+1]

	if err := json.Unmarshal([]byte(jsonData), &backups); err != nil {
		fatal(ctx, "no service specified", nil)
	}

	days, err := strconv.Atoi(os.Getenv("DELETION_DAYS"))
	if err != nil {
		fatal(ctx, "invalid no of days provided", nil)
	}

	for _, backup := range backups {
		if backup.Stats.EndedAt.Before(time.Now().AddDate(0, 0, -days)) {
			cmd := exec.Command(
				"./corso", "backup", "delete", service, "--backup", backup.ID)
			err := cmd.Run()
			if err != nil {
				fatal(ctx, "deletion failed", nil)
			}
		}
	}

}

func fatal(ctx context.Context, msg string, err error) {
	logger.CtxErr(ctx, err).Error("test failure: " + msg)
	fmt.Println(msg+": ", err)
	os.Exit(1)
}
