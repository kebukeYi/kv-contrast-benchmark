package benchmk_2

import (
	"bytes"
	contrast_benchmark "contrast-benchmark"
	"fmt"
	"github.com/kebukeYi/TrainKV/model"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// -benchtime=300s -count=1 -timeout=30m
func BenchmarkDiskUsage(b *testing.B) {
	b.Run("batchPut4K", benchmarkDiskUsageBatchPut)

	b.Run("concurrentBatchPut4K", benchmarkDiskUsageConcurrentBatchPut)
}

func getActualDiskUsage(path string) int64 {
	cmd := exec.Command("du", "-sb", path)

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0
	}

	parts := strings.Fields(out.String())
	if len(parts) < 1 {
		return 0
	}

	size, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0
	}

	return size
}

// print disk usgae per three seconds
var (
	totalBytesWritten int64
	stopCh            chan struct{}
)

func printDiskUsageStat() {
	fmt.Printf("\n")

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	var lastTotal int64
	for {
		select {
		case <-ticker.C:
			current := atomic.LoadInt64(&totalBytesWritten)
			speed := current - lastTotal
			lastTotal = current
			fmt.Printf(
				"Write Speed %.2f MB/s | Write Total: %.2f GB | Disk Usage: %.2f GB\n",
				float64(speed)/1024/1024/3,
				float64(current)/1024/1024/1024,
				float64(getActualDiskUsage(dir))/1024/1024/1024,
			)
		case <-stopCh:
			return
		}
	}
}

func benchmarkDiskUsageBatchPut(b *testing.B) {
	initTrainDB()
	defer db.Close()

	totalBytesWritten = 0
	stopCh = make(chan struct{})
	defer close(stopCh)

	go printDiskUsageStat()

	b.ResetTimer()
	b.ReportAllocs()

	batch := make([]*model.Entry, 0, BatchSize)
	for i := 0; i < b.N; i++ {
		entry := model.NewEntry(contrast_benchmark.GetKey(i), bin4KB)
		entry.Key = model.KeyWithTs(entry.Key)

		batch = append(batch, entry)
		estimateSize := entry.EstimateSize(1 << 20)
		atomic.AddInt64(&totalBytesWritten, int64(estimateSize))
		if i%BatchSize == 0 {
			err := db.BatchSet(batch)
			assert.Nil(b, err)
			batch = batch[:0]
		}
	}
	time.Sleep(6 * time.Second)
}

func benchmarkDiskUsageConcurrentBatchPut(b *testing.B) {
	initTrainDB()
	defer db.Close()

	totalBytesWritten = 0
	stopCh = make(chan struct{})
	defer close(stopCh)

	go printDiskUsageStat()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		var err error
		iteration := 0
		batch := make([]*model.Entry, 0, BatchSize)
		for pb.Next() {
			entry := model.NewEntry(contrast_benchmark.GetKey(iteration), bin4KB)
			entry.Key = model.KeyWithTs(entry.Key)

			batch = append(batch, entry)
			estimateSize := entry.EstimateSize(1 << 20)
			atomic.AddInt64(&totalBytesWritten, int64(estimateSize))
			if iteration%BatchSize == 0 {
				err = db.BatchSet(batch)
				assert.Nil(b, err)
				batch = batch[:0]
			}
			iteration++
		}
	})
	time.Sleep(6 * time.Second)
}
