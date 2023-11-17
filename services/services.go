package services

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Interface interface {
	GetDeliveryTime() (int, error)
}

type Service struct {
}

func NewServices() *Service {
	return &Service{}
}

func (s *Service) GetDeliveryTime() (int, error) {
	return rand.Intn(61), nil
	// because this api is not working i return some random number between 0 and 60
	resp, err := http.Get("https://run.mocky.io/v3/122c2796-5df4-461c-ab75-87c1192b17f7")
	if err != nil {
		return 0, fmt.Errorf("error while getting delivery time: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error while getting delivery time: %v", err)
	}
	return 0, nil
}
