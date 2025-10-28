package cfpx

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"

	"github.com/icdb37/bfsm/internal/infra/logx"
	"go.yaml.in/yaml/v3"
)

// Featurer 功能参数接口
type Featurer interface {
	GetFeature() string
}

// Item 节点
type Item struct {
	Elem `json:",inline" yaml:",inline"`
	Item map[string]*Item `json:"item" yaml:"item"`
}

// feature 功能
type service struct {
	items map[string]*Item `json:"feats" yaml:"feats"`
}

func (s *service) load(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		logx.Fatal("read 'cfpx' config file failed", "path", path, "err", err)
	}
	s.items = map[string]*Item{}
	var fUnmarshal func(data []byte, v any) error
	if strings.HasSuffix(path, "json") || strings.HasSuffix(path, "js") {
		fUnmarshal = json.Unmarshal
	} else {
		fUnmarshal = yaml.Unmarshal
	}
	if err := fUnmarshal(data, s.items); err != nil {
		logx.Fatal("unmarshal 'cfpx' config file failed", "path", path, "err", err)
	}
}

func (s *service) Process(param Featurer) error {
	v := reflect.ValueOf(param)
	if v.Kind() != reflect.Ptr {
		return nil
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}
	features := []string{param.GetFeature(), ""}
	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		iv := v.Field(i)
		f := v.Type().Field(i)
		if f.Anonymous {
			continue
		}
		// 检查是否有 cfpx 标签
		tag := f.Tag.Get("cfpx")
		if tag == "" {
			continue
		}
		tag = strings.TrimSpace(tag)
		pn := parseTagElem(tag)
		features[len(features)-1] = pn.Code
		if tpn := s.getElem(s.items, features...); tpn != nil {
			pn = tpn
		}
		processFmtfn(&iv, pn)
		if err := processCheck(&iv, pn); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) getElem(items map[string]*Item, paths ...string) *Elem {
	for i, p := range paths {
		item, ok := items[p]
		if !ok {
			return nil
		}
		if len(paths) == i+1 {
			return &(item.Elem)
		}
		items = item.Item
	}
	return s.getDefaultElem(items, paths...)
}

func (s *service) getDefaultElem(items map[string]*Item, paths ...string) *Elem {
	paths[0] = "_default_"
	return s.getElem(items, paths...)
}
