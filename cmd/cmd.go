package cmd

import (
	"todo/server"

	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	cli := &cobra.Command{
		Use: "todo",
	}

	cli.AddCommand(newHttpServerCmd())
	cli.AddCommand(newGrpcServerCmd())

	return cli
}

func newHttpServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "start http server",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartHttp()
		},
	}
}

func newGrpcServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "grpc",
		Short: "start grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartGRPC()
		},
	}
}
