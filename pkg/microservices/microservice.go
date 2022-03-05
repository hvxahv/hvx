package microservices

import (
	"fmt"
	"github.com/spf13/viper"
)

type Service struct {
	Name string
}

func (s *Service) GetHost() string {
	return viper.GetString(fmt.Sprintf("microservices.%s.host", s.Name))
}

func (s *Service) GetPort() string {
	return viper.GetString(fmt.Sprintf("microservices.%s.port", s.Name))
}

func (s *Service) GetAddress() string {
	return viper.GetString(fmt.Sprintf("%s:%s", s.GetHost(), s.GetPort()))
}

func NewService(name string) *Service {
	return &Service{Name: name}
}

type Microservice interface {
	GetServiceHost() string
	GetServicePort() string
	GetServiceAddress() string
}
