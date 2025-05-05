package benchmk

import (
	"contrast-benchmark"
	"fmt"
	"github.com/kebukeYi/TrainDB"
	"github.com/kebukeYi/TrainDB/lsm"
	"github.com/kebukeYi/TrainDB/model"
	"testing"
)

var triandb *TrainDB.TrainKVDB

func initTrainDB() {
	// fmt.Println("init TrainDB")
	dir := "./trainDB"
	//dir := "F:\\ProjectsData\\golang\\TrainDB\\benchmk"
	contrast_benchmark.ClearDir(dir)
	trianDB, err, _ := TrainDB.Open(lsm.GetLSMDefaultOpt(dir))
	if err != nil {
		panic(err)
	}
	triandb = trianDB
}

// -benchtime=60s -timeout=30m -count=3
func Benchmark_PutValue_TrainDB(b *testing.B) {
	initTrainDB()
	defer triandb.Close()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entry := model.NewEntry(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue())
		err = triandb.Set(entry)
		if err != nil {
			panic(err)
			return
		}
	}
}

func initTrainDBData() {
	batchSize := 50
	threshold := 500000
	batch := make([]*model.Entry, 0, batchSize)
	for i := 0; i < threshold; i++ {
		batch = append(batch, &model.Entry{
			Key:   model.KeyWithTs(contrast_benchmark.GetKey(i)),
			Value: contrast_benchmark.GetValue(),
		})

		if i%batchSize == 0 {
			err = triandb.BatchSet(batch)
			if err != nil {
				panic(err)
			}
			batch = batch[:0]
		}
	}
}

// -benchtime=60s -timeout=30m -count=3
func Benchmark_GetValue_TrainDB(b *testing.B) {
	initTrainDB()
	initTrainDBData()
	defer triandb.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getKey := contrast_benchmark.GetKey(i)
		_, err = triandb.Get(getKey)
		if err != nil {
			fmt.Printf("getKey: %s \n", getKey)
			panic(err)
		}
	}
}
