# Функции

Это локально-определенный блок кода , имеющий имя (ЯВНОЕ ОПРЕДЕЛЕНИЕ)  
Функцию необходимо определить + Функциию необходимо вызвать  
**Сигнатура**

```go
func functionName(param, param2 int) int {
  //body
}
```

## Variadic parameters

```go
func someStrings(a, b int, words ...string) {
    //body
}
```

## Анонимная функция

```go
 anon := func(a, b int) int {
  return a + b
 }
```

## Возврат нескольких значений

```go
 anon := func(a, b int) (int, string) {
  return paramInt, paramString
 }
```
