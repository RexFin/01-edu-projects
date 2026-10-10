# Go-Refresh

**Go-Refresh** - это консольный инструмент на языке Go, предназначенный для автоматического редактирования, исправления и форматирования текста. Программа работает как пайплайн: читает текст из файла, применяет набор строгих правил форматирования и безопасно сохраняет результат в новый файл.

---

## Основные возможности

### 1. Преобразование систем счисления
* **(hex)**: Преобразует предшествующее шестнадцатеричное число в десятичное.
    * *Пример:* `1E (hex)` → `30`
* **(bin)**: Преобразует предшествующее двоичное число в десятичное.
    * *Пример:* `10 (bin)` → `2`

### 2. Изменение регистра (Case Modification)
* **(up)** / **(up, n)**: Переводит слово (или `n` предшествующих слов) в ВЕРХНИЙ РЕГИСТР.
* **(low)** / **(low, n)**: Переводит слово (или `n` предшествующих слов) в нижний регистр.
* **(cap)** / **(cap, n)**: Пишет слово (или `n` предшествующих слов) С Заглавной Буквы.
    * *Пример:* `This is so exciting (up, 2)` → `This is SO EXCITING`

>**⚠️ Важные правила работы модификаторов:**  
> * **Строгий синтаксис**: Модификаторы должны быть написаны без лишних пробелов внутри скобок. Например, `(up, 5)` сработает корректно, а сломанные форматы вроде `(up , 5)` или `( up, 5 )` будут проигнорированы как обычный текст.
> * **Многострочность**: Если указано число слов (например, `(up, 100)`), действие модификатора распространяется на предшествующие слова даже через переносы строк (переносы строк при подсчете слов назад игнорируются).

### 3. Коррекция пунктуации и грамматики
* **Знаки препинания**: Символы `. , ! ? : ;` автоматически прижимаются к левому слову и отделяются пробелом справа.
* **Одинарные кавычки**: Текст внутри `' '` очищается от лишних пробелов по краям (например, `' awesome '` → `'awesome'`).
* **Артикли**: Автоматически заменяет `a` на `an`, если следующее за ним слово начинается с гласной (`a, e, i, o, u`) или буквы `h`.

### 4. Анализ текста (Дополнительные флаги)
Утилита поддерживает передачу опциональных флагов для анализа текста. Результат выводится в консоль:
* **Флаг `--lang`**: Анализирует текст и выводит предполагаемый язык (English / French) и процент уверенности.
* **Флаг `--keywords`**: Находит и выводит топ-5 самых часто используемых слов в тексте (без учета регистра и общих стоп-слов).

> **Примечание:** В целях безопасности и защиты оперативной памяти установлен лимит на обработку файлов — **до 5 МБ**.

---

## 🛠 Установка

Для работы потребуется установленный [Go](https://go.dev/dl/) (версии 1.26 или выше).

1. Клонируйте репозиторий:
```bash
git clone https://01.tomorrow-school.ai/git/nkishiba/go-refresh.git
```

2. Перейдите в директорию проекта:
```bash
cd go-refresh
```

## Использование

Программа требует указать два обязательных аргумента: путь к входному файлу (`input.txt`) и путь для сохранения результата (`output.txt`).

### Запуск тестов:
```go
go test ./...
```

### Стандартный запуск:
```go
go run cmd/go-refresh/main.go samples/sample1.in.txt samples/result.txt
```

### Запуск с флагами анализа:
```go
go run cmd/go-refresh/main.go samples/sample2.in.txt samples/result.txt --lang
```
Примерный результат:
```bash
Language: English (100%)
Success! Processed text saved to samples/result.txt
```

### Запуск через Makefile
Вы также можете использовать команды `make` для быстрой работы (аргументы передаются через переменные):
```bash
make run IN=samples/sample3.in.txt OUT=samples/result.txt LANG_FLAG=--keywords
```

В консоле будет:
```bash
Keywords: was, there, i, rock, 1e
Success! Processed text saved to samples/result.txt
```

## Пример работы (До / После)

**sample6.in.txt**
```txt
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how did you get in my house (cap, 6) ? "
I am exactly how they describe me: ' awesome '
it (cap) was 10 (bin) years ago.
```

**result.txt**
```txt
If I make you breakfast in bed just say thank you instead of: how Did You Get In My House?
I am exactly how they describe me: 'awesome'
It was 2 years ago.
```

## Структура проекта
```txt
go-refresh/
├─ cmd/
│  └─ go-refresh/
│     └─ main.go        // Точка входа программы
├─ internal/
│  ├─ ai/               // Анализ языка в тексте
│  │  ├─ langdetect_test.go
│  │  └─ langdetect.go
│  ├─ cases/            // Обработка регистр букв
│  │  ├─ transform_test.go
│  │  └─ transform.go
│  ├─ grammar/          // Проверка артикул
│  │  ├─ articles_test.go
│  │  └─ articles.go
│  ├─ modfields/        // Подготовка модфикаторов
│  │  ├─ ModificationFix.go
│  │  └─ ModificationFix_test.go
│  ├─ numbers/          // Конвертация чисел
│  │  ├─ convert_test.go
│  │  └─ convert.go
│  ├─ parser/           // Чтения символов из файла по байтам
│  │  └─ ParseTextFromFile.go
│  │  ├─ ParseTextFromFile_test.go
│  ├─ pipeline/         // Упарвление порядком действии
│  │  ├─ pipeline.go
│  │  └─ pipeline_test.go
│  ├─ punct/            // Обработка пунктуации
│  │  ├─ spacing_test.go
│  │  └─ spacing.go
│  ├─ quotes/           // Обработка кавычки
│  │  ├─ singlequotes_test.go
│  │  └─ singlequotes.go
│  └─ wordcount/        // Нахождение 5 часто используемых слов
│     ├─ frequency_test.go
│     └─ frequency.go
├─ samples              // Примеры
│  ├─ sample1.in.txt
│  ├─ sample1.out.txt
│  ├─ sample2.in.txt
│  ├─ sample2.out.txt
│  ├─ sample3.in.txt
│  ├─ sample3.in.txt
│  ├─ sample4.in.txt
│  ├─ sample5.in.txt    // Размер данного файла превышен 5мб
│  ├─ sample6.in.txt
│  └─ sample6.out.txt
├─ go.mod
├─ Makefile
└─ README.md
```

## Автор
Автор: RexFin  
Проект выполнен в рамках обучения в Tomorrow School