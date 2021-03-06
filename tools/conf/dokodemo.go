package conf

import (
	"v2ray.com/core/common/loader"
	"v2ray.com/core/proxy/dokodemo"
)

type DokodemoConfig struct {
	Host         *Address     `json:"address"`
	PortValue    uint16       `json:"port"`
	NetworkList  *NetworkList `json:"network"`
	TimeoutValue uint32       `json:"timeout"`
	Redirect     bool         `json:"followRedirect"`
}

func (v *DokodemoConfig) Build() (*loader.TypedSettings, error) {
	config := new(dokodemo.Config)
	if v.Host != nil {
		config.Address = v.Host.Build()
	}
	config.Port = uint32(v.PortValue)
	config.NetworkList = v.NetworkList.Build()
	config.Timeout = v.TimeoutValue
	config.FollowRedirect = v.Redirect
	return loader.NewTypedSettings(config), nil
}
