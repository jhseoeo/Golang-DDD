package chapter8

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -package mocks -destination chapter8/mocks/cookies.go github.com/jhseoeo/Golang-DDD/chapter8 CookieStockChecker,CardCharger,EmailSender
