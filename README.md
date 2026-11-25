# facial-geometry-go

Zero-dependency Go package for computing 2D facial landmark geometry metrics (canthal tilt, facial thirds balance, golden ratio delta, and bilateral mirror symmetry).

## Install

```bash
go get github.com/visualcraft-lab/facial-geometry-go
```

## Example

```go
package main

import (
    "fmt"
    fg "github.com/visualcraft-lab/facial-geometry-go"
)

func main() {
    tilt, _ := fg.CanthalTiltDeg(fg.Point2D{X: 100, Y: 200}, fg.Point2D{X: 140, Y: 195})
    fmt.Printf("Canthal tilt: %.2f deg\n", tilt)
}
```

## Reference & Online Calculator

Interactive visual verification and landmark definitions:
- [Looksmax Scale](https://looksmaxscale.com/)
- [Canthal Tilt Measurement Guide](https://looksmaxscale.com/canthal-tilt)
- [Facial Proportion Reference](https://looksmaxscale.com/attractiveness-test)

## License

MIT
