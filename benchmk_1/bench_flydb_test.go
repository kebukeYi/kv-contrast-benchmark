package benchmk_1

import (
	"contrast-benchmark"
	"fmt"
	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/engine"
	"github.com/ByteStorage/FlyDB/flydb"
	_const "github.com/ByteStorage/FlyDB/lib/const"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

var FlyDB *engine.DB
var err error

func initFlyDB() {
	fmt.Println("init FlyDB")
	opts := config.DefaultOptions
	opts.FIOType = config.FileIOType
	opts.DirPath = filepath.Join("benchmark", "flydb")
	//fmt.Printf("opts.DirPath: %s\n", opts.DirPath)
	//opts.DirPath = "F:\\ProjectsData\\golang\\flydb\\benchmark\\"
	//ClearDir(opts.DirPath)
	FlyDB, err = flydb.NewFlyDB(opts)
	if err != nil {
		panic(err)
	}
}

func Benchmark_PutValue_FlyDB(b *testing.B) {
	initFlyDB()
	defer FlyDB.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		err = FlyDB.Put(contrast_benchmark.GetKey(n), contrast_benchmark.GetValue())
		assert.Nil(b, err)
	}
}

func Benchmark_GetValue_FlyDB(b *testing.B) {
	initFlyDB()
	defer FlyDB.Close()

	for i := 0; i < 500000; i++ {
		err = FlyDB.Put(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue())
		assert.Nil(b, err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_, err = FlyDB.Get(contrast_benchmark.GetKey(n))
		if err != nil && err != _const.ErrKeyNotFound {
			panic(err)
		}
	}

}
