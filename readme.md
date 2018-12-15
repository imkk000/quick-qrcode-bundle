# Quick Qrcode Encode & Decode

## Golang require package
- Encode: [https://github.com/skip2/go-qrcode](https://github.com/skip2/go-qrcode)
- Decode: [https://github.com/tuotoo/qrcode](https://github.com/tuotoo/qrcode)

## Example
```go
package main

import (
    "fmt"
    
    // import qrcode bundle library
    "github.com/ImKK-000/quick-qrcode-bundle"
)

func main() {
    // qrcodeEncode -> return []byte
    qrcodeEncode, _ := qrcode.Encode("https://google.com")
    // qrcodeDecode -> return string
    qrcodeDecode, _ := qrcode.Decode(qrcodeEncode)
    // console display -> https://google.com
    fmt.Println("qrcode content", qrcodeDecode)
}
```