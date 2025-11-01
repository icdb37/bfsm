package cmd

import (
	"context"

	webCommodity "github.com/icdb37/bfsm/internal/features/commodity/api/web"
	webCompany "github.com/icdb37/bfsm/internal/features/company/api/web"
	webUser "github.com/icdb37/bfsm/internal/features/user/api/web"
	"github.com/icdb37/bfsm/internal/wire/echox"

	"github.com/spf13/cobra"

	svcCommodity "github.com/icdb37/bfsm/internal/features/commodity/service"
	svcCompany "github.com/icdb37/bfsm/internal/features/company/service"
	svcUser "github.com/icdb37/bfsm/internal/features/user/service"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
	"github.com/icdb37/bfsm/internal/infra/config"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store/sqlite"
	"github.com/icdb37/bfsm/internal/wire"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.ConfigFile, "config", "f", "", "config file")
	cobra.OnInitialize(config.MustInitConfig)
	rootCmd.AddCommand(
		ServerCmd,
	)
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Init server",
	Long:  `Init server`,
	RunE: func(_ *cobra.Command, _ []string) error {
		initInfra()
		provideService()
		wireWeb()
		wire.Start(context.Background())
		return nil
	},
}

func initInfra() {
	sqlite.Init()
	cfpx.Init()
	logx.Init()
}
func provideService() {
	echox.Provide()
	svcCompany.Provide()
	svcUser.Provide()
	svcCommodity.Provide()
}
func wireWeb() {
	webCompany.Wire()
	webUser.Wire()
	webCommodity.Wire()
}
