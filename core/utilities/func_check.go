package utilities

import "time"

func SafeString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func SafeInt(ptr *int) int {
	if ptr != nil {
		return *ptr
	}
	return 0
}

func SafeTime(ptr *time.Time) time.Time {
	if ptr != nil {
		return *ptr
	}
	return time.Time{}
}

func SafeFloat32(ptr *float32) float32 {
	if ptr != nil {
		return *ptr
	}
	return 0.0
}

func SafeFloat64(ptr *float64) float64 {
	if ptr != nil {
		return *ptr
	}
	return 0.0
}

func SafeBool(ptr *bool) bool {
	if ptr != nil {
		return *ptr
	}
	return false
}

func SafeUint(ptr *uint) uint {
	if ptr != nil {
		return *ptr
	}
	return 0
}

//for interface
func SafeInterfaceSlice(ptr *[]interface{}) []interface{} {
	if ptr != nil {
		return *ptr
	}
	return []interface{}{}
}

func SafeInterface(ptr *interface{}) interface{} {
	if ptr != nil {
		return *ptr
	}
	return nil
}