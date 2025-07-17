package game

type Player struct {
	Hunger       int
	Water        int
	Health       int
	Infection    int
	IQ           int
	CraftSkill   int
	AttackSkill  int
	Inventory    map[string]int
	Location     string
	SpeedBoost   bool
	ReadingBook  string
	BookProgress map[string]int
}

type HotDrop struct {
	Location string   // where the hotdrop is located
	Items    []string // 2 random items from hotdrop list
	IsActive bool     // whether it's available
}

type GameState struct {
	Day         int
	IsNight     bool
	Player      Player
	ZombieCount int
	HotDrop     *HotDrop
	Log         []string
	TimePassed  int // in minutes
}

type ActionOption struct {
	ID     int
	Label  string                  // e.g. "Eat Apple"
	Effect func(*GameState) string // applies the effect, returns log message
}

func DefaultInventory() map[string]int {
	return map[string]int{
		"Apple":  2,
		"Rabbit": 1,
		"Water":  2,
		"Stone":  10,
		"Wood":   5,
	}
}

func NewGameState() GameState {
	player := Player{
		Hunger:       70,
		Water:        70,
		Health:       100,
		Infection:    0,
		IQ:           0,
		CraftSkill:   0,
		AttackSkill:  0,
		Inventory:    DefaultInventory(),
		Location:     "Home",
		SpeedBoost:   false,
		ReadingBook:  "",
		BookProgress: map[string]int{"Crafting": 0, "Survival": 0},
	}

	return GameState{
		Day:         1,
		IsNight:     false,
		Player:      player,
		ZombieCount: 0,
		HotDrop:     nil,
		Log:         []string{"Game started. You are at Home."},
		TimePassed:  0,
	}
}
