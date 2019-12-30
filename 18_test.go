package design

import "testing"

func Test_18_1(t *testing.T) {
	tv := new(TV)
	openCommand := new(TvOpenCommand)
	openCommand.Init(tv)
	closeCommand := new(TvCloseCommand)
	closeCommand.Init(tv)
	playCommand := new(TvPlayCommand)
	playCommand.Init(tv)
	tvInvoker := new(TvInvoker)
	tvInvoker.Init(openCommand,closeCommand,playCommand)
	tvInvoker.Open()
	tvInvoker.Play()
	tvInvoker.Close()
}
