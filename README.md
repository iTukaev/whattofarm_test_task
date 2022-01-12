# whattofarm_test_task

# Написать web-приложение (Go + Mongo DB), осуществляющее подсчёт запросов счётного пикселя в формате gif.

Ввозвращается минимальный однопиксельный gif размером 33 байта:
47 49 46 38 39 61 01 00 01 00 00 00 00 21 F9 04 01 00 00 00 00 2C 00 00 00 00 01 00 01 00 00 02 00
с валидным заголовком Content-type).

В качестве параметра при запросе пикселя могут передаваться:

1. Тип действия (view, click, install, ...)
2. Страна в формате (ru, gb, ...)

Пример запроса:
http://localhost:8000/counter.gif?action=view&country=ru

В базе данных данные вызовы должны приводить к подсчёту таких вызовов с группировкой по действиям и странам.

Примерный формат итоговой записи в базе данных:

{
  "_id": ObjectId("..."),
  "total": 123,
  "actions": {
    "view": {
      "total": 100
    },
    "click": {
      "total": 20
    },
    "install": {
      "total": 3
    },
    ...
  },
  "countries": {
    "ru": {
      "total": 50
    },
    "fr": {
      "total": 17
    },
    ...
  }
}

### Дополнительное задание: дополнить вложенные срезы данных:

{
  "_id": ObjectId("..."),
  "total": 123,
  "actions": {
    "view": {
      "total": 100
      "countries": {
        "ru": {
          "total": 80
        },
        "fr": {
          "total": 10
        },
        ...
      }
    },
    "click": {
      "total": 20,
        "countries": {
          "ru": {
            "total": 10
          },
          "fr": {
            "total": 5
          },
          ...
        }
      },
      ...
    },
    "countries": {
      "ru": {
        "total": 50,
        "actions": {
          "view": {
            "total": 40
          },
          "click": {
            "total": 40
          },
          ...
        }
      },
      ...
    }
}

### Дополнительное задание 2:

Сделать все то же самое, только с разбивкой на часовые бины, когда данные накапливаются целый час, после чего создается новый объект.
Отдельный плюс - через функцио агрегации собрать данные из данных часовых бинов в заданный фрагмент (например день с учётом временной зоны)