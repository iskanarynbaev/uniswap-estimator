package uniswap

import (
	"math/big"
)

func CalculateAmountOut(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	feeNumerator := big.NewInt(997)
	feeDenominator := big.NewInt(1000)

	amountInWithFee := new(big.Int).Mul(amountIn, feeNumerator)
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, feeDenominator), amountInWithFee)

	amountOut := new(big.Int).Div(numerator, denominator)
	return amountOut
}
