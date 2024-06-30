# lucky-money（红包）

## 介绍

此项目包含常见的红包拆分算法，随机，普通红包。提供开箱即用逻辑完整的红包拆分算法，避免手动实现时特殊场景考虑不全的问题，造成金额损失。

## 红包算法

### Random（随机）

二倍均值拆分，每个小红包的金额有范围限制 [MinAmount,MaxAmount]

- `MinAmount` 拆分出来金额的最小值，默认为1，可自定义调整
- `MaxAmount` 拆分出来金额的最大值，默认为200，可自定义调整

保证拆分后的金额在范围限制内。

### Identical（普通）

顾名思义，就是拆分后每个小红包的金额是相等的。

但存在`总金额/数量`除不尽的情况，如总金额为10，数量为3，那么拆分出来的小红包金额为 3.33, 3.33, 3.34。

## 必要条件

go version >= v1.10

## 安装

`go get github.com/huiselr/lm`

## 示例

```go
package main

import (
	"fmt"
	"github.com/huiselr/lm"
	"github.com/shopspring/decimal"
)

func main() {
	// 随机
	rps, err := lm.SplitRandom(lm.LuckyMoney{Quantity: 3, Amount: decimal.NewFromFloat(10)})
	if err != nil {
		panic(err)
	}
	for i, rp := range rps {
		fmt.Println(i, rp)
	}

	// 普通
	rps, err = lm.SplitIdentical(lm.LuckyMoney{Quantity: 3, Amount: decimal.NewFromFloat(10)})
	if err != nil {
		panic(err)
	}
	for i, rp := range rps {
		fmt.Println(i, rp)
	}
}

```

## 规划

- 提供惰性拆分，即每次调用时拆且仅拆一个小红包。