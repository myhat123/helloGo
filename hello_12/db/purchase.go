package db

import (
	"fmt"
	"sort"
	"time"

	"gopkg.in/inf.v0"

	log "github.com/sirupsen/logrus"
)

type Purchase struct {
	OrganCode       string
	TradeDate       time.Time
	TradeTime       string
	ApplicationCode string
	TaCode          string
	ProdCode        string
	TradeQuot       *inf.Dec
	TradeCode       string
}

type PurchaseRecords struct {
	purchases []Purchase
}

// 排序必备
func (p PurchaseRecords) Sort() {
	sort.Sort(Purchases(p.purchases))
}

type Purchases []Purchase

func (pur Purchases) Len() int { return len(pur) }

func (pur Purchases) Less(i, j int) bool {
	var t1 = pur[i].TradeDate.Format("20060102") + pur[i].TradeTime
	var t2 = pur[j].TradeDate.Format("20060102") + pur[j].TradeTime
	return t1 < t2
}

func (pur Purchases) Swap(i, j int) {
	pur[i], pur[j] = pur[j], pur[i]
}

// 打印元素
func (p *PurchaseRecords) Print() {
	for _, x := range p.purchases {
		fmt.Println(x.OrganCode, x.TradeDate, x.TradeTime)
	}
}

// 添加元素
func (p *PurchaseRecords) Append(d Purchase) {
	p.purchases = append(p.purchases, d)
}

// 清理为0的记录
func (p *PurchaseRecords) Clean() {

	var nps []Purchase

	for _, x := range p.purchases {
		if x.TradeQuot.Cmp(inf.NewDec(0, 0)) != 0 {
			// fmt.Println(x.TradeQuot)
			nps = append(nps, x)
		}
	}

	p.purchases = nps
}

// 根据赎回记录，定位购买记录日期位置

func (p PurchaseRecords) Index(r Redeem) int {
	pos := -1
	if len(p.purchases) == 0 {
		log.WithFields(log.Fields{
			"TaCode":                 r.TaCode,
			"ProdCode":               r.ProdCode,
			"AppiontApplicationCode": r.AppiontApplicationCode,
		}).Info("定位购买记录-对应的赎回找不到购买记录")
		return pos
	}
	var end string

	// 赎回记录日期时间
	var rd = r.TradeDate.Format("20060102") + r.TradeTime

	// 赎回日期小于第一个购买记录日期
	
	var start = p.purchases[0].TradeDate.Format("20060102") + p.purchases[0].TradeTime
	
	if rd < start {
		return pos
	}
	
	var found = false
	for i, x := range p.purchases {
		end = x.TradeDate.Format("20060102") + x.TradeTime
		if (start <= rd) && (rd <= end) {
			pos = i - 1
			found = true
			break
		}
	}

	// 赎回日期大于所有购买记录
	if !found {
		pos = len(p.purchases) - 1
	}
	//fmt.Println("pos1=====",pos)
	return pos
}

// 根据赎回记录，查找合同号记录
func (p PurchaseRecords) FindApplicationCode(r Redeem) int {
	pos := -1

	for i, x := range p.purchases {
		if x.ApplicationCode == r.AppiontApplicationCode {
			pos = i
		}
	}

	return pos
}

// 扣减购买记录购买量
func (p PurchaseRecords) Reduce(r Redeem, prodType string) (err int) {
	//fmt.Println(r)
	
	if prodType == "0001" || prodType == "0003" {
		x := r.RegeemSequence
		//fmt.Println("RegeemSequence=====",x)
		switch x {
		case 1:
			return p.Reduce1(r)
		case 2:
			return p.Reduce2(r)
		case 3:
			return p.Reduce3(r)
		case 4:
			return p.Reduce4(r)
		}
	}

	if prodType == "0002" {
		return p.ReduceClosureProd(r)
	}

	return 1
}

// 开放式产品和定开产品的赎回
// 分4种：1)指定赎回，2)先进先出，3)后进先出，4)全部赎回

// 指定赎回
func (p *PurchaseRecords) Reduce1(r Redeem) (err int) {
	var i = p.FindApplicationCode(r)

	if i == -1 {
		log.WithFields(log.Fields{
			"TaCode":                 r.TaCode,
			"ProdCode":               r.ProdCode,
			"AppiontApplicationCode": r.AppiontApplicationCode,
		}).Info("指定赎回-找不到购买合同号")
		return -1
	}

	z := new(inf.Dec)
	p.purchases[i].TradeQuot = z.Sub(p.purchases[i].TradeQuot, r.TradeQuot)

	//如果最后购买 < 赎回, 将数据写入日志
	if p.purchases[i].TradeQuot.Cmp(inf.NewDec(0, 0)) == -1 {
		log.WithFields(log.Fields{
			"TaCode":    r.TaCode,
			"ProdCode":  r.ProdCode,
			"TradeQuot": r.TradeQuot,
		}).Info("指定赎回-购买扣减赎回出现负数")
		return -1
	}

	return 1
}

// 先进先出
func (p *PurchaseRecords) Reduce2(r Redeem) (err int) {
	//fmt.Println("r3=====",r)
	
	var pos = p.Index(r)
	//fmt.Println("pos3=====",pos)
	var t = r.TradeQuot

	z1 := new(inf.Dec)
	z2 := new(inf.Dec)

	for i := 0; i <= pos; i++ {
		if t.Cmp(inf.NewDec(0, 0)) <= 0 {
			break
		}

		if p.purchases[i].TradeQuot.Cmp(t) <= 0 {
			t = z1.Sub(t, p.purchases[i].TradeQuot)
			p.purchases[i].TradeQuot = inf.NewDec(0, 0)
		} else {
			p.purchases[i].TradeQuot = z2.Sub(p.purchases[i].TradeQuot, t)
			t = inf.NewDec(0, 0)
		}
	}
	//如果最后购买 < 赎回, 将数据写入日志
	if t.Cmp(inf.NewDec(0, 0)) == 1 {
		log.WithFields(log.Fields{
			"TaCode":    r.TaCode,
			"ProdCode":  r.ProdCode,
			"TradeQuot": r.TradeQuot,
		}).Info("先进先出-购买扣减赎回出现负数")

		return -1
	}

	return 1
}

// 后进先出
func (p *PurchaseRecords) Reduce3(r Redeem) (err int) {
	//fmt.Println("r=====",r)
	
	var pos = p.Index(r)
	//fmt.Println("pos2=====",pos)
	var t = r.TradeQuot

	z1 := new(inf.Dec)
	z2 := new(inf.Dec)

	for i := pos; i >= 0; i-- {

		if t.Cmp(inf.NewDec(0, 0)) <= 0 {
			break
		}

		// log.WithFields(log.Fields{
		// 	"TaCode":       r.TaCode,
		// 	"ProdCode":     r.ProdCode,
		// 	"PurchaseQuot": p.purchases[i].TradeQuot,
		// 	"PurchaseDate": p.purchases[i].TradeDate,
		// 	"RedeemQuot":   t,
		// 	"RedeemDate":   r.TradeDate,
		// }).Info("后进先出-扣减之前")

		if p.purchases[i].TradeQuot.Cmp(t) <= 0 {
			t = z1.Sub(t, p.purchases[i].TradeQuot)
			p.purchases[i].TradeQuot = inf.NewDec(0, 0)

		} else {
			p.purchases[i].TradeQuot = z2.Sub(p.purchases[i].TradeQuot, t)
			t = inf.NewDec(0, 0)
		}

		// log.WithFields(log.Fields{
		// 	"TaCode":       r.TaCode,
		// 	"ProdCode":     r.ProdCode,
		// 	"PurchaseQuot": p.purchases[i].TradeQuot,
		// 	"RedeemQuot":   t,
		// }).Info("后进先出-扣减之后")

	}
	//如果最后购买 < 赎回, 将数据写入日志
	if t.Cmp(inf.NewDec(0, 0)) == 1 {
		log.WithFields(log.Fields{
			"TaCode":    r.TaCode,
			"ProdCode":  r.ProdCode,
			"TradeQuot": r.TradeQuot,
		}).Info("后进先出-购买扣减赎回出现负数")

		return -1
	}

	return 1
}

// 全部赎回 算法和Reduce2一致
func (p *PurchaseRecords) Reduce4(r Redeem) (err int) {
	var pos = p.Index(r)

	var t = r.TradeQuot

	z1 := new(inf.Dec)
	z2 := new(inf.Dec)

	for i := 0; i <= pos; i++ {
		if t.Cmp(inf.NewDec(0, 0)) <= 0 {
			break
		}

		if p.purchases[i].TradeQuot.Cmp(t) <= 0 {
			t = z1.Sub(t, p.purchases[i].TradeQuot)
			p.purchases[i].TradeQuot = inf.NewDec(0, 0)
		} else {
			p.purchases[i].TradeQuot = z2.Sub(p.purchases[i].TradeQuot, t)
			t = inf.NewDec(0, 0)
		}
	}
	//如果最后购买 < 赎回, 将数据写入日志
	if t.Cmp(inf.NewDec(0, 0)) == 1 {
		log.WithFields(log.Fields{
			"TaCode":    r.TaCode,
			"ProdCode":  r.ProdCode,
			"TradeQuot": r.TradeQuot,
		}).Info("全部赎回-购买扣减赎回出现负数")

		return -1
	}

	return 1
}

// 封闭式产品赎回, 赎回序号为空
func (p *PurchaseRecords) ReduceClosureProd(r Redeem) (err int) {
	var pos = p.Index(r)

	// log.WithFields(log.Fields{
	// 	"TaCode":    r.TaCode,
	// 	"ProdCode":  r.ProdCode,
	// 	"TradeQuot": r.TradeQuot,
	// 	"TradeDate": r.TradeDate.Format("20060102"),
	// }).Info("封闭式产品-赎回信息")

	var t = r.TradeQuot

	z1 := new(inf.Dec)

	for i := 0; i <= pos; i++ {
		if t.Cmp(inf.NewDec(0, 0)) <= 0 {
			break
		}

		if p.purchases[i].ProdCode == r.ProdCode {
			t = z1.Sub(t, p.purchases[i].TradeQuot)
			p.purchases[i].TradeQuot = inf.NewDec(0, 0)
		}
	}

	if t.Cmp(inf.NewDec(0, 0)) == 1 {
		log.WithFields(log.Fields{
			"TaCode":    r.TaCode,
			"ProdCode":  r.ProdCode,
			"TradeQuot": r.TradeQuot,
		}).Info("封闭式产品-购买扣减赎回出现负数")
		return -1
	}

	return 1
}
