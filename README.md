# Contrast Benchmark

Based on the results of a simple benchmark test of the open source kv database written in Golang.

## Test database:

- [trainkv](https://github.com/kebukeYi/TrainKV)
- [flydb](https://github.com/ByteStorage/FlyDB)
- [bbolt](https://github.com/etcd-io/bbolt)
- [goleveldb](https://github.com/syndtr/goleveldb)
- [nutsdb](https://github.com/nutsdb/nutsdb)
- [rosedb](https://github.com/flower-corp/rosedb)
- [badger](https://github.com/dgraph-io/badger)
- [pebble](https://github.com/cockroachdb/pebble)

## Options:

- Value size: 512 bytes

## Results

```
goos: windows
goarch: amd64
pkg: contrast-benchmark_1
cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz

Benchmark_PutValue_GoLevelDB-4   	35443	     32395 ns/op	    2490 B/op	      12 allocs/op
Benchmark_PutValue_FlyDB-2-linux   	125092	      9319 ns/op	    2932 B/op	      18 allocs/op
Benchmark_PutValue_RoseDB-4   	        35956	     33159 ns/op	    2556 B/op	      17 allocs/op
Benchmark_PutValue_Badger-4   	        36817	     30879 ns/op	    5314 B/op	      49 allocs/op
Benchmark_PutValue_TrainKV-4   	        40033	     29274 ns/op	    6275 B/op	     130 allocs/op

Benchmark_PutValue_BoltDB-4   	        6697	    172776 ns/op	   22954 B/op	     139 allocs/op
Benchmark_PutValue_Pebble-4   	        44798	     22520 ns/op	    2340 B/op	       8 allocs/op

Benchmark_GetValue_GoLevelDB-4   	273182	      4387 ns/op	    1346 B/op	      15 allocs/op
Benchmark_GetValue_FlyDB-2-linux   	1410592	     788.7 ns/op	     374 B/op	       6 allocs/op
Benchmark_GetValue_RoseDB-4   	        110493	     10933 ns/op	    1326 B/op	      10 allocs/op
Benchmark_GetValue_Badger-4   	        115933	     10195 ns/op	    4862 B/op	      37 allocs/op
Benchmark_GetValue_TrainKV-4   	        61551	     17489 ns/op	   11810 B/op	     270 allocs/op
Benchmark_GetValue_BoltDB-4   	        498140	      2876 ns/op	     720 B/op	      25 allocs/op
Benchmark_GetValue_Pebble-4   	        33540	     38089 ns/op	   14859 B/op	      20 allocs/op
````


```
goos: linux
goarch: amd64
pkg: contrast-benchmark_1
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz

Benchmark_PutValue_RoseDB
Benchmark_PutValue_RoseDB-16       	   69776	     19166 ns/op	    6242 B/op	      59 allocs/op
Benchmark_GetValue_RoseDB
Benchmark_GetValue_RoseDB-16       	 4155183	     298.0 ns/op	     167 B/op	       4 allocs/op

Benchmark_PutValue_FlyDB
Benchmark_PutValue_FlyDB-16        	   95023	     13763 ns/op	    2904 B/op	      16 allocs/op
Benchmark_GetValue_FlyDB
Benchmark_GetValue_FlyDB-16    	 	 2710143	     463.5 ns/op	     259 B/op	       5 allocs/op

Benchmark_PutValue_GoLevelDB
Benchmark_PutValue_GoLevelDB-16    	   71931	     14709 ns/op	    2226 B/op	      12 allocs/op
Benchmark_GetValue_GoLevelDB
Benchmark_GetValue_GoLevelDB-16    	  500736	      2520 ns/op	    1278 B/op	      15 allocs/op


Benchmark_PutValue_Badger
Benchmark_PutValue_Badger-16       	   59331	     22711 ns/op	    6006 B/op	      48 allocs/op
Benchmark_GetValue_Badger
Benchmark_GetValue_Badger-16       	  158686	      7686 ns/op	   10844 B/op	      42 allocs/op

Benchmark_PutValue_BoltDB
Benchmark_PutValue_BoltDB-16       	   32637	     56519 ns/op	   21009 B/op	     123 allocs/op
Benchmark_GetValue_BoltDB
Benchmark_GetValue_BoltDB-16       	  655971	     24327 ns/op	     723 B/op	      26 allocs/op 

Benchmark_PutValue_NutsDB
Benchmark_PutValue_NutsDB-16       	   78801	     13582 ns/op	    3242 B/op	      22 allocs/op
Benchmark_GetValue_NutsDB
Benchmark_GetValue_NutsDB-16       	  373124	      5702 ns/op	    1392 B/op	      14 allocs/op

Benchmark_PutValue_Pebble
Benchmark_PutValue_Pebble-16       	   91304	     21877 ns/op	    2720 B/op	       8 allocs/op
Benchmark_GetValue_Pebble
Benchmark_GetValue_Pebble-16       	   66135	     15837 ns/op	   17193 B/op	      22 allocs/op
PASS
```

