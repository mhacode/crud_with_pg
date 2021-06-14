package main

import "fmt"

func main() {
	res := divide(10, 20)
	fmt.Println(res)

}

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0.0, "Can't Divide by Zero"
	}
	return a / b, ""
}

func formatTimeWithMessage(hours, minutes int) (string, error) {
	formatted, err := formatTime(hours, minutes)
	if err != nil {
		return "", fmt.Errorf("Format time with error message: %v", err)
	}
	return "It is " + formatted + " o'clock", nil
}
