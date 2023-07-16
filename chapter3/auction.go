package chapter3

import "time"

type Money = int

type Auction struct {
	ID            int
	startingPrice Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}
