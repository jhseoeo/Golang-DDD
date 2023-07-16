package chapter3

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type AuctionRefactored struct {
	id            uuid.UUID
	startingPrice Money
	sellerID      uuid.UUID
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

func (a *AuctionRefactored) GetAuctionElapsedDuration() time.Duration {
	return a.auctionStart.Sub(a.auctionEnd)
}

func (a *AuctionRefactored) GetAuctionEndTimeInUTC() time.Time {
	return a.auctionEnd
}

func (a *AuctionRefactored) SetAuctionEnd(auctionEnd time.Time) error {
	err := a.validateTimeZone(auctionEnd)
	if err != nil {
		return err
	}
	a.auctionEnd = auctionEnd
	return nil
}

func (a *AuctionRefactored) GetAuctionStartTimeInUTC() time.Time {
	return a.auctionStart
}

func (a *AuctionRefactored) SetAuctionStartTimeInUTC(auctionStart time.Time) error {
	err := a.validateTimeZone(auctionStart)
	if err != nil {
		return err
	}
	a.auctionStart = auctionStart
	return nil
}

func (a *AuctionRefactored) GetId() uuid.UUID {
	return a.id
}

func (a *AuctionRefactored) validateTimeZone(t time.Time) error {
	tz, _ := t.Zone()
	if tz != time.UTC.String() {
		return errors.New("time zone must be UTC")
	} else {
		return nil
	}
}
