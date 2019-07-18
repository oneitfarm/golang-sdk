package golang_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/oneitfarm/golang-sdk/cienv"
	"sync"
)

var sdk *Sdk
var sdkInitOnce = sync.Once{}

func GetSdk() *Sdk {
	sdkInitOnce.Do(func() {
		sdk = &Sdk{
			services:   map[string]string{},
			gatewayUrl: "",
		}
		sdk.init()
	})
	return sdk
}

type Sdk struct {
	services   map[string]string
	gatewayUrl string
}

func (s *Sdk) init() {
	str := cienv.GetEnv("services")
	if str != "" {
		j := map[string]interface{}{}
		if err := json.Unmarshal([]byte(str), &j); err != nil {
			return
		}
		for k, item := range j {
			s.services[k] = item.(string)
		}
	}
}

func (s *Sdk) GetServiceUrl(serviceName string) string {
	url, exists := s.services[serviceName]
	if exists {
		return url
	}

	if s.gatewayUrl != "" {
		url = fmt.Sprintf("%s%s/", s.gatewayUrl, serviceName)
		s.services[serviceName] = url
		return url
	}

	url = cienv.GetEnv(fmt.Sprintf("DEPLOYMENT_%s_HOST"))
	if url != "" {
		s.services[serviceName] = url
		return url
	}

	url = cienv.GetEnv(fmt.Sprintf("WORKSPACE_%s_HOST"))
	if url != "" {
		s.services[serviceName] = url
		return url
	}

	return ""
}
