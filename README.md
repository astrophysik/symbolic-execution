# Символьное исполнение — 2026

Домашние задания курса и общая кодовая база символьного анализатора Go.
Задания и каркас перенесены из [курса 2025, main@7d98354](https://github.com/CaelmBleidd/symbolic_execution_2025/tree/7d983541d76dd257596a01ade93cf2f7e4c63b6d).

## Домашние задания

| Задание | Тема |
| ------- | ---- |
| [ДЗ1](homework1/README.md) | Построение SSA |
| [ДЗ2](homework2/README.md) | Символьные выражения и трансляция в SMT |
| [ДЗ3](homework3/README.md) | Внутрипроцедурный анализ основных конструкций |
| [ДЗ4](homework4/README.md) | Символьная память и алиасинг |
| [ДЗ5](homework5/README.md) | Генерация тестов |
| [ДЗ6](homework6/README.md) | Межпроцедурный анализ оставшихся конструкций |

Реализации развиваются в общей директории `internal/`. В папках заданий находятся условия,
примеры и программы запуска. [Итоговый набор примеров](final_tests/README.md) предназначен
для проверки готового анализатора. Заготовки содержат TODO; решения заданий не включены.

## Требования

- Go 1.26.2 или новее, как в исходном репозитории с заданиями.
- C-компилятор и включённый cgo.
- Z3: библиотека и заголовочные файлы.

macOS:

```bash
brew install z3
export CGO_CFLAGS="-I$(brew --prefix z3)/include"
export CGO_LDFLAGS="-L$(brew --prefix z3)/lib"
```

Ubuntu/Debian:

```bash
sudo apt-get install build-essential libz3-dev
```

## Запуск

```bash
git clone https://github.com/CaelmBleidd/symbolic_execution_2026.git
cd symbolic_execution_2026
make deps
make build
make test
make examples
```

`make lint` запускает `go vet` для каркаса и программ запуска, `make fmt` форматирует код,
`make coverage` строит отчёт о покрытии.
Все команды доступны через `make help`.

Успешные `make build` и `make test` подтверждают сборку каркаса и работу существующих тестов
обёртки Z3, но не выполнение домашних заданий. Проверки конкретного задания описаны в его README.
Файлы в `homework*/examples/` служат входом анализатора; некоторые имеют `package main`
без функции `main`, поэтому не являются самостоятельными исполняемыми программами.
Входные примеры намеренно содержат недостижимый и другой проблемный код, поэтому
`make lint` не проверяет их и `final_tests/`.

## Структура

```text
homework1/ … homework6/        # Условия заданий, примеры и программы запуска
internal/                     # Общий каркас анализатора с TODO
final_tests/                  # Примеры для приёмочной проверки анализатора
examples/basic_z3_example.go   # Уравнения, модель и push/pop
pkg/z3wrapper/solver.go        # Обёртка над Z3
pkg/z3wrapper/solver_test.go   # Тесты обёртки
go.mod, go.sum                # Зависимости Go
Makefile                     # Сборка и проверки
```
