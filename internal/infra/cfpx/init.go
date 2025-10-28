package cfpx

import (
	"github.com/icdb37/bfsm/internal/infra/config"
)

var pService *service

func Init() {
	path := config.GetCfpx()
	pService = &service{
		itemFeature: &Item{
			Item: make(map[string]*Item),
		},
		itemDefault: &Item{
			Item: make(map[string]*Item),
		},
	}
	pService.load(path)
}
