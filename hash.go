package main

import (
    "fmt"
    "os"
    "hash/fnv"
    "encoding/binary"
    "bytes"
)

func hash32(s string) []byte{
    h := fnv.New32a()
    h.Write([]byte(s)) 
    buf := new(bytes.Buffer)
    err := binary.Write(buf, binary.LittleEndian, h.Sum32())
    if err != nil {
        fmt.Println("binary.Write failed:", err)
    }
    return buf.Bytes() 
}

func hashDirTo4Level(dir string) string {
    h := hash32(dir);
    var buffer [7]byte  
    for i := 0; i < 7; i++ {
        if (i % 2 != 0) {
            buffer[i] = '/'
        } else {
            buffer[i] = h[i / 2] % 16 + 'A'
        }
    }
    return string(buffer[:])
}

func main() {
    fmt.Printf("%s\n", hashDirTo4Level(os.Args[1]))
}
