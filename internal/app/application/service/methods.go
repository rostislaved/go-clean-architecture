package service

import (
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/order"
)

func (svc *ApiService) Start() {
	ticker := time.NewTicker(svc.config.UpdateInterval)

	for ; true; <-ticker.C {
		err := svc.loop()
		if err != nil {
		}
	}
}

func (svc *ApiService) loop() (err error) {
	return nil
}

func (svc *ApiService) Method1([]order.Order) (interface{}, error) {
	return nil, nil
}
