package codec 

import (
    "io"
)

type Codec interface {
    Encode(w io.Writer) error 
}
