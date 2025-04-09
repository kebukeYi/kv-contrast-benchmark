package benchmk

import (
	"contrast-benchmark"
	"fmt"
	"github.com/cockroachdb/pebble"
	"math"
	"testing"
)

var pebbledb *pebble.DB

func initPebbleDB() {
	fmt.Println("init pebble")
	//dir := filepath.Join("benchmark", "pebble")
	dir := "F:\\ProjectsData\\golang\\pebble\\benchmark"
	contrast_benchmark.ClearDir(dir)
	opts := &pebble.Options{
		BytesPerSync: math.MaxInt,
	}
	var err error
	pebbledb, err = pebble.Open(dir, opts)
	if err != nil {
		panic(err)
	}
}

func Benchmark_PutValue_Pebble(b *testing.B) {
	initPebbleDB()
	defer pebbledb.Close()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		pebbledb.Set(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue(), &pebble.WriteOptions{
			Sync: false,
		})
	}
}

func Benchmark_GetValue_Pebble(b *testing.B) {
	initPebbleDB()
	defer pebbledb.Close()

	for i := 0; i < 500000; i++ {
		pebbledb.Set(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue(), &pebble.WriteOptions{
			Sync: false,
		})
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pebbledb.Get(contrast_benchmark.GetKey(i))
	}
}
