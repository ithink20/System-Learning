### Request payload parsing optimisation

#### Introduction

Gin is a high-performance HTTP web framework written in Golang.
It has a martini-like API and claims to be up to 40 times faster. Gin contains a set of commonly used functionalities, such as: Routing, Middleware support, Rendering.

In Gin, **BindJSON** reads the body buffer and de-serializes it into a struct.


#### Technical Details

* In our service, we use gin BindJson() to parse the http request payload and de-serialize into a struct
* This bindJSON() internally decode_json which in overall costs ~8% of our CPU

This marks a crucial juncture where optimization is paramount, aiming for systems with both high query per second (QPS) and optimal performance.


#### Benchmarking Details
 
1. Convert Request payload from JSON key-val string to Json Array format
2. Compare the performance of de-serialization

- Using String Parser (only string operations - no library use to make it faster)
- Using Bytedance/Sonic Unmarshal
- Using Json Unmarshal


#### Result
```
goos: darwin
goarch: amd64
pkg: golang/pkg/worker-pool/golang/pkg/HttpRequest
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkStringParse-12    	 3219116	       370.1 ns/op
BenchmarkStringParse-12    	 3249013	       374.3 ns/op
BenchmarkStringParse-12    	 3285272	       367.0 ns/op
BenchmarkUnmarshal-12      	   59209	     20213 ns/op
BenchmarkUnmarshal-12      	   58171	     20431 ns/op
BenchmarkUnmarshal-12      	   59199	     20273 ns/op
BenchmarkSonic-12          	  348633	      2916 ns/op
BenchmarkSonic-12          	  366554	      2944 ns/op
BenchmarkSonic-12          	  364891	      3023 ns/op
PASS
```
As we can observe, Simple String parsing is more faster even from sonic

* Addition test for checking performance on URL encode/decode

```
BenchmarkEncode-12             	   57801	     20486 ns/op
BenchmarkEncode-12             	   56664	     20427 ns/op
BenchmarkEncode-12             	   57770	     20752 ns/op
BenchmarkDecode-12             	  436692	      2767 ns/op
BenchmarkDecode-12             	  441338	      2712 ns/op
BenchmarkDecode-12             	  426259	      2729 ns/op
BenchmarkEncodeUnmarshal-12    	   18840	     63807 ns/op
BenchmarkEncodeUnmarshal-12    	   18798	     63353 ns/op
BenchmarkEncodeUnmarshal-12    	   18712	     63218 ns/op
PASS
```

Interesting: Encoding is expensive than decoding

url.QueryEscape() --> can be optimised more
-  https://github.com/VictoriaMetrics/VictoriaMetrics/issues/2114
- https://github.com/golang/go/issues/17860