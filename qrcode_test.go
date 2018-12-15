package qrcode_test

import (
	"io/ioutil"
	. "qrcode_bundle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Decode_Input_QrCode_Should_Be_Hello_QrCode(t *testing.T) {
	expectedContent := "https://google.com"
	qrcode, _ := Encode("https://google.com")

	actualContent, _ := Decode(qrcode)

	assert.Equal(t, len(expectedContent), len(actualContent))
	assert.Equal(t, expectedContent, actualContent)
}

func Test_Encode_Input_Text_Hello_QrCode_Should_Be_QrCode(t *testing.T) {
	expectedQrCode, _ := ioutil.ReadFile("resources/qrcode.png")

	actualQrCode, _ := Encode("https://google.com")

	assert.Equal(t, len(expectedQrCode), len(actualQrCode))
	assert.Equal(t, expectedQrCode, actualQrCode)
}
