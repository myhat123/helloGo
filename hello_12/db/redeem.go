package db

import (
	"sort"
	"time"

	"gopkg.in/inf.v0"
)

type Redeem struct {
	OrganCode              string
	TradeDate              time.Time
	TradeTime              string
	ApplicationCode        string
	TaCode                 string
	ProdCode               string
	TradeQuot              *inf.Dec
	TradeCode              string
	RegeemSequence         int64
	AppiontApplicationCode string
}

type RedeemRecords struct {
	redeems []Redeem
}

// 排序必备
func (r RedeemRecords) Sort() {
	sort.Sort(Redeems(r.redeems))
}

type Redeems []Redeem

func (red Redeems) Len() int {
	return len(red)
}

func (red Redeems) Less(i, j int) bool {
	var t1 = red[i].TradeDate.Format("20060102") + red[i].TradeTime
	var t2 = red[j].TradeDate.Format("20060102") + red[j].TradeTime
	return t1 < t2
}

func (red Redeems) Swap(i, j int) {
	red[i], red[j] = red[j], red[i]
}

// 添加元素
func (red *RedeemRecords) Append(r Redeem) {
	red.redeems = append(red.redeems, r)
}
