package logx

import (
	"encoding/json"

	"github.com/icdb37/bfsm/internal/infra/config"
)

func Init() {
	strConfig := config.GetLogx()
	opts := &Options{}
	if err := json.Unmarshal([]byte(strConfig), opts); err != nil {
		return
	}
	if sugared, err := newZapSugared(opts); err == nil {
		log = newZapLogger(sugared)
		emptyWithLog = log.With()
	}
}
