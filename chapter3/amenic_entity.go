package chapter3

import "time"

type AnemicAuction struct {
	id            int
	startingPrice Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

func (a *AnemicAuction) GetID() int {
	return a.id
}

func (a *AnemicAuction) StartingPrice() Money {
	return a.startingPrice
}

func (a *AnemicAuction) SetStartingPrice(startingPrice Money) {
	a.startingPrice = startingPrice
}

func (a *AnemicAuction) GetSellerID() int {
	return a.sellerID
}

func (a *AnemicAuction) SetSellerID(sellerID int) {
	a.sellerID = sellerID
}

func (a *AnemicAuction) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *AnemicAuction) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

func (a *AnemicAuction) GetAuctionStart() time.Time {
	return a.auctionStart
}

func (a *AnemicAuction) SetAuctionStart(auctionStart time.Time) {
	a.auctionStart = auctionStart
}

func (a *AnemicAuction) GetAuctionEnd() time.Time {
	return a.auctionEnd
}

func (a *AnemicAuction) SetAuctionEnd(auctionEnd time.Time) {
	a.auctionEnd = auctionEnd
}
