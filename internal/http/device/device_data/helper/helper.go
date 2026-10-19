package helper

import (
	"errors"
	"time"
)

// ============================================================
// BUILD SLOT START
// ============================================================

func BuildSlotStart(
	baseDate time.Time,
	shiftStart time.Time,
	shiftEnd time.Time,
	slotStart time.Time,
	loc *time.Location,
) time.Time {

	result := time.Date(
		baseDate.Year(),
		baseDate.Month(),
		baseDate.Day(),

		slotStart.Hour(),
		slotStart.Minute(),
		slotStart.Second(),
		0,

		loc,
	)

	// Overnight shift.
	//
	// Example:
	//
	// Shift B = 20:00 → 08:00
	//
	// Slot = 00:00 → 01:00
	//
	// The slot belongs to the next calendar day.

	if isOvernightShift(
		shiftStart,
		shiftEnd,
	) {

		if slotStart.Hour() < shiftStart.Hour() {

			result = result.AddDate(
				0,
				0,
				1,
			)
		}
	}

	return result
}

// ============================================================
// BUILD SLOT END
// ============================================================

func BuildSlotEnd(
	start time.Time,
	slotEnd time.Time,
	loc *time.Location,
) time.Time {

	result := time.Date(
		start.Year(),
		start.Month(),
		start.Day(),

		slotEnd.Hour(),
		slotEnd.Minute(),
		slotEnd.Second(),
		0,

		loc,
	)

	// Example:
	//
	// 23:00 → 00:00
	//
	// End must be next day.

	if !result.After(start) {

		result = result.AddDate(
			0,
			0,
			1,
		)
	}

	return result
}

// ============================================================
// CHECK OVERNIGHT SHIFT
// ============================================================

func isOvernightShift(
	shiftStart time.Time,
	shiftEnd time.Time,
) bool {

	startMinutes :=
		shiftStart.Hour()*60 +
			shiftStart.Minute()

	endMinutes :=
		shiftEnd.Hour()*60 +
			shiftEnd.Minute()

	return startMinutes > endMinutes
}

var ErrInvalidRequest = errors.New("invalid production request")
