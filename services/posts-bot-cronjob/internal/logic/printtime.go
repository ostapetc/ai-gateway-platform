package logic

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func PrintTime(_ *cobra.Command, _ []string) error {
	fmt.Println(time.Now().Format(time.RFC3339))
	return nil
}
