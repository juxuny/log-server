package main

import (
	"fmt"
	"github.com/juxuny/log-server/log"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

var rootFlag = struct {
	KeepDuration  string
	Dir           string
	Scripts       string
	ScriptsFile   string
	Ext           string
	CheckDuration string
}{}

func walk(scriptContent string, keepDuration time.Duration) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	currentTime := time.Now()
	fmt.Println("check at: ", currentTime.Format("2006-01-02 15:04:05"))
	if err := filepath.Walk(rootFlag.Dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		var env = []string{
			fmt.Sprintf("FILE_NAME=%s", info.Name()),
			fmt.Sprintf("FULL_NAME=%s", filepath.Join(path, info.Name())),
			fmt.Sprintf("FILE_EXT=%s", strings.Trim(filepath.Ext(info.Name()), ".")),
		}
		d := currentTime.Sub(info.ModTime())
		ext := strings.Trim(filepath.Ext(info.Name()), ".")
		if d > keepDuration && ext == rootFlag.Ext {
			if err := execShell(scriptContent, env); err != nil {
				fmt.Println(err)
			}
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use: "cleaner",
	Run: func(cmd *cobra.Command, args []string) {
		duration, err := parseDuration(rootFlag.KeepDuration)
		if err != nil {
			log.Fatal(err)
		}
		rootFlag.Ext = strings.Trim(rootFlag.Ext, ".")
		if rootFlag.Ext == "" {
			log.Fatal("ext cannot empty")
		}
		scriptContent := ""
		if rootFlag.Scripts != "" {
			scriptContent = rootFlag.Scripts
		} else if rootFlag.ScriptsFile != "" {
			scriptData, err := ioutil.ReadFile(rootFlag.ScriptsFile)
			if err != nil {
				log.Fatal(err)
			}
			scriptContent = string(scriptData)
		} else {
			log.Fatal("--script and --script-file are empty")
		}
		checkDuration, err := parseDuration(rootFlag.CheckDuration)
		if err != nil {
			log.Fatal(err)
		}
		ticker := time.NewTicker(checkDuration)
		for range ticker.C {
			walk(scriptContent, duration)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootFlag.KeepDuration, "keep", "30d", "s => seconds, m => minutes, h => hour, d => days, w => weeks")
	rootCmd.PersistentFlags().StringVar(&rootFlag.Dir, "dir", ".", "")
	rootCmd.PersistentFlags().StringVar(&rootFlag.Scripts, "script", "", "shell command, env variable: FILE_NAME, FULL_NAME, FILE_EXT")
	rootCmd.PersistentFlags().StringVar(&rootFlag.ScriptsFile, "script-file", "", "shell script file path")
	rootCmd.PersistentFlags().StringVar(&rootFlag.Ext, "ext", "log", "file extension name")
	rootCmd.PersistentFlags().StringVar(&rootFlag.CheckDuration, "check", "1m", "check file every 1 minute")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
