package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch nType := n.(type) {
	case directMessage:
		return nType.senderUsername, nType.importance()
	case groupMessage:
		return nType.groupName, nType.importance()
	case systemAlert:
		return nType.alertCode, nType.importance()
	default:
		return "", 0
	}
}

func main() {
	newUrgentDM := directMessage{
		senderUsername: "Tharani",
		messageContent: "Vanakam di Mapla",
		priorityLevel:  25,
		isUrgent:       true,
	}
	newDM := directMessage{
		senderUsername: "Tharani",
		messageContent: "Calm down",
		priorityLevel:  10,
		isUrgent:       false,
	}
	name, importanceVal := processNotification(newUrgentDM)
	fmt.Println("Name: ", name, " importance: ", importanceVal)

	name, importanceVal = processNotification(newDM)
	fmt.Println("Name: ", name, " importance: ", importanceVal)

	newGM := groupMessage{
		groupName:      "Vanma Vanam",
		messageContent: "Angutu poda",
		priorityLevel:  5,
	}
	name, importanceVal = processNotification(newGM)
	fmt.Println("Name: ", name, " importance: ", importanceVal)

	newSA := systemAlert{
		alertCode:      "FUCKED",
		messageContent: "Nah! We are fucked",
	}
	name, importanceVal = processNotification(newSA)
	fmt.Println("Name: ", name, " importance: ", importanceVal)
}
