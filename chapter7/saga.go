package chapter7

import "context"

type Saga interface {
	Execute(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type OrderCreator struct{}

func (o *OrderCreator) Execute(ctx context.Context) error {
	return o.createOrder(ctx)
}

func (o *OrderCreator) Rollback(ctx context.Context) error {
	return o.deleteOrder(ctx)
}

func (o *OrderCreator) createOrder(ctx context.Context) error {
	// create order
	return nil
}

func (o *OrderCreator) deleteOrder(ctx context.Context) error {
	// delete order
	return nil
}

type PaymentCreator struct{}

func (p *PaymentCreator) Execute(ctx context.Context) error {
	return p.createPayment(ctx)
}

func (p *PaymentCreator) Rollback(ctx context.Context) error {
	return p.deletePayment(ctx)
}

func (p *PaymentCreator) createPayment(ctx context.Context) error {
	// create payment
	return nil
}

func (p *PaymentCreator) deletePayment(ctx context.Context) error {
	// delete payment
	return nil
}

type SagaManager struct {
	actions []Saga
}

func (s *SagaManager) Execute(ctx context.Context) {
	for i, action := range s.actions {
		err := action.Execute(ctx)
		if err != nil {
			for j := 0; j < i; j++ {
				err := s.actions[j].Rollback(ctx)
				if err != nil {
					// One of our compensation actions failed;
					// by emitting a message to a a messagebus, we need to handle it
				}
			}
		}
	}
}
