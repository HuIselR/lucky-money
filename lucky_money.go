package lm

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand/v2"
)

var (
	// MaxAmount 红包的最大金额，可自行定义
	MaxAmount = decimal.NewFromFloat(200)
	// MinAmount 红包的最小金额，可自行定义
	MinAmount = decimal.NewFromFloat(1)
	// Threshold 红包阈值
	Threshold = MaxAmount.Sub(MinAmount)
)

const (
	// Random 随机金额红包
	Random = iota
	// Identical 普通红包（每个红包金额相等）
	Identical
)

type LuckyMoney struct {
	// 数量
	Quantity int32
	// 总金额
	Amount decimal.Decimal
}

type Checker interface {
	Check() error
}

func (l *LuckyMoney) Check() error {
	if l.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	if l.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than 'decimal.Zero'")
	}
	return nil
}

func (l *LuckyMoney) String() string {
	return fmt.Sprintf("{Quantity:%d, Amount:%s}", l.Quantity, l.Amount)
}

// Split 拆分算法
// st: 可以选择 Random 和 Identical
func Split(l LuckyMoney, st int) ([]decimal.Decimal, error) {
	if err := Check(l); err != nil {
		return nil, err
	}
	switch st {
	case Random:
		return random(l)
	case Identical:
		return identical(l)
	default:
		return nil, errors.New("invalid st")
	}
}

func SplitRandom(l LuckyMoney) ([]decimal.Decimal, error) {
	return Split(l, Random)
}

func SplitIdentical(l LuckyMoney) ([]decimal.Decimal, error) {
	return Split(l, Identical)
}

func random(l LuckyMoney) ([]decimal.Decimal, error) {
	qd := decimal.NewFromInt32(l.Quantity)
	if l.Amount.Equal(MaxAmount.Mul(qd)) {
		return identical(l)
	}
	// 每个红包预留最小金额 #{MinAmount}
	sm := l.Amount.Sub(MinAmount.Mul(qd))
	if sm.Equal(decimal.Zero) {
		return identical(l)
	}
	var rps = make([]decimal.Decimal, 0)
	// 二倍拆分法，最后一个红包用减法，因为四舍五入存在金额丢失的问题
	for l.Quantity > 1 && sm.GreaterThan(decimal.Zero) {
		avg := sm.Div(qd)
		n := decimal.NewFromFloat(rand.Float64()).Mul(avg).Mul(decimal.NewFromInt(2)).Truncate(2)
		// 拆分出来的红包有可能超过最大阈值，超过则处理为最大金额
		if n.GreaterThan(Threshold) {
			n = Threshold
		}
		rps = append(rps, n.Add(MinAmount))
		sm = sm.Sub(n)
		l.Quantity--
	}
	// 判断最后一个是否超过最大金额
	// 超过则随机分配到其他红包种
	lst := sm.Add(MinAmount)
	sub := lst.Sub(MaxAmount)
	if sub.GreaterThan(decimal.Zero) {
		lst = MaxAmount
		// 将超出的金额随机分配到其他红包中
		for {
			i := rand.IntN(len(rps))
			rp := rps[i].Add(sub)
			if rp.LessThanOrEqual(MaxAmount) {
				rps[i] = rp
				break
			}
			rps[i] = MaxAmount
			sub = rp.Sub(MaxAmount)
		}
	}
	rps = append(rps, lst)
	return rps, nil
}

func identical(l LuckyMoney) ([]decimal.Decimal, error) {
	q := l.Quantity
	rps := make([]decimal.Decimal, 0)
	avg := l.Amount.Div(decimal.NewFromInt32(q)).Round(2)
	for ; q > 1; q-- {
		rps = append(rps, avg)
	}
	return append(rps, l.Amount.Sub(avg.Mul(decimal.NewFromInt32(l.Quantity-1)))), nil
}

func Check(l LuckyMoney) error {
	if err := l.Check(); err != nil {
		return err
	}
	qd := decimal.NewFromInt32(l.Quantity)
	if l.Amount.Sub(qd.Mul(MinAmount)).LessThan(decimal.Zero) {
		return errors.New("insufficient amount, should be greater than 'MinAmount'")
	}
	if l.Amount.Sub(MaxAmount.Mul(qd)).GreaterThan(decimal.Zero) {
		return errors.New("overflow amount, should be less than 'MaxAmount'")
	}
	return nil
}
