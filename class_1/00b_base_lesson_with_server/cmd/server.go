package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tj10200/golang-class/class_1/00b_base_lesson_with_server/pkg/bluecore"
)

var addr string
var port int

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "boiler plate for golang server",
		Long:  `boiler plate for golang server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			router := gin.Default()
			router.GET("/hello/:world", bluecore.HelloWorldHandler)
			return router.Run(fmt.Sprintf("%s:%d", addr, port))
		},
	}

	cmd.PersistentFlags().StringVar(&addr, "addr", "0.0.0.0", "The host address to serve against")
	cmd.PersistentFlags().IntVar(&port, "port", 9876, "the server port")

	return cmd
}
