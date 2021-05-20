package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

func parseTime(s string) (ret time.Time, err error) {
	hyphenIndex := strings.Index(s, "-")
	colonIndex := strings.Index(s, ":")
	spaceIndex := strings.Index(s, " ")
	if hyphenIndex > 0 && colonIndex > 0 {
		return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	}
	if hyphenIndex > 0 && spaceIndex > 0 {
		return time.ParseInLocation("2006-01-02 15", s, time.Local)
	}
	if len(s) == len("2006010215") {
		return time.ParseInLocation("2006010215", s, time.Local)
	}

	return ret, fmt.Errorf("invalid time: %s", s)
}

func convertTimeToFileList(prefix string, start, end string) (ret []string, err error) {
	s, err := parseTime(start)
	if err != nil {
		return nil, err
	}
	e, err := parseTime(end)
	if err != nil {
		return nil, err
	}
	if e.Before(s) {
		return nil, fmt.Errorf("invalid time duration")
	}
	for s.Format("2006010215") <= e.Format("2006010215") {
		ret = append(ret, path.Join(rootFlag.Dir, strings.Join([]string{prefix, s.Format("20060102_15") + ".log"}, rootFlag.Separator)))
		s = s.Add(time.Hour)
	}
	return ret, nil
}

func grep(fileName, expr string) error {
	sh := fmt.Sprintf("cat %s | grep \"%s\"", fileName, expr)
	cmd := exec.Command("sh")
	in := bytes.NewBuffer([]byte(sh))
	cmd.Stdin = in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
