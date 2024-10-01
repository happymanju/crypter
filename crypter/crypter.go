package crypter

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type decodeFunc func([]byte) ([]byte, error)
type encodeFunc func([]byte) ([]byte)
type modeFuncs struct {
	encodeFunc
	decodeFunc
}

var modesMap map[string]modeFuncs = map[string]modeFuncs {
	"base64": {toBase64,fromBase64},
	"hex": {toHex, fromHex},
}



// CLI parses command line arguments and executes app with provided parameters, and returns an exit code
func CLI(args []string) int {
	encodeFlags := flag.NewFlagSet("encode", flag.ExitOnError)
	encodeFormat := encodeFlags.String("f", "base64", "format to encode")
	encodeOutput := encodeFlags.String("o", "", "write encoded output to file")
	encodeInputFile := encodeFlags.String("i", "", "read contents of file for encode")

	decodeFlags := flag.NewFlagSet("decode", flag.ExitOnError)
	decodeFormat := decodeFlags.String("f", "base64", "format to decode from")
	decodeOutput := decodeFlags.String("o", "stdout", "write decoded output to file")
	decodeInputFile := decodeFlags.String("i", "", "read contents of file for decode")



	switch args[1] {
	case "encode":
		encodeFlags.Parse(args[2:])

		// TODO: figure out to switch outputs for appContext
		ac := &appContext{}
		ac.fromArgs(modesMap[*encodeFormat].encodeFunc,nil,os.Stdout,)

		err := RunApp(ac, encodeFlags.Args())
		if err != nil {
			log.Printf("Error running app: %v\n", err)
			return 1
		}
	
	case "decode":
		decodeFlags.Parse(args[2:])

		ac.fun = modesMap[*decodeFormat].decodeFunc
		if *decodeOutput == "stdout" {
			ac.output = os.Stdout
		} else {
			f, err := os.Open(*decodeOutput)
			if err != nil {
				log.Printf("Error with -o flag: %v\n", err)
				return 1
			}
			ac.output = f
		}

		err := RunApp(ac, decodeFlags.Args())
		if err != nil {
			log.Printf("Error running app: %v\n", err)
			return 1
		}
	default:
		fmt.Printf("invalid flags")
		encodeFlags.Usage()

	}

	return 0
}

// RunApp - use parsed CLI options to execute app
func RunApp(ac appContext, tailArgs []string) error {
	// TODO
	input := tailArgs[0]

	return nil
}

type appContext struct {
	enc encodeFunc
	dec decodeFunc
	input io.Reader
	output io.Writer
	payload []byte
}

func (app *appContext) fromArgs(enc encodeFunc, dec decodeFunc, input string, output io.Writer, payload []byte) *appContext {
	app.enc = enc
	app.dec = dec
	app.input = input
	app.output = output
	app.payload = payload

	return app
}

func toBase64(src []byte) []byte {
	var encodedData []byte = make([]byte, base64.StdEncoding.EncodedLen(len(src)))

	base64.StdEncoding.Encode(encodedData, src)
	return encodedData
}

func fromBase64(src []byte) ([]byte, error) {
	var decodedData []byte = make([]byte, 0)

	n, err := base64.StdEncoding.Decode(decodedData, src)
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(os.Stdout, "Decoded %v bytes from base64", n)

	return decodedData, nil
}

func toHex(src []byte) []byte {
	var encodedData []byte = make([]byte, hex.EncodedLen(len(src)))

	n := hex.Encode(encodedData, src)

	fmt.Fprintf(os.Stdout, "Encoded %v bytes to hex\n", n)

	return encodedData
}

func fromHex(src []byte) ([]byte, error) {
	var decodedData []byte = make([]byte, hex.DecodedLen(len(src)))

	n, err := hex.Decode(decodedData,src)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stdout, "Decoded %v bytes\n", n)
	return decodedData, nil

}