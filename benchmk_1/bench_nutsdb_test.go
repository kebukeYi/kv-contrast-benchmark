package benchmk_1

import (
	"contrast-benchmark"
	"fmt"
	"github.com/nutsdb/nutsdb"
	"testing"
)

var nutsDB *nutsdb.DB

func initNutsDB() {
	fmt.Println("init nutsdb")
	opts := nutsdb.DefaultOptions
	//opts.Dir = "benchmark/nutsdb"
	opts.Dir = "F:\\ProjectsData\\golang\\nutsdb\\benchmark"
	contrast_benchmark.ClearDir(opts.Dir)
	opts.SyncEnable = false
	opts.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	var err error
	nutsDB, err = nutsdb.Open(opts)
	if err != nil {
		panic(err)
	}
}

func initNutsDBData() {
	for i := 0; i < 500000; i++ {
		nutsDB.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put("test-bucket", contrast_benchmark.GetKey(i), contrast_benchmark.GetValue(), 0)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func Benchmark_PutValue_NutsDB(b *testing.B) {
	initNutsDB()
	defer nutsDB.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nutsDB.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put("test-bucket", contrast_benchmark.GetKey(i), contrast_benchmark.GetValue(), 0)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func Benchmark_GetValue_NutsDB(b *testing.B) {
	initNutsDB()
	initNutsDBData()
	defer nutsDB.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nutsDB.View(func(tx *nutsdb.Tx) error {
			_, err := tx.Get("test-bucket", contrast_benchmark.GetKey(i))
			if err != nil && err != nutsdb.ErrKeyNotFound {
				panic(err)
			}
			return nil
		})
	}
}
