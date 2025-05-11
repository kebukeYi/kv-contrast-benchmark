package benchmk_1

import (
	"contrast-benchmark"
	"fmt"
	"github.com/rosedblabs/rosedb/v2"
	"testing"
)

var roseDB *rosedb.DB

func initRoseDB() {
	fmt.Println("init rosedb")
	opts := rosedb.DefaultOptions
	//opts.DirPath = filepath.Join("benchmark", "rosedb")
	opts.DirPath = "F:\\ProjectsData\\golang\\rosedb\\benchmark"
	contrast_benchmark.ClearDir(opts.DirPath)
	var err error
	roseDB, err = rosedb.Open(opts)
	if err != nil {
		panic(err)
	}
}

func initRoseDBData() {
	for i := 0; i < 500000; i++ {
		err := roseDB.Put(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue())
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_PutValue_RoseDB(b *testing.B) {
	initRoseDB()
	defer roseDB.Close()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := roseDB.Put(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue())
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_GetValue_RoseDB(b *testing.B) {
	initRoseDB()
	initRoseDBData()
	defer roseDB.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := roseDB.Get(contrast_benchmark.GetKey(i))
		if err != nil && err != rosedb.ErrKeyNotFound {
			panic(err)
		}
	}
}
