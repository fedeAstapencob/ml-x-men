package db

const (
	selectStatsSQL = `
SELECT SUM(CASE WHEN p.is_mutant = 1 THEN 1 ELSE 0 END) AS count_mutant_dna, SUM(CASE WHEN p.is_mutant = 0 THEN 1 ELSE 0 END) AS count_human_dna, SUM(CASE WHEN p.is_mutant=1 THEN 1 ELSE 0 END)/COUNT(*) AS ratio FROM  person p

`
)

func (db DB) StatsGetMutantVsHuman() (mutantCount, humanCount int64, ratio float32, err error) {
	row := db.Raw(selectStatsSQL).Row()
	err = row.Scan(&mutantCount, &humanCount, &ratio)
	return
}
