package plant

import "fmt"

// 植物インターフェース
type Plant interface {
	Grow() string
	Bloom() string
	Describe() string
	Water() Plant     // メソッドチェーン用
	Fertilize() Plant // メソッドチェーン用
	GetHealth() string
	ChangeName(newName string) Plant // 植物の名前を変更する
}

// 基本植物構造体
type BasePlant struct {
	Name    string
	Age     int // 年数
	Watered bool
	Healthy bool
}

// 花構造体
type Flower struct {
	BasePlant
	Color string // 花の色
}

// 木構造体
type Tree struct {
	BasePlant
	Height float64 // 高さ（メートル）
}

// 定数
const (
	DefaultColor  = "赤"
	DefaultHeight = 1.0
)

// 花のコンストラクタ
func NewFlower(name string, age int) *Flower {
	return &Flower{
		BasePlant: BasePlant{
			Name:    name,
			Age:     age,
			Watered: false,
			Healthy: true,
		},
		Color: DefaultColor,
	}
}

// 木のコンストラクタ
func NewTree(name string, age int) *Tree {
	return &Tree{
		BasePlant: BasePlant{
			Name:    name,
			Age:     age,
			Watered: false,
			Healthy: true,
		},
		Height: DefaultHeight,
	}
}

// 花のインターフェース実装
func (f *Flower) Grow() string {
	f.Age++
	return "花がすくすく成長しています"
}

func (f *Flower) Bloom() string {
	return fmt.Sprintf("%sの美しい花が咲きました", f.Color)
}

func (f *Flower) Describe() string {
	return fmt.Sprintf("%s（%d年目、%sの花）", f.Name, f.Age, f.Color)
}

func (f *Flower) Water() Plant {
	f.Watered = true
	return f // メソッドチェーン用
}

func (f *Flower) Fertilize() Plant {
	f.Healthy = true
	return f // メソッドチェーン用
}

func (f *Flower) GetHealth() string {
	// 短縮変数宣言
	health := "普通"
	if f.Watered && f.Healthy {
		health = "とても元気"
	} else if f.Watered {
		health = "潤っている"
	} else if f.Healthy {
		health = "健康"
	} else {
		health = "元気がない"
	}
	return health
}

// 花専用のメソッド
func (f *Flower) SetColor(color string) {
	f.Color = color
}

// 名前を変えるメソッド
func (f *Flower) ChangeName(newName string) Plant {
	f.Name = newName
	return f
}

// 木のインターフェース実装
func (t *Tree) Grow() string {
	t.Age++
	t.Height += 0.5
	return "木がどんどん大きくなっています"
}

func (t *Tree) Bloom() string {
	return "木に小さな花が咲きました"
}

func (t *Tree) Describe() string {
	return fmt.Sprintf("%s（%d年目、高さ%.1fm）", t.Name, t.Age, t.Height)
}

func (t *Tree) Water() Plant {
	t.Watered = true
	return t // メソッドチェーン用
}

func (t *Tree) Fertilize() Plant {
	t.Healthy = true
	return t // メソッドチェーン用
}

func (t *Tree) GetHealth() string {
	// 短縮変数宣言
	health := "普通"
	if t.Watered && t.Healthy {
		health = "青々と茂っている"
	} else if t.Watered {
		health = "水分十分"
	} else if t.Healthy {
		health = "丈夫"
	} else {
		health = "枯れ気味"
	}
	return health
}

// 木専用のメソッド
func (t *Tree) SetHeight(height float64) {
	t.Height = height
}

// 名前を変えるメソッド
func (t *Tree) ChangeName(newName string) Plant {
	t.Name = newName
	return t
}
