package qrcode_bundle

import (
	qrcode_encode "github.com/skip2/go-qrcode"
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
func Encode(content string, recoveryLevel int, size int) ([]byte, error) {
	recoveryLevelForEncode := convertRecoveryLevel(recoveryLevel)
	qrcode, err := qrcode_encode.Encode(content, recoveryLevelForEncode, size)
	return qrcode, err
}
