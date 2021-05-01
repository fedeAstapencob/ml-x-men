package json_formatter

type StatResp struct {
	MutantCount int64   `json:"count_mutant_dna"`
	HumanCount  int64   `json:"count_human_dna"`
	Ratio       float32 `json:"ratio"`
}

func NewStatResp(mutantCount, humanCount int64, ratio float32) StatResp {
	return StatResp{
		MutantCount: mutantCount,
		HumanCount:  humanCount,
		Ratio:       ratio,
	}
}
