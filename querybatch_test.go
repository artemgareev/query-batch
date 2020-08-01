package querybatch

import (
	"reflect"
	"testing"
)

//nolint
func TestGetLimitsBatches(t *testing.T) {
	t.Parallel()

	type args struct {
		batchSize     int64
		recordsNumber int64
	}
	tests := []struct {
		name string
		args args
		want []QueryBatches
	}{
		{
			"regular batching",
			args{500, 1540},
			[]QueryBatches{
				{500, 0},
				{500, 500},
				{500, 1000},
				{40, 1500},
			},
		},
		{
			"batch greater than recordsNumber",
			args{1000, 100},
			[]QueryBatches{
				{100, 0},
			},
		},
		{
			"small batches",
			args{300, 1234},
			[]QueryBatches{
				{300, 0},
				{300, 300},
				{300, 600},
				{300, 900},
				{34, 1200},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQueryBatches(tt.args.batchSize, tt.args.recordsNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueryBatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint
func TestGetLimitsBatchesWithOffset(t *testing.T) {
	t.Parallel()

	type args struct {
		batchSize     int64
		recordsNumber int64
		startOffset   int64
	}
	tests := []struct {
		name string
		args args
		want []QueryBatches
	}{
		{
			"regular batching with round numbers",
			args{100, 1500, 1300},
			[]QueryBatches{
				{100, 1300},
				{100, 1400},
			},
		},
		{
			"regular batching with not rounded numbers",
			args{12, 40, 10},
			[]QueryBatches{
				{12, 10},
				{12, 22},
				{6, 34},
			},
		},
		{
			"offset bigger than  records number",
			args{10, 20, 30},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQueryBatchesWithOffset(tt.args.batchSize, tt.args.recordsNumber, tt.args.startOffset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueryBatchesWithOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
