package wrapper

import (
	"fmt"
	"gateway/common/response"
	"github.com/pkg/errors"
	"time"
)

func NewServiceWrapper(name string) {
	c := &Config{
		Namespace:              name,
		Timeout:                1 * time.Second, // TODO 建议加在配置文件里面
		MaxConcurrentRequests:  100,
		RequestVolumeThreshold: 10,
		SleepWindow:            5 * time.Second,
		ErrorPercentThreshold:  50,
	}

	g := NewGroup(c)
	if err := g.Do(name, func() error {
		return errors.New(response.GetMsg(response.ErrorServiceUnavailable))
	}); err != nil {
		fmt.Println("err", err)
	}
}
