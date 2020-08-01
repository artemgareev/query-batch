package querybatch

type QueryBatches struct {
	Limit  int64
	Offset int64
}

func GetQueryBatches(batchSize int64, recordsNumber int64) []QueryBatches {
	return GetQueryBatchesWithOffset(batchSize, recordsNumber, 0)
}

func GetQueryBatchesWithOffset(batchSize int64, recordsNumber int64, startOffset int64) []QueryBatches {
	if startOffset > recordsNumber {
		return nil
	}

	if batchSize > recordsNumber {
		return []QueryBatches{{recordsNumber, 0}}
	}

	var limitBatches []QueryBatches
	for i := startOffset + batchSize; i < recordsNumber+batchSize; i += batchSize {
		offset := i - batchSize
		limit := batchSize

		if offset+limit >= recordsNumber {
			limit = recordsNumber - offset
		}
		limitBatches = append(limitBatches, QueryBatches{limit, offset})
	}

	return limitBatches
}
