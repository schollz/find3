package models

import "math"

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
	return BinaryStats{
		TruePositives:  tp,
		FalsePositives: fp,
		TrueNegatives:  tn,
		FalseNegatives: fn,

		Sensitivity:  sensitivity,
		Specificity:  specificity,
		Informedness: specificity + sensitivity - 1,
		MCC:          (tpf*tnf - fpf*fnf) / math.Sqrt((tpf+fpf)*(tpf+fnf)*(tnf+fpf)*(tnf+fnf)),
	}
}
