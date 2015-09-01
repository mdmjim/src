package Rsystem

import ( 
    //"io" 
	"bytes"
    "github.com/qiniu/iconv"
	"errors"
)

type Encoder struct{
	
}

func (enc *Encoder) DecodeToUtf8(s []byte,orig_charset string) ([]byte, error) {
    cd, err := iconv.Open("utf-8", orig_charset) // convert gbk to utf8
    if err != nil { 
        return nil ,errors.New("iconv.Open failed!")
    }
    defer cd.Close()

    input := bytes.NewReader(s)   // eg. input := os.Stdin || input, err := os.Open(file)
    bufSize := 0 // default if zero
    r := iconv.NewReader(cd, input, bufSize)
    buff :=make([]byte,len(s))
	_,err = r.Read(buff)
	return buff,err
} 
 