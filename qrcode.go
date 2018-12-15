package qrcode_bundle

import (
	"bytes"

	qrcode_encode "github.com/skip2/go-qrcode"
	qrcode_decode "github.com/tuotoo/qrcode"
)

const (
	// RecoveryLevelLow - 7% error recovery.
	RecoveryLevelLow int = iota

	// RecoveryLevelMedium - 15% error recovery.
	RecoveryLevelMedium

	// RecoveryLevelHigh - 25% error recovery.
	RecoveryLevelHigh

	// RecoveryLevelHighest - 30% error recovery.
	RecoveryLevelHighest

	Size = 256
)

func convertRecoveryLevel(recoveryLevelIndex int) qrcode_encode.RecoveryLevel {
	qrcodeRecoveryLevel := []qrcode_encode.RecoveryLevel{
		qrcode_encode.Low,
		qrcode_encode.Medium,
		qrcode_encode.High,
		qrcode_encode.Highest,
	}
	if len(qrcodeRecoveryLevel) > recoveryLevelIndex {
		recoveryLevelIndex = RecoveryLevelHighest
	}
	return qrcodeRecoveryLevel[recoveryLevelIndex]
}

// Encode - qrcode encoder
func Encode(content string) ([]byte, error) {
	recoveryLevelForEncode := convertRecoveryLevel(RecoveryLevelMedium)
	qrcode, err := qrcode_encode.Encode(content, recoveryLevelForEncode, Size)
	return qrcode, err
}

// Decode - qrcode decoder
func Decode(qrcode []byte) (string, error) {
	qrcode_decode.SetDebug(false)
	buffer := bytes.NewReader(qrcode)
	matrix, err := qrcode_decode.Decode(buffer)
	return matrix.Content, err
}
