# Uniq-Util
Утилита для поиска уникальных строк.

# Параметры.

`-c` - подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.

`-d` - вывести только те строки, которые повторились во входных данных.

`-u` - вывести только те строки, которые не повторились во входных данных.

`-f num_fields` - не учитывать первые <num_fields> полей в строке. Полем в строке является непустой набор символов отделенный пробелом.

`-s num_chars` - не учитывать первые <num_chars> символов в строке. При использовании вместе с параметром <-f> учитываются первые символы после <num_fields> полей (не учитывая пробел-разделитель после последнего поля).

`-i` - не учитывать регистр букв.

# Использование.

`uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]`

1. Все параметры опциональны. Поведение утилиты без параметров - простой вывод уникальных строк из входных данных.

2. Параметры `-c, -d, -u` взаимозаменяемы. Необходимо учитывать, что параллельно эти параметры не имеют никакого смысла. При передаче одного вместе с другим отображается пользователю правильное использование утилиты.

3. Если не передан `input_file`, то входным потоком считается stdin

4. Если не передан `output_file`, то выходным потоком считается stdout.
