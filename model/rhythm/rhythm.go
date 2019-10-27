package rhythm

import (
	"time"

	"github.com/cisordeng/beego/orm"
)

type RhythmSet struct { // 歌单
	Id int
	Nid int64

	Name string
	Avatar string
	PlayedCount int `orm:"default(0)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`

	Index int
}

func (this *RhythmSet) TableName() string {
	return "rhythm_rhythm_set"
}

type RhythmSetRhythm struct { // 歌单包含的单歌
	Id int
	RhythmSetId int
	RhythmId int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`

	Index int
}

func (this *RhythmSetRhythm) TableName() string {
	return "rhythm_rhythm_set_rhythm"
}

type Rhythm struct { // 单歌
	Id int // id
	Nid int64

	Name string
	Avatar string
	Url string
	Lyric string `orm:"type(text);default('')"`
	TranslatedLyric string `orm:"type(text);default('')"`
	SingerName string
	PlayedCount int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *Rhythm) TableName() string {
	return "rhythm_rhythm"
}

type RhythmListen struct { // 听歌记录
	Id int // id
	UserId int
	RhythmSetId int
	RhythmId int
	ListenerIp string
	ListenedAt time.Time `orm:"auto_now_add;type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *RhythmListen) TableName() string {
	return "rhythm_rhythm_listen"
}

func init() {
	orm.RegisterModel(new(RhythmSet))
	orm.RegisterModel(new(Rhythm))
	orm.RegisterModel(new(RhythmSetRhythm))
	orm.RegisterModel(new(RhythmListen))
}