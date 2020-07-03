### ginicoef

Ginicoef is a library used to calculate gini coefficients.


### Installation
```
go get github.com/vearne/ginicoef
```


###  Example
```
package main

import (
	"fmt"
	"github.com/vearne/ginicoef"
)

func main() {
	incoms := []float64{1, 2, 3, 4}
	data := ginicoef.GiniCoefficient(incoms)
	fmt.Println(data)
}

```


### wiki
1. [Gini coefficient](https://en.wikipedia.org/wiki/Gini_coefficient)   
2. [Trapezoidal rule](https://en.wikipedia.org/wiki/Trapezoidal_rule)
