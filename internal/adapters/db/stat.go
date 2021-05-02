package db

const (
	selectStatsSQL = `
SELECT SUM(CASE WHEN p.is_mutant = 1 THEN 1 ELSE 0 END) AS count_mutant_dna, SUM(CASE WHEN p.is_mutant = 0 THEN 1 ELSE 0 END) AS count_human_dna, SUM(CASE WHEN p.is_mutant=1 THEN 1 ELSE 0 END)/COUNT(*) AS ratio FROM  person p

`
)

func (db DB) StatsGetMutantVsHuman() (mutantCount, humanCount int64, ratio float32, err error) {
	row := db.Raw(selectStatsSQL).Row()
	//if the table is empty the result will be nil
	var mutantPointer, humanPointer *int64
	var ratioPointer *float32
	err = row.Scan(&mutantPointer, &humanPointer, &ratioPointer)
	if mutantPointer != nil {
		mutantCount = *mutantPointer
	}
	if humanPointer != nil {
		humanCount = *humanPointer
	}
	if ratioPointer != nil {
		ratio = *ratioPointer
	}

	return
}
