package cmd

import (
	webUser "github.com/icdb37/bfsm/internal/features/user/api/web"
	"github.com/icdb37/bfsm/internal/infra/config"
	"github.com/icdb37/bfsm/internal/infra/store/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.ConfigFile, "config", "f", "", "config file")
	cobra.OnInitialize(config.MustInitConfig)
	rootCmd.AddCommand(
		ServerCmd,
	)
	sqlite.Init()
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Init server",
	Long:  `Init server`,
	RunE: func(_ *cobra.Command, _ []string) error {
		e := echo.New()
		if err := webUser.Init(e); err != nil {
			return err
		}
		if err := e.Start(":8080"); err != nil {
			return err
		}
		return nil
	},
}
