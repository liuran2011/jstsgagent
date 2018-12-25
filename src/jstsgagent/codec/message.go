package codec

type Message struct {
    Magic uint32 
    TotalLength uint32
    Blocks *map[uint16]TLV
}
