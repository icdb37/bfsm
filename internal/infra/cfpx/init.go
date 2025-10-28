package cfpx

import (
	"github.com/icdb37/bfsm/internal/infra/config"
)

var pService *service

func Init() {
	path := config.GetCfpx()
	pService = &service{
		items: map[string]*Item{},
	}
	pService.load(path)
}
