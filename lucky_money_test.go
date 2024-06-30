package lm

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestSuccessLMSRandom(t *testing.T) {
	lml := make([]LuckyMoney, 0)
	lml = append(lml, LuckyMoney{Quantity: 33, Amount: decimal.NewFromFloat(33)})
	lml = append(lml, LuckyMoney{Quantity: 99, Amount: decimal.NewFromFloat(101)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(32)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(1.01)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(11)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(2000)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(1988)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(1999)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(1500)})
	for _, lm := range lml {
		a := lm.Amount
		split, err := SplitRandom(lm)
		if err != nil {
			t.Errorf("拆分红包失败 %s %s", lm.String(), err)
			continue
		}
		sum := decimal.Zero
		for _, d := range split {
			if d.GreaterThan(MaxAmount) || d.LessThan(MinAmount) {
				t.Errorf("红包金额异常 %s", d)
			}
			sum = sum.Add(d)
		}
		if !a.Equal(sum) {
			t.Errorf("红包拆分异常 %s %s != %s", lm.String(), sum, a)
		}
	}
}

func TestErrorLMSRandom(t *testing.T) {
	lml := make([]LuckyMoney, 0)
	lml = append(lml, LuckyMoney{Quantity: -1, Amount: decimal.NewFromFloat(3)})
	lml = append(lml, LuckyMoney{Quantity: 0, Amount: decimal.NewFromFloat(10)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(0)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(-1.01)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(0.89)})
	lml = append(lml, LuckyMoney{Quantity: 3, Amount: decimal.NewFromFloat(2.99)})
	lml = append(lml, LuckyMoney{Quantity: 100, Amount: decimal.NewFromFloat(99.999)})
	lml = append(lml, LuckyMoney{Quantity: 9, Amount: decimal.NewFromFloat(1801)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(2500)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(2001)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(3000)})
	fn := 0
	for _, lm := range lml {
		_, err := SplitRandom(lm)
		if err != nil {
			fn++
		}
	}
	if fn != len(lml) {
		t.Errorf("拆分红包失败: 预期失败用例 %d，实际失败用例 %d", fn, len(lml))
	}
}

func TestSuccessLMSIdentical(t *testing.T) {
	lml := make([]LuckyMoney, 0)
	lml = append(lml, LuckyMoney{Quantity: 999, Amount: decimal.NewFromFloat(1000)})
	lml = append(lml, LuckyMoney{Quantity: 9999, Amount: decimal.NewFromFloat(10000)})
	lml = append(lml, LuckyMoney{Quantity: 333, Amount: decimal.NewFromFloat(30000)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(2000)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(1999.99)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(16)})
	lml = append(lml, LuckyMoney{Quantity: 100, Amount: decimal.NewFromFloat(1600)})
	lml = append(lml, LuckyMoney{Quantity: 1000, Amount: decimal.NewFromFloat(1600)})
	lml = append(lml, LuckyMoney{Quantity: 1000, Amount: decimal.NewFromFloat(1000.01)})
	lml = append(lml, LuckyMoney{Quantity: 1000, Amount: decimal.NewFromFloat(1001.99)})
	lml = append(lml, LuckyMoney{Quantity: 99, Amount: decimal.NewFromFloat(19800)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(1)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(200)})
	lml = append(lml, LuckyMoney{Quantity: 3, Amount: decimal.NewFromFloat(4)})
	for _, lm := range lml {
		a := lm.Amount
		split, err := SplitIdentical(lm)
		if err != nil {
			t.Errorf("拆分红包失败 %s %s", lm.String(), err)
			continue
		}
		sum := decimal.Zero
		for _, d := range split {
			if d.GreaterThan(MaxAmount) || d.LessThan(MinAmount) {
				t.Errorf("红包金额异常 %s %s", lm.String(), d)
			}
			sum = sum.Add(d)
		}
		if !a.Equal(sum) {
			t.Errorf("红包拆分异常 %s != %s", sum, a)
		}
	}
}

func TestErrorLMSIdentical(t *testing.T) {
	lml := make([]LuckyMoney, 0)
	lml = append(lml, LuckyMoney{Quantity: -1, Amount: decimal.NewFromFloat(3)})
	lml = append(lml, LuckyMoney{Quantity: 0, Amount: decimal.NewFromFloat(10)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(0)})
	lml = append(lml, LuckyMoney{Quantity: 1, Amount: decimal.NewFromFloat(-1.01)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(0.89)})
	lml = append(lml, LuckyMoney{Quantity: 3, Amount: decimal.NewFromFloat(2.99)})
	lml = append(lml, LuckyMoney{Quantity: 100, Amount: decimal.NewFromFloat(99.999)})
	lml = append(lml, LuckyMoney{Quantity: 9, Amount: decimal.NewFromFloat(1801)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(2500)})
	lml = append(lml, LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(2001)})
	lml = append(lml, LuckyMoney{Quantity: 15, Amount: decimal.NewFromFloat(3001)})
	fn := 0
	for _, lm := range lml {
		_, err := SplitIdentical(lm)
		if err != nil {
			fn++
		}
	}
	if fn != len(lml) {
		t.Errorf("拆分红包失败: 预期失败用例 %d，实际失败用例 %d", fn, len(lml))
	}
}

func BenchmarkLMSIdentical10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(33)}
		_, _ = SplitIdentical(lm)
	}
}

func BenchmarkLMSIdentical100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 100, Amount: decimal.NewFromFloat(333)}
		_, _ = SplitIdentical(lm)
	}
}

func BenchmarkLMSIdentical1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 1000, Amount: decimal.NewFromFloat(3333)}
		_, _ = SplitIdentical(lm)
	}
}
func BenchmarkLMSIdentical10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 10000, Amount: decimal.NewFromFloat(33333)}
		_, _ = SplitIdentical(lm)
	}
}
func BenchmarkLMSIdentical100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 100000, Amount: decimal.NewFromFloat(333333)}
		_, _ = SplitIdentical(lm)
	}
}

func BenchmarkLMSRandom10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 10, Amount: decimal.NewFromFloat(19)}
		_, _ = SplitRandom(lm)
	}
}

func BenchmarkLMSRandom100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 100, Amount: decimal.NewFromFloat(199)}
		_, _ = SplitRandom(lm)
	}
}

func BenchmarkLMSRandom1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 1000, Amount: decimal.NewFromFloat(1999)}
		_, _ = SplitRandom(lm)
	}
}

func BenchmarkLMSRandom10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 10000, Amount: decimal.NewFromFloat(19999)}
		_, _ = SplitRandom(lm)
	}
}

func BenchmarkLMSRandom100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lm := LuckyMoney{Quantity: 100000, Amount: decimal.NewFromFloat(199999)}
		_, _ = SplitRandom(lm)
	}
}
