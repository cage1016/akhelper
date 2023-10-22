package akhelper

import (
	"fmt"
	"os/exec"
	"strings"
)

const notifier = `tell application id "com.runningwithcrayons.Alfred" to run trigger "%s" in workflow "%s" with argument "%s"`

type Notifier struct {
	Title      string
	TriggerID  string
	WorkflowID string
}

func (an *Notifier) Notify(msg string) {
	exec.Command("osascript", "-e", fmt.Sprintf(notifier, an.TriggerID, an.WorkflowID, strings.Join([]string{an.Title, msg}, ","))).Run()
}

func NewNotifier(title, triggerID, workflowID string) *Notifier {
	return &Notifier{
		Title:      title,
		TriggerID:  triggerID,
		WorkflowID: workflowID,
	}
}
