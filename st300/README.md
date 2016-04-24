# ST300 Parser

## Model ST300, ST340 and ST340LC

Report | Supported
 --- | ---
Status Report | yes

Zip aren't supported yet.

## Usage

To parse one STT from a string, you can do something like:

```golang
package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/larixsource/suntech/st300"
)

func main() {
	data := "ST300STT;100850000;01;010;20081017;07:41:56;00100;+37.478519;+126.886819;000.012;000.00;9;1;0;15.30;001100;1;0072;0;4.5;1\r"
	opts := st300.ParserOpts{}
	p := st300.ParseString(data, opts)
	for p.Next() {
		spew.Dump(p.Msg())
	}
	if p.Error() != nil {
		log.Printf("parsing error: %s", p.Error())
	}
}
```

Also, there a ParseBytes function, that accepts a byte slice.

However, the parser is designed to operate over a stream, extracting frames in a loop from a reader (a file, socket, etc):

```golang
reader := ... // stream of frames
opts := st300.ParserOpts{
    SkipUnknownFrames: true,
}
p := st300.ParseString(reader, opts)
for p.Next() {
    spew.Dump(p.Msg())
}
if p.Error() != nil {
    log.Printf("parsing error: %s", p.Error())
}
```
