package common

import (
	"testing"
)

func TestCommand(t *testing.T) {
	ctrl := &Control{}
	lightOnCommand := new(LightOnCommand)
	lightOffCommand := new(LightOffCommand)
	videoOnCommand := new(VideoOnCommand)
	videoOffCommand := new(VideoOffCommand)

	// 开灯
	ctrl.SetCommand(lightOnCommand)
	ctrl.Execute()
	// 关灯
	ctrl.SetCommand(lightOffCommand)
	ctrl.Execute()
	// 开电视
	ctrl.SetCommand(videoOnCommand)
	ctrl.Execute()
	// 关电视
	ctrl.SetCommand(videoOffCommand)
	ctrl.Execute()
}
