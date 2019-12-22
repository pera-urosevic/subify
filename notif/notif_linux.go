package notif

import (
	"fmt"
	"os/exec"
)

// SendSubtitleDownloadSuccess sends a notification when download went well
func SendSubtitleDownloadSuccess(successAPI string) {
	_ = Info("I found a subtitle for your video :)", fmt.Sprintf("Thank you %s <3", successAPI))
}

// SendSubtitleCouldNotBeDownloaded sends a notification when download went bad
func SendSubtitleCouldNotBeDownloaded(noSucessAPIs string) {
	_ = Error("!! I didn't found any subtitle :'(", fmt.Sprintf("No match for your video in : %s. Try later !", noSucessAPIs))
}

func sendMessage(title, message string) error {
	subifyIcon := downloadIcon()
	return exec.Command("notify-send", "-i", subifyIcon, fmt.Sprintf("Subify - %s", title), message).Run()
}

// Error send a notification error
func Error(title, message string) error {
	return sendMessage(title, message)
}

// Info send a notification information
func Info(title, message string) error {
	return sendMessage(title, message)
}
