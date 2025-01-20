package jetton

import (
	"math"
	"strconv"

	"github.com/xssnick/tonutils-go/tlb"
)

func UseDecimals(value float64, isTon bool, dec int) tlb.Coins {
	if isTon {
		return tlb.MustFromTON(strconv.FormatFloat(value, 'f', 9, 64))
	}

	exp := float64(dec)
	withDecimals := value * math.Pow(10, exp)
	return tlb.FromNanoTONU(uint64(withDecimals))
}
