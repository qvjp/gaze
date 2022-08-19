package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
)

const max = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"

var (
	bigIECExp = big.NewInt(1024)

	bigByte   = big.NewInt(1)
	bigKiByte = (&big.Int{}).Mul(bigByte, bigIECExp)
	bigMiByte = (&big.Int{}).Mul(bigKiByte, bigIECExp)
	bigGiByte = (&big.Int{}).Mul(bigMiByte, bigIECExp)
	bigTiByte = (&big.Int{}).Mul(bigGiByte, bigIECExp)
	bigPiByte = (&big.Int{}).Mul(bigTiByte, bigIECExp)
	bigEiByte = (&big.Int{}).Mul(bigPiByte, bigIECExp)
	bigZiByte = (&big.Int{}).Mul(bigEiByte, bigIECExp)
	bigYiByte = (&big.Int{}).Mul(bigZiByte, bigIECExp)
)

var bigIBytesSizeSlice = []*big.Int{
	bigByte, bigKiByte, bigMiByte, bigGiByte, bigTiByte, bigPiByte, bigEiByte, bigZiByte, bigYiByte,
}

func humaneBigBytes(s, base *big.Int, sizes []string) string {
	if s.Cmp(big.NewInt(10)) <= 0 {
		return fmt.Sprintf("%d B", s)
	}

	var ret strings.Builder

	for s.Cmp(big.NewInt(10)) > 0 {
		c := (&big.Int{}).Set(s)
		val, mag := oom(c, base, len(sizes)-1)
		suffix := sizes[mag]
		ret.WriteString(fmt.Sprintf("%.0f %s ", math.Floor(val), suffix))

		s = (&big.Int{}).Sub(s, (&big.Int{}).Mul(big.NewInt(int64(val)), bigIBytesSizeSlice[mag]))
	}

	return strings.TrimSpace(ret.String())
}

func oom(n, b *big.Int, max int) (float64, int) {
	mag := 0
	m := &big.Int{}
	for n.Cmp(b) >= 0 {
		n.DivMod(n, b, m)
		mag++
		if mag == max && max >= 0 {
			break
		}
	}
	return float64(n.Int64()) + (float64(m.Int64()) / float64(b.Int64())), mag
}

func IBytes(num string) string {
	base := 10
	if strings.HasPrefix(num, "0x") || strings.HasPrefix(num, "0X") {
		base = 16
		num = strings.TrimPrefix(strings.TrimPrefix(num, "0x"), "0X")
	}

	count, ok := big.NewInt(0).SetString(num, base)
	if !ok {
		return "invalid num"
	}

	maxNum, _ := big.NewInt(0).SetString(max, 16)
	if count.Cmp(maxNum) > 0 {
		return "too big"
	}

	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	return humaneBigBytes(count, bigIECExp, sizes)
}

func main() {
	for k, v := range os.Args {
		if k == 0 {
			continue
		}

		fmt.Println(IBytes(v))
	}
}
