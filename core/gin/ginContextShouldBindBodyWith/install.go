package ginContextShouldBindBodyWith

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/gin-gonic/gin"
	"go-agent/model"
)

func init() {
	model.HookMap["ginContextShouldBindBodyWith"] = new(GinContextShouldBindBodyWith)
}

type GinContextShouldBindBodyWith struct {
}

func (h *GinContextShouldBindBodyWith) Hook() {
	context := &gin.Context{}
	err := gohook.HookMethod(context, "ShouldBindBodyWith", ShouldBindBodyWith, ShouldBindBodyWithT)
	if err != nil {
		fmt.Println(err, "GinContextShouldBindBodyWith")
	} else {
		fmt.Println("GinContextShouldBindBodyWith")
	}
}

func (h *GinContextShouldBindBodyWith) UnHook() {
	context := &gin.Context{}
	gohook.UnHookMethod(context, "ShouldBindBodyWith")
}
