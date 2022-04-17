package microservices

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

type Service struct {
	Name string
}

func (s *Service) GetHost() string {
	return viper.GetString(fmt.Sprintf("microservices.%s.host", s.Name))
}

func (s *Service) GetGRPCPort() int {
	p, err := strconv.Atoi(viper.GetString(fmt.Sprintf("microservices.%s.port", s.Name)))
	if err != nil {
		return 0
	}
	return p
}

func (s *Service) GetHTTPPort() int {
	p, err := strconv.Atoi(viper.GetString(fmt.Sprintf("microservices.%s.http_port", s.Name)))
	if err != nil {
		return 0
	}
	return p
}

func (s *Service) GetGRPCAddress() string {
	return fmt.Sprintf("%s:%s", s.GetHost(), strconv.Itoa(s.GetGRPCPort()))
}

func (s *Service) GetHTTPAddress() string {
	return fmt.Sprintf("%s:%s", s.GetHost(), strconv.Itoa(s.GetHTTPPort()))
}

func NewService(name string) *Service {
	return &Service{Name: name}
}

type Microservice interface {
	GetServiceHost() string
	GetServicePort() string
	GetServiceAddress() string
}
