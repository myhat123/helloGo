package db

import (
	// "fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"gopkg.in/inf.v0"
)

//d 是购买
var d1 = Purchase{"36015828", time.Date(2017, time.October, 20, 0, 0, 0, 0, time.UTC), "132112", "171020710975085783", "710910010742065", "1300099C89", inf.NewDec(5000000, 2), "603302"}
var d2 = Purchase{"36015828", time.Date(2018, time.February, 22, 0, 0, 0, 0, time.UTC), "120642", "180222710986153573", "710910010742065", "1300099C89", inf.NewDec(7000000, 2), "603302"}
var d3 = Purchase{"36015828", time.Date(2018, time.February, 28, 0, 0, 0, 0, time.UTC), "125759", "180228710987050869", "710910010742065", "1300099C89", inf.NewDec(5000000, 2), "603302"}

//red 是赎回
var red1 = Redeem{"36015828", time.Date(2018, time.February, 19, 0, 0, 0, 0, time.UTC), "141759", "180225710987050869", "710910010742065", "1300099C89", inf.NewDec(5000000, 2), "603302", 1, "171020710975085783"}
var red4 = Redeem{"36015828", time.Date(2018, time.February, 27, 0, 0, 0, 0, time.UTC), "141759", "180225710987050869", "710910010742065", "1300099C89", inf.NewDec(12000000, 2), "603302", 1, "171020710975085783"}
var red5 = Redeem{"36015828", time.Date(2018, time.February, 27, 0, 0, 0, 0, time.UTC), "141759", "180225710987050869", "710910010742065", "1300099C89", inf.NewDec(25000000, 2), "603302", 1, "171020710975085783"}
var red6 = Redeem{"36015828", time.Date(2018, time.February, 27, 0, 0, 0, 0, time.UTC), "141759", "180225710987050869", "710910010742065", "1300099C89", inf.NewDec(7000000, 2), "603302", 1, "171020710975085783"}

//d 是购买
var d4 = Purchase{"36015828", time.Date(2017, time.October, 20, 0, 0, 0, 0, time.UTC), "132112", "171020710975085783", "710910010742065", "1300099C89", inf.NewDec(5000000, 2), "603302"}
var d5 = Purchase{"36015828", time.Date(2018, time.February, 22, 0, 0, 0, 0, time.UTC), "120642", "180222710986153573", "710910010742065", "1300099C89", inf.NewDec(7000000, 2), "603302"}
var d6 = Purchase{"36015828", time.Date(2018, time.February, 28, 0, 0, 0, 0, time.UTC), "125759", "180228710987050869", "710910010742065", "1300099C89", inf.NewDec(5000000, 2), "603302"}

//red 是赎回
var red2 = Redeem{"36015828", time.Date(2018, time.February, 19, 0, 0, 0, 0, time.UTC), "141759", "180225710987050869", "710910010742065", "1300099C89", inf.NewDec(2000000, 2), "603302", 1, "171020710975085783"}

//red 是赎回
var red3 = Redeem{"36015828", time.Date(2018, time.February, 19, 0, 0, 0, 0, time.UTC), "141759", "180225710987050869", "710910010742065", "1300099C89", inf.NewDec(20000000, 2), "603302", 1, "171020710975085783"}

// 测试1
func TestPurchase1(t *testing.T) {
	r := PurchaseRecords{
		purchases: []Purchase{d3, d2, d1},
	}

	x := r.purchases[0]
	assert.Equal(t, len(r.purchases), 3, "3个元素")
	assert.Equal(t, x.ApplicationCode, "180228710987050869", "第一个元素")

	r.Sort()

	x = r.purchases[0]
	assert.Equal(t, x.TradeDate.Format("20060102"), "20171020", "日期时间")
	assert.Equal(t, x.ApplicationCode, "171020710975085783", "第一个元素")
}

// 测试2
func TestPurchase2(t *testing.T) {
	r := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r.purchases = append(r.purchases, d3, d2, d1)

	r.Sort()

	x := r.purchases[0]
	assert.Equal(t, x.TradeDate.Format("20060102"), "20171020", "日期时间")
	assert.Equal(t, x.ApplicationCode, "171020710975085783", "第一个元素")
}

// 测试清理
func TestPurchaseClean(t *testing.T) {
	r := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r.purchases = append(r.purchases, d3, d2, d1)

	r.purchases[1].TradeQuot = inf.NewDec(0, 0)

	r.Clean()

	assert.Equal(t, len(r.purchases), 2, "2个元素")

	x := &(r.purchases[1])
	x.TradeQuot = inf.NewDec(0, 0)

	r.Clean()
	assert.Equal(t, len(r.purchases), 1, "1个元素")
}

// 测试定位
func TestPurchaseIndex(t *testing.T) {
	r := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r.purchases = append(r.purchases, d3, d2, d1)

	r.Sort()

}

// 测试查找指定合同号
func TestPurchaseFindApplicationCode(t *testing.T) {

	r := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r.purchases = append(r.purchases, d3, d2, d1)

	r.Sort()

	assert.Equal(t, r.FindApplicationCode(red1), 0, "查找指定合同号")
}

// 测试指定赎回
func TestPurchaseReduce1(t *testing.T) {

	r := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r.purchases = append(r.purchases, d3, d2, d1)

	r.Sort()

	r.Reduce1(red1)
	r.Clean()

	assert.Equal(t, len(r.purchases), 2, "指定赎回，购买等于赎回")

	r1 := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r1.purchases = append(r1.purchases, d4, d5, d6)

	r1.Sort()

	r1.Reduce1(red2)
	r1.Clean()

	assert.Equal(t, len(r1.purchases), 3, "指定赎回，购买大于赎回")

	r2 := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r2.purchases = append(r2.purchases, d4, d5, d6)

	r2.Sort()

	r2.Reduce1(red3)
	r2.Clean()

	assert.Equal(t, 3, len(r2.purchases), "指定赎回，购买小于赎回")

}

// 测试先进先出赎回
func TestPurchaseReduce2(t *testing.T) {

	r := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r.purchases = append(r.purchases, d3, d2, d1)

	r.Sort()

	r.Reduce2(red1)
	r.Clean()

	assert.Equal(t, len(r.purchases), 2, "先进先出赎回")

	r1 := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r1.purchases = append(r1.purchases, d3, d2, d1)

	r1.Sort()

	r1.Reduce2(red4)
	r1.Clean()

	assert.Equal(t, 1, len(r1.purchases), "先进先出赎回")

	r2 := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r2.purchases = append(r2.purchases, d3, d2, d1)

	r2.Sort()

	r2.Reduce2(red5)
	r2.Clean()

	assert.Equal(t, len(r2.purchases), 1, "先进先出赎回，购买小于赎回")
}

// 测试后进先出赎回
func TestPurchaseReduce3(t *testing.T) {

	r := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r.purchases = append(r.purchases, d3, d2, d1)

	r.Sort()

	r.Reduce3(red6)
	r.Clean()

	assert.Equal(t, len(r.purchases), 2, "后进先出")

	r1 := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r1.purchases = append(r1.purchases, d3, d2, d1)

	r1.Sort()

	r1.Reduce3(red4)
	r1.Clean()

	assert.Equal(t, 1, len(r1.purchases), "后进先出")

	r2 := new(PurchaseRecords)
	// r.purchases = make([]Purchase, 0)
	r2.purchases = append(r2.purchases, d3, d2, d1)

	r2.Sort()

	r2.Reduce3(red5)
	r2.Clean()

	assert.Equal(t, len(r2.purchases), 1, "后进先出，购买小于赎回")
}
