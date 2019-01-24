Benchmark this, like so:

```shell
go test -benchtime=10 -bench=.
```

Pool is working with things that are expensive to create can drastucally improve response time.
The object pool design pattern is best used either when you have concurrent processes that require objects, but dispose of them 
very repidly after instantiation, or when construction of these objects could negatively impact memory.