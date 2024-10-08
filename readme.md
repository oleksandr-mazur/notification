### Examples
https://github.com/Aristat/golang-example-app
https://github.com/topics/golang-app

[TOC]

# Syntax

## Variables

```go
s := ""
var s string
var s  = ""
var s string = ""
```

##  Constants

```go
const NAME string = "Hello"

const (
	one = 1
	two = 2
)
```

## Массивы
Массив имеет определенный набор элементов

```go
var block [3]int

block := [3]int

var block [...]init{1, 2, 3} //использовать длину при инициализации

block[1] = 4 // изменение значения

var block [3][3]int
```

## Slice
Срез это часть массива, срез может быть расширен

```go

var slice []init

slice = append(slice, 10) //добавление элемента в срез

var slice []init{1, 2, 3} // инициализации с данными

slice = append(slice, slice1...) // добавить данные среза slice1 в slice

var slice2 [][]int // слайс слайсов
slice2 = append(slice2, slice)

slice3 := make([]int, 5, 10) // создать слайс нужной длины 5 - длина, 10 - капасити

slice10 := make([]int, len(slice), len(slice)) // Создать слайс нужной длины
copy(slice10, slice) // копировать срез

```

## Map

```go

var mm map[string]int // обьявление

var mm map[string] map[string]int // обьявление + инициализации
mm := map[string]int{} // обьявление + инициализации
var mm = make(map[string]int) // обьявление + инициализации

mm["test"] = 10 // записать значение

ss := mm["test"] // получить значение

ss, ok := mm["test"] // получить значение + проверка

delete(ss, "test")  // удаление значения

```

## IOTA

**iota** - автоинкремент в рамках блока

```go
const (
	one = iota
	two
	tree
)

const (
	one int = iota * 1
	two
	tree
)
```

## IF/ELSE

Значение должно быть булевым

```go
if a == b {
    println(a)
}

if value, ok := range variable; ok {
    println(value)
} else {
    println("false")
}

if value, ok := range variable; ok {
    println(value)
} else  if value, ok := range variable2; ok {
    println("false")
} else {
    println("false")
}

```

## Цыклы FOR, SWITCH

```go
for {
    println("Бесконечный цыкл while")
}

for i :=1; i < 10; i++ {
    println(i)
}

for key, value := range mm {
    println(key)
}
```

**range** - индекс по значениям

```go
Mymetka:
    switch variable {
    case 1:
        println(1)
    case 2:
        println(2)
        fallthrough //перейти в следующий вариант
    case 3:
        println(3)
        break // выход из кейса
        break Mymetka // выход из цыкла по метке
    default:
        println(0)
```

## Функции FUNC

```go
func Test() int {
    return 10
}

func Test(param1 ...int) int { // неизвесное количество аргументов
    return (param...) неизвесное количество аргументов на выходе
}

func Test(param1 ...int) (res int) { // возврат переменной res
    res = 1
    return res
}

func init () {    // функция init выполняется во время начала
    some actions  // выполнения програмы
}

f := func () {    // анонимная функция
    some actions  // выполнения програмы
}
```

**defer** - выполнения после выполнения функции

```go
func Test(param1 int) int {
    fmt.Println("One")
    defer fmt.Println("Two")
    return (param)
}
```

```go
    func Ny() func() {
        fmt.Println("Func")
    }
    return func () {
        fmt.Println("Return func")
    }
```

## Структуры

```go
type Typp struct {
    name string
    }

// обьявление методов
type NFR int

func (m *NFR) add(i NFR) { // метод для типа Typp
    *m = *m + NFR(i)
}

type MySlice []int // тип может быть слайсом или мапом
```


#### Вложеность в go 

```go

type Type1 struct {
    name string
}

type Type2 struct {
    Type1
    LastName string
}

Type3 := Type2{LastName: "Ivanov"}
Type3.Type1.name = "Ivan"
```

## Интерфейсы

```go
type MyTyper interface { // интерфейсы определяют поведение
    MyFunc()             // принято добавлять суфикс -er
}                        // интерфейс не может содержать данные, только методы

MyTyper.(string) //  проверка на типа

var i interface{} = "Hello"


type error interface {
    Error() string
}
```

### Join

```go
strings.Join(os.Args[1:], " ")
```

# Library

### fmt
```go
fmt.Fprintf(os.Stderr, "Error: %v\n", err)
```

### os
```go
os.Stdin
os.Stdout
os.Open
```
### io/ioutils
```go
ioutils.ReadFile(filename)
```
### string
```go
strings.Split(string(data)., ''\n")
```
* **fmt**
* **bufio**
```go
scan := bufio.NewScanner(f)
for scan.Scan() {
    fmt.Println(scan.Text())
}
```