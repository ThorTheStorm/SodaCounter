package logging

import "fmt"

type (
	LogType string
)

const (
	InfoLog  = LogType("INFO")
	WarnLog  = LogType("WARN")
	ErrorLog = LogType("ERROR")
	DebugLog = LogType("DEBUG")
)

func AppLog(logType LogType, message any) error {
	// Create the logMessage based on logType
	switch logType {
	case InfoLog:
		fmt.Println("[INFO]: " + fmt.Sprint(message))
		return nil
	case WarnLog:
		fmt.Println("[WARN]: " + fmt.Sprint(message))
		return nil
	case ErrorLog:
		fmt.Println("[ERROR]: " + fmt.Sprint(message))
		return nil
	case DebugLog:
		fmt.Println("[DEBUG]: " + fmt.Sprint(message))
		return nil
	default:
		return fmt.Errorf("invalid log type: %s", logType)
	}
}
