package purchase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	coffeeco "github.com/jhseoeo/Golang-DDD/chapter5/internal"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/membership"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/payment"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/store"
	"time"
)

type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	total              coffeeco.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("puchase must have at least one product")
	}

	p.total = 0
	for _, v := range p.ProductsToPurchase {
		p.total += v.BasePrice
	}

	if p.total == 0 {
		return errors.New("total price must be greater than 0")
	}

	p.id = uuid.New()
	p.timeOfPurchase = time.Now()
	return nil
}

type CardChargeService interface {
	ChargeCard(ctx context.Context, amount coffeeco.Money, cardToken string) error
}

type StoreService interface {
	GetStoreSpecificDiscount(ctx context.Context, storeId uuid.UUID) (float32, error)
}

type Service struct {
	cardService  CardChargeService
	purchaseRepo Repository
	storeService StoreService
}

func NewService(cardService CardChargeService, purchaseRepo Repository, storeService StoreService) *Service {
	return &Service{
		cardService:  cardService,
		purchaseRepo: purchaseRepo,
		storeService: storeService,
	}
}

func (s Service) CompletePurchase(ctx context.Context, storeId uuid.UUID, purchase *Purchase, coffeeBuxCard *membership.CoffeeBux) error {
	err := purchase.validateAndEnrich()
	if err != nil {
		return err
	}

	err = s.calculateStoreSpecificDiscount(ctx, storeId, purchase)
	if err != nil {
		return err
	}

	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		err := s.cardService.ChargeCard(ctx, purchase.total, *purchase.CardToken)
		if err != nil {
			return errors.New("card charge is failed")
		}

	case payment.MEANS_CASH:
	// do nothing

	case payment.MEANS_COFFEEBUX:
		err := coffeeBuxCard.Pay(ctx, purchase.ProductsToPurchase)
		if err != nil {
			return fmt.Errorf("failed to charge membership card: %w", err)
		}

	default:
		return errors.New("unknown payment type")
	}

	err = s.purchaseRepo.Store(ctx, *purchase)
	if err != nil {
		return errors.New("failed to store purchase")
	}

	if coffeeBuxCard != nil {
		coffeeBuxCard.AddStamp()
	}

	return nil
}

func (s *Service) calculateStoreSpecificDiscount(ctx context.Context, storeId uuid.UUID, purchase *Purchase) error {
	discount, err := s.storeService.GetStoreSpecificDiscount(ctx, storeId)
	if err != nil && !errors.Is(err, store.ErrNoDiscount) {
		return fmt.Errorf("failed to get discount: %w", err)
	}

	purchasePrice := purchase.total
	if discount > 0 {
		purchase.total = purchasePrice * coffeeco.Money(100-discount)
	}

	return nil
}
