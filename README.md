# Get and convert currency from cbr.ru

### Features

  - Get all currency from cbr.ru
  - Convert from one currency to another
  - Get the rate of a currency for one unit in rubles
  
  
### Installation

```sh
$ go get github.com/k0v4back/cbr-valute/src 
```

### Description of methods
  - src.Convert("USD", "RUB", 1)
    - USD - currency from which from convert
    - RUB - currency to convert
    - 1   - amount of first currency
  - src.GetValuteByName("USD")
    - USD - get the price of one dollar in rubles


### List available currency

```sh
AUD
AZN
GBP
AMD
BYN
BGN
BRL
... AND MANY OTHER (https://www.cbr.ru/scripts/XML_daily_eng.asp)
```
[All list](https://www.cbr.ru/scripts/XML_daily_eng.asp)


#
### Example
```sh
package main

import (
	"fmt"

	"github.com/k0v4back/cbr-valute/src"
)

func main() {
	fmt.Println(src.Convert("USD", "RUB", 1))
	fmt.Println(src.GetValuteByName("USD"))
}
```


The MIT License (MIT)

Copyright (c) 2019-present Vadim Kosolapov

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
