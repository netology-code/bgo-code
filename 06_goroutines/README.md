# Трассировка

Для трассировки в `main` используется стандартный пакет `trace`:

```go
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()
```

После запуска и появления файла `trace.out` запустить просмотрщик:

```shell
go tool trace trace.out
```

# Benchmark'и

Для бенчмарков используется команда запуска:

```shell
go test -bench=. -benchtime=1000x ./...
```
