package application

func (i interactor) StatsGetMutantVsHuman() (mutantCount, humanCount int64, ratio float32, err error) {
	return i.storage.StatsGetMutantVsHuman()
}
