package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	// packed := Encode(data)
	packed := ""

	err = io.WriteFile(packedFileName(), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}
}

func packedFileName()
