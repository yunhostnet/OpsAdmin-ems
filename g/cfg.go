package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type HttpConfig struct {
	Enable bool   `json:"enable"`
	Listen string `json:"listen"`
}
type MailConfig struct {
	Enable            bool   `json:"enable"`
	SendConcurrent    int    `json:"sendConcurrent"`
	MaxQueueSize      int    `json:"maxQueueSize"`
	FromUser          string `json:"fromUser"`
	MailServerHost    string `json:"mailServerHost"`
	MailServerPort    int    `json:"mailServerPort"`
	MailServerAccount string `json:"mailServerAccount"`
	MailServerPasswd  string `json:"mailServerPasswd"`
}
type WechatConfig struct {
	Enable         bool   `json:"enable"`
	SendConcurrent int    `json:"sendConcurrent"`
	MaxQueueSize   int    `json:"maxQueueSize"`
	Url            string `json:"url"`
	Ak             string `json:"ak"`
	Sk             string `json:"sk"`
}

type GlobalConfig struct {
	Debug  bool          `json:"debug"`
	Http   *HttpConfig   `json:"http"`
	Mail   *MailConfig   `json:"mail"`
	Wechat *WechatConfig `json:"wechat"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func GetConfig() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func LoadConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var v GlobalConfig
	err = json.Unmarshal([]byte(configContent), &v)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &v

	log.Println("g.ParseConfig ok, file ", cfg)
}
