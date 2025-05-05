package benchmk_2

import (
	"contrast-benchmark"
	"crypto/sha1"
	"fmt"
	"github.com/kebukeYi/TrainDB"
	"github.com/kebukeYi/TrainDB/lsm"
	"github.com/kebukeYi/TrainDB/model"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync/atomic"
	"testing"
)

var (
	dir string
	db  *TrainDB.TrainKVDB
)

const BatchSize = 50

func genTestKey(i int) []byte {
	return []byte(strconv.Itoa(i))
}

func Gen1KBytes() []byte {
	buf := make([]byte, 1024)
	for i := 0; i < 128; i++ {
		copy(buf[i*8:], []byte("01234567"))
	}
	return buf
}

func GenNKBytes(n int) []byte {
	bytes1KB := Gen1KBytes()
	buf := make([]byte, 1024*n)
	for i := 0; i < n; i++ {
		copy(buf[i*1024:], bytes1KB)
	}
	return buf
}

var (
	bin4KB = GenNKBytes(4)
	ns     = sha1.Sum([]byte("benchmark"))
)

func initTrainDB() {
	// fmt.Println("init TrainDB")
	dir := "./TrainDB"
	contrast_benchmark.ClearDir(dir)
	trainDB, err, _ := TrainDB.Open(lsm.GetLSMDefaultOpt(dir))
	if err != nil {
		panic(err)
	}
	db = trainDB
}

// -benchtime=60s -count=3 -timeout=50m
func BenchmarkPutGet(b *testing.B) {
	b.Run("put4K", benchmarkPut)           // ok
	b.Run("batchPut4K", benchmarkBatchPut) // ok

	b.Run("get4K", benchmarkGet) //

	b.Run("concurrentPut4K", benchmarkConcurrentPut)
	b.Run("concurrentGet4K", benchmarkConcurrentGet)

	b.Run("concurrentBatchPut4K", benchmarkConcurrentBatchPut)
}

func benchmarkPut(b *testing.B) {
	initTrainDB()
	defer db.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entry := model.NewEntry(contrast_benchmark.GetKey(i), bin4KB)
		err := db.Set(entry)
		assert.Nil(b, err)
	}
}

func benchmarkConcurrentPut(b *testing.B) {
	initTrainDB()
	defer db.Close()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		iteration := 0
		for pb.Next() {
			entry := model.NewEntry(contrast_benchmark.GetKey(iteration), bin4KB)
			err := db.Set(entry)
			assert.Nil(b, err)
			iteration++
		}
	})
}

func benchmarkBatchPut(b *testing.B) {
	initTrainDB()
	defer db.Close()

	b.ResetTimer()
	b.ReportAllocs()

	batch := make([]*model.Entry, 0, BatchSize)
	for i := 0; i < b.N; i++ {
		entry := model.NewEntry(contrast_benchmark.GetKey(i), bin4KB)
		entry.Key = model.KeyWithTs(entry.Key)
		batch = append(batch, entry)
		if i%BatchSize == 0 {
			err := db.BatchSet(batch)
			assert.Nil(b, err)
			batch = batch[:0]
		}
	}

	if len(batch) != 0 {
		err := db.BatchSet(batch)
		assert.Nil(b, err)
	}
}

func getPrepare(b *testing.B) {
	batch := make([]*model.Entry, 0, BatchSize)
	for i := 0; i < 100001; i++ {
		entry := model.NewEntry(contrast_benchmark.GetKey(i), bin4KB)
		entry.Key = model.KeyWithTs(entry.Key)

		batch = append(batch, entry)

		if i%BatchSize == 0 {
			err := db.BatchSet(batch)
			assert.Nil(b, err)
			batch = batch[:0]
		}
	}
}

func benchmarkGet(b *testing.B) {
	initTrainDB()
	defer db.Close()

	getPrepare(b)

	b.ResetTimer()
	b.ReportAllocs()
	total := 0
	for i := 0; i < b.N; i++ {
		_, err := db.Get(contrast_benchmark.GetKey(i % 100000))
		if err != nil {
			total++
			// assert.Nilf(b, err, "i: %v, err: %v", i, err)
		}
	}
	fmt.Println("benchmarkGet.total: ", total)
}

func benchmarkConcurrentGet(b *testing.B) {
	initTrainDB()
	defer db.Close()

	getPrepare(b)

	b.ResetTimer()
	b.ReportAllocs()
	var total int32
	b.RunParallel(func(pb *testing.PB) {
		iteration := 0
		for pb.Next() {
			_, err := db.Get(contrast_benchmark.GetKey(iteration % 100000))
			if err != nil {
				atomic.AddInt32(&total, 1)
				// assert.Nilf(b, err, "i: %v, err: %v", i, err)
			}
			// assert.Nil(b, err)
			iteration++
		}
	})
	fmt.Println("benchmarkConcurrentGet.total: ", total)
}

func benchmarkConcurrentBatchPut(b *testing.B) {
	initTrainDB()
	defer db.Close()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		iteration := 0
		batch := make([]*model.Entry, 0, BatchSize)
		for pb.Next() {

			entry := model.NewEntry(contrast_benchmark.GetKey(iteration), bin4KB)
			entry.Key = model.KeyWithTs(entry.Key)

			batch = append(batch, entry)

			if iteration%BatchSize == 0 {
				err := db.BatchSet(batch)
				assert.Nil(b, err)
				batch = batch[:0]
			}

			iteration++
		}
	})
}
