# ORM Performance Testing

- For benchmark performance between PGX, Gorm, Go-pg, Xorm on postgres.

- Gorm, Go-pg and Xorm using both their orm and exec raw sql.

- Benchmark with Insert, Update, Delete method.

## Getting Started

1. Start Postgres docker

```bash
$ ./start-docker.sh
```

2. Initialize database

```bash
$ ./migrate-db.sh
```

\*note if you cannot run script.

```bash
$ chmod 755 start-docker.sh migrate-db.sh
```

## Run Benchmark

- Benchmark Insert

```bash
$ go test -bench=BenchmarkInsertTest
```

- Benchmark Update

```bash
$ go test -bench=BenchmarkUpdateTest
```

- Benchmark Delete

```bash
$ go test -bench=BenchmarkDeleteTest
```

\*note if you want to run all benchmark.

```bash
$ go test -bench=.
```

## Result

- Benchmark Insert

![plot](./img/insert.jpg)

| Engine     | #1 (ns/op,B/op,allocs/op) | #2 (ns/op,B/op,allocs/op) | #3 (ns/op,B/op,allocs/op) | #avg (ns/op,B/op,allocs/op) |
| ---------- | ------------------------- | ------------------------- | ------------------------- | --------------------------- |
| PGX        | 2103503, 342, 10          | 1907730, 335, 10          | 2323930, 326, 10          | 2111721, 334, 10            |
| Gorm       | 4029284, 5803, 88         | 4383680, 5894, 88         | 4870865, 5820, 88         | 4427943, 5839, 88           |
| Gorm(raw)  | 1946235, 1295, 25         | 2090575, 1299, 25         | 2094899, 1311, 25         | 2043903, 1301.67, 25        |
| Go-pg      | 2012819, 972, 18          | 1987296, 971, 18          | 1950745, 971, 18          | 1983620, 97.33, 18          |
| Go-pg(raw) | 1999979, 276, 11          | 2178022, 276, 11          | 1826669, 275, 11          | 2001556.67, 275.67, 11      |
| Xorm       | 3539050, 3977, 89         | 3468434, 3975, 89         | 3440560, 3975, 89         | 3482681.33, 3975.67, 89     |
| Xorm(raw)  | 3518896, 2304, 44         | 3997685, 2305, 44         | 4397684, 2306, 44         | 3971421.67, 2305, 44        |

- Benchmark Update

- Benchmark Delete

## Contact

[Thanunya](mailto:b.beemmps@gmail.com)
