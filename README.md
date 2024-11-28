# Origin / Credits
This is the Go Version of the [CS2 Inspect Generator](https://github.com/dr3fty/cs2-inspect-gen) from [dr3fty](https://github.com/dr3fty)

# Usage
- go install github.com/lian-yang/cs2-inspect-gen@latest
- cs2-inspect-gen -h

```bash
Usage of cs2-inspect-gen:
  -cn string
        customname
  -di uint
        defindex (default 1)
  -pi uint
        paintindex
  -ps uint
        paintseed
  -pw float
        paintwear
  -r uint
        rarity
  -si1 uint
        sticker 1 id
  -si2 uint
        sticker 2 id
  -si3 uint
        sticker 3 id
  -si4 uint
        sticker 4 id
  -si5 uint
        sticker 5 id
  -st uint
        stattrak (default 1)
  -stc uint
        stattrak count (default 123)
  -sw1 float
        sticker 1 wear
  -sw2 float
        sticker 2 wear
  -sw3 float
        sticker 3 wear
  -sw4 float
        sticker 4 wear
  -sw5 float
        sticker 5 wear

```

# example

Used as a dependency

go get github.com/lian-yang/cs2-inspect-gen@latest

```go
package main

import (
	"fmt"

	"github.com/lian-yang/cs2-inspect-gen"
	"github.com/lian-yang/cs2-inspect-gen/econ_pb2"
)

func main() {
	inspect, err := cs2_inspect_gen.GenerateInspect(&econ_pb2.CEconItemPreviewDataBlock{
		Defindex:   cs2_inspect_gen.Uint32Ptr(0),
		Paintindex: cs2_inspect_gen.Uint32Ptr(0),
		Rarity:     cs2_inspect_gen.Uint32Ptr(0),
		Paintwear:  cs2_inspect_gen.Uint32Ptr(0),
		Paintseed:  cs2_inspect_gen.Uint32Ptr(0),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(inspect) // 输出 001800200028003800400055AA61AA
}

```