package models

import (
	"math"
	"math/big"
)

// BinaryStats is a structure that derives the following metrics https://en.wikipedia.org/wiki/Sensitivity_and_specificity
type BinaryStats struct {
	TruePositives  int `json:"true_positives"`
	FalsePositives int `json:"false_positives"`
	TrueNegatives  int `json:"true_negatives"`
	FalseNegatives int `json:"false_negatives"`

	// Sensitivity or true positive rate
	Sensitivity float64 `json:"sensitivity"`
	// Specificity or true negative rate
	Specificity float64 `json:"specificity"`
	// Informedness (TPR + SPC - 1)
	Informedness float64 `json:"informedness"`
	// Martthews Correlation coefficient
	MCC float64 `json:"mcc"`
	// Fisher's P test
	FisherP float64 `json:"fisher_p"`
}

// NewBinaryStats returns a binary stats object
func NewBinaryStats(tp, fp, tn, fn int) BinaryStats {
	tpf := float64(tp)
	fpf := float64(fp)
	tnf := float64(tn)
	fnf := float64(fn)
	sensitivity := float64(0)
	if tpf+fnf != 0 {
		sensitivity = tpf / (tpf + fnf)
	}
	specificity := float64(0)
	if tnf+fpf != 0 {
		specificity = tnf / (tnf + fpf)
	}
	mcc := float64(0)
	if math.Sqrt((tpf+fpf)*(tpf+fnf)*(tnf+fpf)*(tnf+fnf)) > 0 {
		mcc = (tpf*tnf - fpf*fnf) / math.Sqrt((tpf+fpf)*(tpf+fnf)*(tnf+fpf)*(tnf+fnf))
	}
	fisher_p := float64(1)
	if NChooseK(tpf+fpf+tnf+fnf, tpf+fpf) > 0 {
		fisher_p = NChooseK(tpf+fnf, tpf) * NChooseK(fpf+tnf, fpf) / NChooseK(tpf+fpf+tnf+fnf, tpf+fpf)
	}

	return BinaryStats{
		TruePositives:  tp,
		FalsePositives: fp,
		TrueNegatives:  tn,
		FalseNegatives: fn,

		Sensitivity:  sensitivity,
		Specificity:  specificity,
		Informedness: specificity + sensitivity - 1,
		MCC:          mcc,
		FisherP:      fisher_p,
	}
}

func NChooseK(n float64, k float64) float64 {
	a := big.NewInt(0)
	a.Binomial(int64(n), int64(k))
	f, _ := (new(big.Float).SetInt(a).Float64())
	return f
}
