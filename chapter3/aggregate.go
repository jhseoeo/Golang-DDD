package chapter3

import (
	"errors"
	"github.com/google/uuid"
)

type WalletItem interface {
	GetBalance() (Money, error)
}

type Wallet struct {
	id          uuid.UUID
	ownerId     uuid.UUID
	walletItems []WalletItem
}

func (w Wallet) GetWalletBalance() (*Money, error) {
	var result Money
	for _, v := range w.walletItems {
		itemBal, err := v.GetBalance()
		if err != nil {
			return nil, errors.New("failed to get balance")
		}
		result += itemBal
	}
	return &result, nil
}
