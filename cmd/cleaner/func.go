package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func parseDuration(s string) (time.Duration, error) {
	s = strings.ToLower(s)
	b := time.Second
	if strings.Contains(s, "s") {
		b = time.Second
		s = strings.ReplaceAll(s, "s", "")
	} else if strings.Contains(s, "m") {
		b = time.Minute
		s = strings.ReplaceAll(s, "m", "")
	} else if strings.Contains(s, "h") {
		b = time.Hour
		s = strings.ReplaceAll(s, "h", "")
	} else if strings.Contains(s, "d") {
		b = time.Hour * 24
		s = strings.ReplaceAll(s, "d", "")
	} else if strings.Contains(s, "w") {
		b = time.Hour * 24 * 7
		s = strings.ReplaceAll(s, "w", "")
	} else {
		return 0, fmt.Errorf("unknown duration format: %v", s)
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid duration: %v", s)
	}

	return time.Duration(v) * b, nil
}

func execShell(script string, env []string) error {
	cmd := exec.Command("sh")
	cmd.Env = env
	cmd.Dir = rootFlag.Dir
	in := bytes.NewBuffer([]byte(script))
	cmd.Stdin = in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
