package crypter

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

// CLI parses command line arguments and executes app with provided parameters, and returns an exit code
func CLI(args []string) int {
	encodeFlags := flag.NewFlagSet("encode", flag.ExitOnError)
	encodeFormat := encodeFlags.String("f", "base64", "format to encode")
	encodeOutput := encodeFlags.String("o", "stdout", "write encoded output to stdout or file")

	decodeFlags := flag.NewFlagSet("decode", flag.ExitOnError)
	decodeFormat := decodeFlags.String("f", "base64", "format to decode from")
	decodeOutput := decodeFlags.String("o", "stdout", "write decoded output to stdout or file")

	var ac AppContext = AppContext{}
	var encodeMap map[string]func([]byte) []byte
	encodeMap["base64"] = toBase64
	encodeMap["hex"] = toHex

	var decodeMap map[string]func(string) ([]byte, error)
	decodeMap["base64"] = fromBase64String
	decodeMap["hex"] = fromHexString

	switch args[1] {
	case "encode":
		encodeFlags.Parse(args[2:])
		// TODO

	}

	return 0
}

func Run(ac AppContext) error {
	// TODO
	return nil
}

type AppContext struct {
	fun    interface{}
	output io.Writer
}

func (app *AppContext) fromArgs(fun interface{}, output io.Writer) *AppContext {
	app.fun = fun
	app.output = output

	return app
}

func (app *AppContext) Run() error {
	// TODO
	return nil
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

func fromBase64String(src string) ([]byte, error) {
	decodedData := make([]byte, base64.StdEncoding.DecodedLen(len(src)))

	decodedData, err := fromBase64([]byte(src))
	if err != nil {
		return nil, err
	}
	return decodedData, nil

}

func toBase64String(src []byte) string {
	encodedString := string(toBase64(src))
	return encodedString
}

func toHex(src []byte) []byte {
	var encodedData []byte = make([]byte, hex.EncodedLen(len(src)))

	n := hex.Encode(encodedData, src)

	fmt.Fprintf(os.Stdout, "Encoded %v bytes to hex", n)

	return encodedData
}

func fromHexString(src string) ([]byte, error) {
	var decodedData []byte = make([]byte, 0)

	decodedData, err := hex.DecodeString(src)

	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
