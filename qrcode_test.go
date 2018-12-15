package qrcode_bundle_test

import (
	"io/ioutil"
	. "qrcode_bundle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encode_Input_Text_Hello_QrCode_Should_Be_QrCode(t *testing.T) {
	expectedQrCode, _ := ioutil.ReadFile("resources/qrcode.png")

	actualQrCode, _ := Encode("Hello QrCode", RecoveryLevelHighest, 256)

	assert.Equal(t, len(expectedQrCode), len(actualQrCode))
	assert.Equal(t, expectedQrCode, actualQrCode)
}
