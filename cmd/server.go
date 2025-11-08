package cmd

import (
	"context"

	webBill "github.com/icdb37/bfsm/internal/features/bill/api/web"
	webCommodity "github.com/icdb37/bfsm/internal/features/commodity/api/web"
	webCompany "github.com/icdb37/bfsm/internal/features/company/api/web"
	webInventory "github.com/icdb37/bfsm/internal/features/inventory/api/web"
	webPurchase "github.com/icdb37/bfsm/internal/features/purchase/api/web"
	webUser "github.com/icdb37/bfsm/internal/features/user/api/web"
	"github.com/icdb37/bfsm/internal/wire/echox"

	"github.com/spf13/cobra"

	svcBill "github.com/icdb37/bfsm/internal/features/bill/service"
	svcCommodity "github.com/icdb37/bfsm/internal/features/commodity/service"
	svcCompany "github.com/icdb37/bfsm/internal/features/company/service"
	svcInventory "github.com/icdb37/bfsm/internal/features/inventory/service"
	svcPurchase "github.com/icdb37/bfsm/internal/features/purchase/service"
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
	svcInventory.Provide()
	svcBill.Provide()
	svcCompany.Provide()
	svcUser.Provide()
	svcCommodity.Provide()
	svcPurchase.Provide()
}
func wireWeb() {
	webInventory.Wire()
	webCompany.Wire()
	webUser.Wire()
	webCommodity.Wire()
	webPurchase.Wire()
	webBill.Wire()
}
