package recommendation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PartnershipAdapter struct {
	client *http.Client
	url    string
}

func NewPartnershipAdapter(client *http.Client, url string) (*PartnershipAdapter, error) {
	if client == nil {
		return nil, errors.New("client is required")
	}
	if url == "" {
		return nil, errors.New("url is required")
	}
	return &PartnershipAdapter{
		client: client,
		url:    url,
	}, nil
}

type partnershipResponse struct {
	AvailableHotels []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	} `json:"availableHotels"`
}

func (p PartnershipAdapter) GetAvailability(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string) ([]Option, error) {
	from := fmt.Sprintf("%d-%d-%d", tripEnd.Year(), tripEnd.Month(), tripEnd.Day())
	to := fmt.Sprintf("%d-%d-%d", tripStart.Year(), tripStart.Month(), tripStart.Day())
	url := fmt.Sprintf("%s/partnerships?from=%s&to=%s&location=%s", p.url, from, to, location)
	res, err := p.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call partnership service: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad request to partnership service: %d", res.StatusCode)
	}

	var pr partnershipResponse
	err = json.NewDecoder(res.Body).Decode(&pr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode partnership response: %w", err)
	}

	opts := make([]Option, len(pr.AvailableHotels))
	for i, p := range pr.AvailableHotels {
		opts[i] = Option{
			HotelName:     p.Name,
			Location:      location,
			PricePerNight: Money(p.PriceInUSDPerNight),
		}
	}
	return opts, nil
}
