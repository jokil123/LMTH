package main

import (
	"flag"
	c "lmth/convert"
	"log"
)

var (
	encode = flag.Bool("encode", false, "encode")
	decode = flag.Bool("decode", false, "decode")
	i      = flag.String("i", "", "input file")
	o      = flag.String("o", "", "output file")
)

func init() {
	flag.Parse()
}

func main() {
	if *encode && *decode || !*encode && !*decode {
		log.Println("Please specify either encode or decode")
		return
	}

	args := flag.Args()

	if len(args) != 2 {
		log.Println("Please specify input and output files")
		return
	}

	i := args[0]
	o := args[1]

	log.Printf("Encode: %t, Decode: %t, i: %s, o: %s", *encode, *decode, i, o)

	if *encode {
		c.EncodeFile(i, o)
	} else if *decode {
		c.DecodeFile(i, o)
	}
}
