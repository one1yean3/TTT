package service

import (
	"context"
	"fmt"
)

func (wr WebhookService) WebhookService() error {
	data, err := wr.storage.FindAll(context.TODO())
	if err != nil {
		return err
	}

	fmt.Println("TEST")
	fmt.Println("TEST")
	fmt.Println("TEST")
	fmt.Println("TEST")
	fmt.Println("data = ", data)
	return nil
}
