package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type cartWrapper struct {
	client.Client
}

func (wrapper *cartWrapper) Call(ctx context.Context, req client.Request, rsp interface {
}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                3000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  50,
		SleepWindow:            5000,
	}
	hystrix.ConfigureCommand(cmdName, config)

	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		return err
	})
}

func NewCartWrapper(c client.Client) client.Client {
	return &cartWrapper{c}
}
