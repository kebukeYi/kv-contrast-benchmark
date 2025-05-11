package benchmk_1

import (
	"contrast-benchmark"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"log"
	"testing"
)

var (
	levelDb *leveldb.DB
)

func initLevelDB() {
	fmt.Println("init leveldb")
	//dir := "benchmark/leveldb"
	dir := "F:\\ProjectsData\\golang\\leveldb\\benchmark"
	contrast_benchmark.ClearDir(dir)
	var err error
	levelDb, err = leveldb.OpenFile(dir, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initLevelDBValue() {
	for i := 0; i < 500000; i++ {
		key := contrast_benchmark.GetKey(i)
		val := contrast_benchmark.GetValue()
		err := levelDb.Put(key, val, nil)
		if err != nil {
			log.Fatal("leveldb write data err.", err)
		}
	}
}

func Benchmark_PutValue_GoLevelDB(b *testing.B) {
	initLevelDB()
	defer levelDb.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := contrast_benchmark.GetKey(i)
		val := contrast_benchmark.GetValue()
		err := levelDb.Put(key, val, &opt.WriteOptions{Sync: false})
		if err != nil {
			log.Fatal("leveldb write data err.", err)
		}
	}
}

func Benchmark_GetValue_GoLevelDB(b *testing.B) {
	initLevelDB()
	initLevelDBValue()
	defer levelDb.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := levelDb.Get(contrast_benchmark.GetKey(i), nil)
		if err != nil && err != leveldb.ErrNotFound {
			log.Fatal("leveldb read data err.", err)
		}
	}
}
