package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootFlag = struct {
	Start      string
	End        string
	Expression string
	Separator  string
	Prefix     string
	Dir        string
}{}

var (
	rootCmd = &cobra.Command{
		Use:   "glog",
		Short: "glog",
		Run: func(cmd *cobra.Command, args []string) {
			fs, err := convertTimeToFileList(rootFlag.Prefix, rootFlag.Start, rootFlag.End)
			if err != nil {
				Fatal(err)
			}
			for _, item := range fs {
				if _, err := os.Stat(item); os.IsNotExist(err) {
					continue
				}
				if err := grep(item, rootFlag.Expression); err != nil {
					Fatal(err)
				}
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&rootFlag.Prefix, "prefix", "api", "file prefix")
	rootCmd.PersistentFlags().StringVar(&rootFlag.Start, "start", time.Now().Format("2006-01-02 15:04:05"), "start time")
	rootCmd.PersistentFlags().StringVar(&rootFlag.End, "end", time.Now().Format("2006-01-02 15:04:05"), "end time")
	rootCmd.PersistentFlags().StringVar(&rootFlag.Expression, "expr", "", "regular expression")
	rootCmd.PersistentFlags().StringVar(&rootFlag.Separator, "separator", "_", "separator")
	rootCmd.PersistentFlags().StringVar(&rootFlag.Dir, "dir", ".", "working dir")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
