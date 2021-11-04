package jwt

import (
	"github.com/Ankr-network/kit/util"
	"sync"
	"time"
)

var (
	out      *Config
	confOnce sync.Once
	Sign     Signer
)

type Config struct {
	Iss             string        `env:"JWT_ISSUER" envDefault:"blue-dashboard"`
	ExpiredDuration time.Duration `env:"JWT_EXPIRED_TIME" envDefault:"24h"`
	Key             string        `env:"JWT_KEY" envDefault:"blue-dashboard"`
}

func MustLoadConfig() *Config {
	confOnce.Do(func() {
		out = new(Config)
		util.MustLoadConfig(out)
	})

	return out
}

func InitGlobal(config *Config) {
	var err error
	Sign, err = NewSigner(SignKey(config.Key), Iss(config.Iss), ExpireDuration(config.ExpiredDuration))
	if err != nil {
		panic(err)
	}
}
