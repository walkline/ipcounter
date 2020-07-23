# ipcounter

Counts the unique IP addresses that comes to the `POST /logs` endpoint and provides metrics for Prometheus.
Supports only IPv4 format.

#### RUN
To run service use `$ make run` command.

### Tests
To run unit tests you can use `$ make test` command.

### Benchmarks
To run go bencmarks use `$ make bench`.

### Load tests
You can run load tests for `POST /logs` using [vegeta](https://github.com/tsenart/vegeta).
Samples:
```
$ ./vegeta-targets-generator.sh | vegeta attack -format=json -rate=0 -max-workers=100 -duration=30s | vegeta report
$ ./vegeta-targets-generator.sh | vegeta attack -format=json -rate=5000 -duration=30s | vegeta report
```
`vegeta-targets-generator.sh` generates 1000 vegetas targets with 1000 random ip addresss.

Load tests result for `./vegeta-targets-generator.sh | vegeta attack -format=json -rate=0 -max-workers=100 -duration=30s | vegeta report` on MacBook Pro (2016) 2,6 GHz Intel Core i7:
```
Requests      [total, rate, throughput]  1082112, 36070.47, 36070.27
Duration      [total, attack, wait]      30.000108397s, 29.999941837s, 166.56µs
Latencies     [mean, 50, 95, 99, max]    297.355µs, 194.033µs, 830.594µs, 1.821919ms, 49.776776ms
Bytes In      [total, mean]              0, 0.00
Bytes Out     [total, mean]              25182918, 23.27
Success       [ratio]                    100.00%
Status Codes  [code:count]               202:1082112
```
