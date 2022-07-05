package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/tj10200/golang-class/class_1/02_data_processing/pkg/bluecore"
)

var addr string
var port int
var filename string

func NewLesson() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lesson",
		Short: "boiler plate for golang server",
		Long:  `boiler plate for golang server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return bluecore.Lesson(bluecore.Config{
				FilePath: filename,
				Fs:       afero.NewOsFs(),
			})
		},
	}

	cmd.PersistentFlags().StringVar(&addr, "addr", "0.0.0.0", "The host address to serve against")
	cmd.PersistentFlags().IntVar(&port, "port", 9876, "the server port")
	cmd.PersistentFlags().StringVar(&filename, "file_name", "", "the file to process")

	return cmd
}
