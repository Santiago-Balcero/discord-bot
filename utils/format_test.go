package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatIntegerOne(t *testing.T) {
	result := FormatInteger(1)
	assert.Equal(
		t,
		"1",
		result,
		"Should return integer as string if integer less than 1.000",
	)
}

func TestFormatIntegerThousands(t *testing.T) {
	result := FormatInteger(9999)
	assert.Equal(t,
		"9.999",
		result,
		"Should format integers greater than 999",
	)
}

func TestFormatIntegerMillions(t *testing.T) {
	result := FormatInteger(111222333)
	assert.Equal(t,
		"111.222.333",
		result,
		"Should format integers greater than 999.999",
	)
}

func TestFormatIntegerBillions(t *testing.T) {
	result := FormatInteger(111222333444)
	assert.Equal(t,
		"111.222.333.444",
		result,
		"Should format integers greater than 999.999.999",
	)
}

func TestMillisecondsToTimeMilliseconds(t *testing.T) {
	result := MillisecondsToTime(100)
	assert.Equal(
		t,
		"00:00:00",
		result,
		"Should format 100 ms as 00:00:00",
	)
}

func TestMillisecondsToTimeOneSecond(t *testing.T) {
	result := MillisecondsToTime(900)
	assert.Equal(
		t,
		"00:00:01",
		result,
		"Should format 900 ms as 00:00:01",
	)
}

func TestMillisecondsToTimeSeconds(t *testing.T) {
	result := MillisecondsToTime(10000)
	assert.Equal(
		t,
		"00:00:10",
		result,
		"Should format 1000 ms as 00:00:10",
	)
}

func TestMillisecondsToTimeMinutes(t *testing.T) {
	result := MillisecondsToTime(120000)
	assert.Equal(
		t,
		"00:02:00",
		result,
		"Should format 120000 ms as 00:02:00",
	)
}

func TestMillisecondsToTimeHours(t *testing.T) {
	result := MillisecondsToTime(7200000)
	assert.Equal(
		t,
		"02:00:00",
		result,
		"Should format 7200000 ms as 02:00:00",
	)
}

func TestMillisecondsToTimeDays(t *testing.T) {
	result := MillisecondsToTime(90000000)
	assert.Equal(
		t,
		"01:00:00",
		result,
		"Should format 90000000 ms as 01:00:00",
	)
}
