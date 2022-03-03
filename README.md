# testtask


Куда смотреть? 

Под каждое задание своя папка:
RomanNum, Anagrams,.... 

Установить go
Перейти в папку (RomanNum/Anagrams/...) 
В терминале запустить go test

Чтобы посмотреть, что находится в тестах:

Перейти в файл с постфиксом _test

Можно добавить свой блок или отредактировать уже имеющийся 
В args вставить то, что в функцию хотите подать на входе, 
В want - то, что хотите получить в результате
 WantErr - bool параметр. Ожидается ли от этого теста ошибка или нет



Each task has its own folder:
RomanNum, Anagrams,....

Install go
Go to task folder (RomanNum/Anagrams/...)
In terminal run "go test"

To see what's in the tests:

Go to file with postfix _test
You can add your own block or edit an existing one.
In args, insert what you want to feed into the function at the input,
In want - what you want to get as a result
  WantErr - bool parameter. Whether this test is expected to fail or not


Задания:

 1) Написать функцию, которая принимает в качестве аргумента десятичное число и возвращает число, написанное римскими числами. (вариации задачи: перевести число из одной системы счисления в другую, например, из двоичной в десятичную)
 ----
  Пример: 55
  Ответ: LV
 ----
  Пример2: 234
  Ответ: CCXXXIV
  
  
   2) Нам дана произвольная строка, содержащая числа, символы и спец. символы. Необходимо найти самое большое число в строке.
 (вариации задачи: найти сумму всех чисел в строке, найти самое длинное слово в строке)
 
 Пример: "fRETre3445 43#$% 32424234 #$@$34 2323"
 Ответ: 32424234
 
 
  3) Нам дан список слов, нужно найти слова, которые имеют анаграммы в данном списке. Анаграмма - это когда переставив буквы в одном слове мы можем получить второе слово (cat > act)
 
 Пример: ['cat', 'act', 'dinosaur', 'testing', 'setting']
 Ответ: ['cat', 'act', 'testing', 'setting']
