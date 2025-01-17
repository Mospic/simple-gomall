package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type userWrapper struct {
	client.Client
}

func (wrapper *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
<<<<<<< HEAD
		Timeout:                30000,
		RequestVolumeThreshold: 20,   //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
=======
		Timeout:                3000000,
		RequestVolumeThreshold: 20,     //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,     //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            500000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
>>>>>>> c17bcb1179efba304c408856318b693d3faf3def
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

// NewUserWrapper 初始化Wrapper
func NewUserWrapper(c client.Client) client.Client {
	return &userWrapper{c}
}
