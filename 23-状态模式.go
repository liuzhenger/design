package design

import "fmt"

// 对象的行为跟随状态的改变而改变

// 状态标识
const (
	PowerOff = 1 // 关机状态
	Play     = 2 // 播放状态
	StandBy  = 3 // 待机状态
)

// 电视机接口(行为接口)
type IState interface {
	PowerOn()
	PowerOff()
	Play()
	Standby()
}

// 上下文角色
type ControlCtx struct {
	StateMap map[int]IState
	State    IState
}

func (c *ControlCtx) AddState(key int, s IState) {
	c.StateMap[key] = s
}

func (c *ControlCtx) RemoveState(key int) {
	delete(c.StateMap, key)
}

func (c *ControlCtx) SetState(s int) {
	var ok bool
	if c.State, ok = c.StateMap[s]; !ok {
		panic(fmt.Sprint("not found state: ", s))
	}
}

func (c *ControlCtx) PowerOn() {
	c.State.PowerOn()
}

func (c *ControlCtx) PowerOff() {
	c.State.PowerOff()
}

func (c *ControlCtx) Play() {
	c.State.Play()
}

func (c *ControlCtx) Standby() {
	c.State.Standby()
}

// 状态基类
type BaseState struct {
	Ctx *ControlCtx
}

func (b *BaseState) PowerOn() {
}

func (b *BaseState) PowerOff() {
}

func (b *BaseState) Play() {
}

func (b *BaseState) Standby() {
}

// 待机状态
type StandByState struct {
	BaseState
}

func (s *StandByState) PowerOff() {
	fmt.Println("关机...")
	s.Ctx.SetState(PowerOff)
	s.Ctx.State.PowerOff()
}

func (s *StandByState) Play() {
	fmt.Println("播放...")
	s.Ctx.SetState(Play)
	s.Ctx.State.Play()
}

func (s *StandByState) Standby() {
	fmt.Println("do...待机")
}

// 关机状态
type PowerOffState struct {
	BaseState
}

func (p *PowerOffState) PowerOn() {
	fmt.Println("开机...")
	p.Ctx.SetState(StandBy)
	p.Ctx.State.Standby()
}

func (p *PowerOffState) PowerOff() {
	fmt.Println("do...关机")
}

// 播放状态
type PlayState struct {
	BaseState
}

func (p *PlayState) PowerOff() {
	fmt.Println("关机...")
	p.Ctx.SetState(PowerOff)
	p.Ctx.State.PowerOff()
}

func (p *PlayState) Play() {
	fmt.Println("do...播放")
}

func (p *PlayState) Standby() {
	fmt.Println("待机...")
	p.Ctx.SetState(StandBy)
	p.Ctx.State.Standby()
}
