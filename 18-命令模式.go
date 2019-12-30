package design

import "fmt"

// 接收者
// 真正的业务逻辑处理者
type Receiver struct {
}

func (r *Receiver) DoSomething() {
	fmt.Println("执行业务逻辑")
}

// 命令接口
// 不同的命令都实现这个接口
type ICommand interface {
	Execute()
}

// 具体命令
// 每个命令都对应一个实例对象，都实现ICommand接口，使命令调用统一
// 包含接收者，因为要利用接收者实现具体命令
type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) Init(r Receiver)  {
	c.receiver = r
}

func (c *ConcreteCommand) Execute() {
	c.receiver.DoSomething()
}

// 请求者
// 根据拥有不同的具体命令，而执行不同命令
type Invoker struct {
	C ICommand
}

func (i *Invoker) Action() {
	i.C.Execute()
}

// 实例
type TV struct {
}

func (t *TV) Open() {
	fmt.Println("打开电视")
}

func (t *TV) Close() {
	fmt.Println("关闭电视")
}

func (t *TV) Play() {
	fmt.Println("播放节目")
}

type ITvCommand interface {
	Execute()
}

type TvOpenCommand struct {
	tv *TV
}

func (t *TvOpenCommand) Init(tv *TV) {
	t.tv = tv
}

func (t *TvOpenCommand) Execute() {
	t.tv.Open()
}

type TvCloseCommand struct {
	tv *TV
}

func (t *TvCloseCommand) Init(tv *TV) {
	t.tv = tv
}

func (t *TvCloseCommand) Execute() {
	t.tv.Close()
}

type TvPlayCommand struct {
	tv *TV
}

func (t *TvPlayCommand) Init(tv *TV) {
	t.tv = tv
}

func (t *TvPlayCommand) Execute() {
	t.tv.Play()
}

type TvInvoker struct {
	open *TvOpenCommand
	close *TvCloseCommand
	play *TvPlayCommand
}

func (t *TvInvoker) Init(open *TvOpenCommand,close *TvCloseCommand,play *TvPlayCommand)  {
	t.open = open
	t.close = close
	t.play = play
}

func (t *TvInvoker) Open() {
	t.open.Execute()
}

func (t *TvInvoker) Close() {
	t.close.Execute()
}

func (t *TvInvoker) Play() {
	t.play.Execute()
}