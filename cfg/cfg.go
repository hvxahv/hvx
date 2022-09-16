package cfg

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

const (
	ConsulProvider = "consul"
)

type Option func(o *options)

type options struct {
	provider  string
	endpoint  string
	path      string
	cfgType   string
	watchTime time.Duration
	secure    string
}

type Cfg struct {
	opts *options
}

type cfg interface {
	Dial()
}

func WithConfigType(ct string) Option {
	return func(o *options) {
		o.cfgType = ct
	}
}

func WithWatchTime(time time.Duration) Option {
	return func(o *options) {
		o.watchTime = time
	}
}

func WithSecureProvider(key string) Option {
	return func(o *options) {
		o.secure = key
	}
}

func New(provider, endpoint, path string, opts ...Option) *Cfg {
	o := &options{
		provider: provider,
		endpoint: endpoint,
		path:     path,
	}
	for _, opt := range opts {
		opt(o)
	}
	switch {
	case o.cfgType == "":
		o.cfgType = "yaml"
	case o.watchTime == 0:
		o.watchTime = time.Second * 5
	default:
	}
	return &Cfg{
		opts: o,
	}
}

func (cfg *Cfg) Dial() error {
	if cfg.opts.secure != "" {
		if err := viper.AddSecureRemoteProvider(cfg.opts.provider, cfg.opts.endpoint, cfg.opts.path, cfg.opts.secure); err != nil {
			return err
		}
	} else {
		if err := viper.AddRemoteProvider(cfg.opts.provider, cfg.opts.endpoint, cfg.opts.path); err != nil {
			return err
		}
	}
	viper.SetConfigType(cfg.opts.cfgType)
	if err := viper.ReadRemoteConfig(); err != nil {
		return err
	}
	if err := viper.ReadInConfig(); err == nil {
		return err
	}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			if err := viper.WatchRemoteConfig(); err != nil {
				fmt.Errorf("unable to read remote config: %v", err)
			}
		}
	}()

	return nil
}
