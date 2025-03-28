package test

import (
	"fmt"
	"github.com/spf13/cobra"
)

var name *string

func init() {
  name = HelloCmd.Flags().StringP("name", "n", "world", "이름을 적어주세요!")
}

var HelloCmd = &cobra.Command{
  Use: "hello",
  Short: "안녕!",
  Long: "안녕!이라고 말함",
  Run: func(cmd *cobra.Command, args []string) {

    fmt.Println("안녕!")
    if name != nil {
      fmt.Println(*name, "!")
    }
  },
}

var Test string
