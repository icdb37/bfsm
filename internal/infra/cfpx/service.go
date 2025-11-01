package cfpx

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/icdb37/bfsm/internal/infra/logx"
)

const (
	defaultFeature = "_default_"
)

// Featurer 功能参数接口
type Featurer interface {
	GetFeature() string
}

// Item 节点
type Item struct {
	Elem   `json:",inline" yaml:",inline"`
	Item   map[string]*Item `json:"item" yaml:"item"`
	parent *Item            `json:"-" yaml:"-"`
}

func (i *Item) getParentCodes() []string {
	if i.parent == nil {
		return []string{}
	}
	return append(i.parent.getParentCodes(), i.parent.Code)
}
func (i *Item) GetField() string {
	codes := append(i.getParentCodes(), i.Code)
	return strings.Join(codes[1:], ".")
}

// feature 功能
type service struct {
	itemFeature *Item
	itemDefault *Item
}

func (s *service) load(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		logx.Fatal("read 'cfpx' config file failed", "path", path, "err", err)
	}
	var fUnmarshal func(data []byte, v any) error
	if strings.HasSuffix(path, "json") || strings.HasSuffix(path, "js") {
		fUnmarshal = json.Unmarshal
	} else {
		fUnmarshal = yaml.Unmarshal
	}
	if err := fUnmarshal(data, s.itemFeature.Item); err != nil {
		logx.Fatal("unmarshal 'cfpx' config file failed", "path", path, "err", err)
	}
	if _, ok := s.itemFeature.Item[defaultFeature]; ok {
		s.itemDefault = s.itemFeature.Item[defaultFeature]
		delete(s.itemFeature.Item, defaultFeature)
	}
	for _, i := range s.itemFeature.Item {
		for _, si := range i.Item {
			s.initItemParent(si)
		}
	}
	s.initItemParent(s.itemDefault)
}
func (s *service) initItemParent(item *Item) {
	for _, subItem := range item.Item {
		subItem.parent = item
		s.initItemParent(subItem)
	}
}

// Process 修改数据格式化校验
func (s *service) Process(param Featurer) error {
	v := reflect.ValueOf(param)
	if v.Kind() != reflect.Ptr {
		return nil
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}
	features := strings.Split(param.GetFeature(), ".")
	pi := s.itemFeature
	for _, f := range features {
		if pi == nil {
			return nil
		}
		if item, ok := pi.Item[f]; ok {
			pi = item
			continue
		}
		pi = s.itemFeature.Item[defaultFeature]
	}
	return s.process(&v, pi)
}
func (s *service) process(pv *reflect.Value, parentItem *Item) error {
	if parentItem == nil {
		return nil
	}
	if len(parentItem.Item) == 0 {
		processFmtfn(pv, parentItem)
		if err := processCheck(pv, parentItem); err != nil {
			return err
		}
		return nil
	}
	switch pv.Kind() {
	case reflect.Ptr:
		iv := pv.Elem()
		return s.process(&iv, parentItem)
	case reflect.Struct:
		for i := 0; i < pv.NumField(); i++ {
			iv := pv.Field(i)
			f := pv.Type().Field(i)
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
			subxItem, ok := parentItem.Item[pn.Code]
			if !ok { //从默认校验匹配
				subxItem = s.itemDefault.Item[pn.Code]
			}
			if subxItem == nil {
				subxItem = pn
			}
			if err := s.process(&iv, subxItem); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, key := range pv.MapKeys() {
			iv := pv.MapIndex(key)
			subxItem, ok := parentItem.Item[key.String()]
			if !ok {
				continue
			}
			if err := s.process(&iv, subxItem); err != nil {
				return err
			}
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < pv.Len(); i++ {
			iv := pv.Index(i)
			if err := s.process(&iv, parentItem); err != nil {
				return err
			}
		}
	}
	return nil
}
