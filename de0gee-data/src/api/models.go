package api

// BinaryStats is a structure that derives the following metrics https://en.wikipedia.org/wiki/Sensitivity_and_specificity
type BinaryStats struct {
	TruePositives  int
	FalsePositives int
	TrueNegatives  int
	FalseNegatives int

	// Sensitivity or true positive rate
	Sensitivity float64
	// Specificity or true negative rate
	Specificity float64
	// Informedness (TPR + SPC - 1)
	Informedness float64
}

// NewBinaryStats returns a binary stats object
func NewBinaryStats(tp, fp, tn, fn int) BinaryStats {
	tpf := float64(tp)
	fpf := float64(fp)
	tnf := float64(tn)
	fnf := float64(fn)
	return BinaryStats{
		TruePositives:  tp,
		FalsePositives: fp,
		TrueNegatives:  tn,
		FalseNegatives: fn,

		Sensitivity:  tpf / (tpf + fnf),
		Specificity:  tnf / (tnf + fpf),
		Informedness: tpf/(tpf+fnf) + tnf/(tnf+fpf) - 1,
	}
}
