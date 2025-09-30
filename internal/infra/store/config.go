package store

// Config 存储配置
type Config struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Debug bool   `json:"debug"`
}
