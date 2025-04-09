package benchmk

import (
	"contrast-benchmark"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"testing"
)

var badgerdb *badger.DB

func initBadgerDB() {
	fmt.Println("init badgerdb")
	dir := "F:\\ProjectsData\\golang\\badger\\benchmark"
	// ClearDir(dir)
	opts := badger.DefaultOptions(dir)
	opts.SyncWrites = false
	badgerdb, err = badger.Open(opts)
	if err != nil {
		panic(err)
	}
}

func Benchmark_PutValue_Badger(b *testing.B) {
	initBadgerDB()
	defer badgerdb.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		badgerdb.Update(func(txn *badger.Txn) error {
			return txn.Set(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue())
		})
	}
}

func Benchmark_GetValue_Badger(b *testing.B) {
	initBadgerDB()
	defer badgerdb.Close()

	for i := 0; i < 500000; i++ {
		badgerdb.Update(func(txn *badger.Txn) error {
			return txn.Set(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue())
		})
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		badgerdb.View(func(txn *badger.Txn) error {
			txn.Get(contrast_benchmark.GetKey(i))
			return nil
		})
	}
}
