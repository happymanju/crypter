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
	encodeFormat := encodeFlags.String("f", "base64", "format to encode: base64 or md5")
	encodeOutput := encodeFlags.String("o", "stdout", "write encoded output to stdout or file")

	decodeFlags := flag.NewFlagSet("decode", flag.ExitOnError)
	decodeFormat := decodeFlags.String("f", "base64", "format to decode from")
	decodeOutput := decodeFlags.String("o", "stdout", "write decoded output to stdout or file")

	return 0
}

type AppContext struct {
	mode   string
	decode bool
	output io.Writer
}

func (app *AppContext) fromArgs(mode string, decode bool, output io.Writer) *AppContext {
	app.mode = mode
	app.decode = decode
	app.output = output

	return app
}

func (app *AppContext) Run() error {
	// TODO
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

func fromBase64String(src string) (string, error) {
	decodedData := make([]byte, base64.StdEncoding.DecodedLen(len(src)))

	decodedData, err := fromBase64([]byte(src))
	if err != nil {
		return "", err
	}
	return string(decodedData), nil

}

func toBase64String(src []byte) string {
	encodedString := string(toBase64(src))
	return encodedString
}

func toHexString(src []byte) string {
	var encodedData []byte = make([]byte, hex.EncodedLen(len(src)))

	n := hex.Encode(encodedData, src)

	fmt.Fprintf(os.Stdout, "Encoded %v bytes to hex", n)

	return string(encodedData)
}

func fromHexString(src string) ([]byte, error) {
	var decodedData []byte = make([]byte, 0)

	decodedData, err := hex.DecodeString(src)

	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
