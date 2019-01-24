Benchmark this, like so:

```shell
go test -benchtime=10 -bench=.
```

Pool is working with things that are expensive to create can drastucally improve response time.