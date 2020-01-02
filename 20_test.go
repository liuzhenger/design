package design

import (
	"fmt"
	"testing"
)

func Test_20_1(t *testing.T) {
	c := ControlCtx{
		StateMap: map[int]IState{},
	}
	c.AddState(PowerOff, &PowerOffState{BaseState{Ctx: &c}})
	c.AddState(Play, &PlayState{BaseState{Ctx: &c}})
	c.AddState(StandBy, &StandByState{BaseState{Ctx: &c}})
	c.SetState(PowerOff)
	c.PowerOn()
	c.Play()
	c.Standby()
	c.PowerOff()
	fmt.Println("------------------")
	c.Play()
	fmt.Println("------------------")
	c.PowerOn()
	c.Play()
	c.PowerOn()
}
