# Методы алгебраической геометрии в криптографии

## Вход:
1. Тип:
  * P (char != 2; char != 3)
  * 2S (char = 2; суперсингулярная)
  * 2N (char = 2; несуперсингулярная)
2. Поле:
  * При P строка всегда равна `1`
  * При 2S и 2N - неприводимый многочлен над полем Z2 степени n, порождающий для GF(2^n). Вид многочлена `x^n + ... + x + 1`
3. Кривая:
  * При P - `a b`
  * При 2S и 2N - `a b c` - битовые строки разложения по базису α^(n - 2), ..., α^2, α, 1 (α - корень неприводимого многочлена)

4. Список задач:
  * Читать пока не кончится stdin
    * У, точка, число - умножить точку на число
    * С, точка, точка - сложить точку с точкой

## Выход:
```
точка * число = точка
...
точка + точка = точка
```

## Число (num) имеет вид
  * Число в указанной системе счислений (с помощью ключа -base)


## Точка имеет вид
* Для P точка имеет вид `x y`, где x и y некоторые числа в указанной системе счисления
* Для 2N/2S точка имеет вид `x y`, где x и y - битовые строки, представляющие x и y разложением по базису α^(n - 1), α^(n - 2), ..., α^2, α, 1 

## Примеры использования
``` bash
$ ./magvk_linux_amd64 -base 16 -i input_file -o output_file
```
