package design

// 游戏接口
type IGame interface {
	GetName() string
	Run(us []*User)
}

type User struct {
	Name string
}

// 游戏场景
type Scene struct {
	users []*User
	game  IGame
}

func (s *Scene) Name() string {
	return s.game.GetName()
}

func (s *Scene) Run() {
	s.game.Run(s.users)
}

func (s *Scene) Enter(u *User) {
	s.users = append(s.users, u)
}

// 场景管理器
type SceneManager struct {
	Scene
}

func (s *SceneManager) GetPlaysNumber() int {
	return len(s.users)
}

// 场景管理器包含游戏场景对象，这是类的功能层次结构（继承）
// 游戏接口，实现游戏接口，这是实现层次结构（实现）
// 桥接模式就是将功能层次结构和实现层次结构相分离
