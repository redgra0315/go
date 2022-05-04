package cat

type Cat struct {
	name  string `json:"name"`
	age   int    `json:"age"`
	color string `json:"color"`
	eat   string `json:"eat"`
	voice string `json:"voice"`
}

type BlackCat struct {
	Cat
}

// 构造函数
func NewCat(name string, age int, color, eat, voice string) *Cat {
	return &Cat{
		name:  name,
		age:   age,
		color: color,
		eat:   eat,
		voice: voice,
	}
}

// 构造函数
func NewBlackCat(name string, age int, color, eat, voice string) *BlackCat {
	return &BlackCat{Cat{
		name:  name,
		age:   age,
		color: color,
		eat:   eat,
		voice: voice,
	},
	}
}
