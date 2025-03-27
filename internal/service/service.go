package service

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
	"time"
)

func CheckServiceStatus(serviceName string) (string, error) {
	m, err := mgr.Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to manager service: %v", err)

	}
	defer func(m *mgr.Mgr) {
		err := m.Disconnect()
		if err != nil {

		}
	}(m)

	s, err := m.OpenService(serviceName)
	if err != nil {
		return "", fmt.Errorf("error getting service: %v", err)
	}
	defer func(s *mgr.Service) {
		err := s.Close()
		if err != nil {

		}
	}(s)

	status, err := s.Query()
	if err != nil {
		return "", fmt.Errorf("could not query service %s: %v", serviceName, err)
	}

	var statusStr string
	switch status.State {
	case windows.SERVICE_STOPPED:
		statusStr = "stopped"
	case windows.SERVICE_START_PENDING:
		statusStr = "start pending"
	case windows.SERVICE_STOP_PENDING:
		statusStr = "stop pending"
	case windows.SERVICE_RUNNING:
		statusStr = "running"
	case windows.SERVICE_CONTINUE_PENDING:
		statusStr = "continue pending"
	case windows.SERVICE_PAUSE_PENDING:
		statusStr = "pause pending"
	case windows.SERVICE_PAUSED:
		statusStr = "paused"
	default:
		statusStr = "unknown"
	}

	return statusStr, nil

}

func MonitorService(serviceName, expectedStatus string, interval int, callback func(message string)) {
	for {
		status, err := CheckServiceStatus(serviceName)
		if err != nil {
			callback(fmt.Sprintf("Error checking service %s: %v", serviceName, err))
		} else if status != expectedStatus {
			callback(fmt.Sprintf("Alert: Service %s is %s, expected: %s", serviceName, status, expectedStatus))
		} else {
			callback(fmt.Sprintf("Service %s is %s, expected: %s", serviceName, status, expectedStatus))
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
