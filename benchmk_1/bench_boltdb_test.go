package benchmk_1

import (
	"contrast-benchmark"
	"fmt"
	"go.etcd.io/bbolt"
	"os"
	"path/filepath"
	"testing"
)

var boltDB *bbolt.DB

func initBoltDB() {
	fmt.Println("init boltDB")
	opts := bbolt.DefaultOptions
	opts.NoSync = true
	var err error
	path := "F:\\ProjectsData\\golang\\boltdb\\benchmark"
	contrast_benchmark.ClearDir(path)
	_ = os.MkdirAll(path, os.ModePerm)
	filePath := filepath.Join(path, "bolt.data")
	boltDB, err = bbolt.Open(filePath, 0644, opts)
	if err != nil {
		panic(err)
	}

	boltDB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket([]byte("test-bucket"))
		if err != nil {
			panic(err)
		}
		return nil
	})
	initBoltDBData()
}

func initBoltDBData() {
	var k int
	for i := 0; i < 5; i++ {
		boltDB.Update(func(tx *bbolt.Tx) error {
			for j := 0; j < 100000; j++ {
				err := tx.Bucket([]byte("test-bucket")).Put(contrast_benchmark.GetKey(k), contrast_benchmark.GetValue())
				if err != nil {
					panic(err)
				}
				k++
			}
			return nil
		})
	}
}

func Benchmark_PutValue_BoltDB(b *testing.B) {
	initBoltDB()
	defer boltDB.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		boltDB.Update(func(tx *bbolt.Tx) error {
			err := tx.Bucket([]byte("test-bucket")).Put(contrast_benchmark.GetKey(i), contrast_benchmark.GetValue())
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func Benchmark_GetValue_BoltDB(b *testing.B) {
	initBoltDB()
	defer boltDB.Close()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		boltDB.View(func(tx *bbolt.Tx) error {
			tx.Bucket([]byte("test-bucket")).Get(contrast_benchmark.GetKey(i))
			return nil
		})
	}
}
