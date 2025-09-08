package objects

import "nso/entity"

type LevelManager struct {
	levelToPPoint map[int]int16
	levelToSPoint map[int]int16
	levels        []*entity.LevelEntity
	levelToMaxExp map[int]int64
}

func NewLevelManager(levels []*entity.LevelEntity) *LevelManager {
	this := &LevelManager{}
	this.levelToPPoint = make(map[int]int16, len(levels))
	this.levelToSPoint = make(map[int]int16, len(levels))
	this.levelToMaxExp = make(map[int]int64, len(levels))
	this.levels = levels
	for _, level := range levels {
		var exp int64 = 0
		for i := 0; i < level.Level; i++ {
			exp += this.GetLevel(i).Exps
		}
		this.levelToMaxExp[level.Level] = exp

		var sPoint int
		for i := 0; i < len(this.levels); i++ {
			if this.GetLevel(i).Level <= level.Level {
				sPoint += this.GetLevel(i).Spoint
			}
		}
		this.levelToSPoint[level.Level] = int16(sPoint)

		var ppoint int
		for i := 0; i < len(this.levels); i++ {
			if this.GetLevel(i).Level <= level.Level {
				ppoint += this.GetLevel(i).Ppoint
			}
		}
		this.levelToPPoint[level.Level] = int16(ppoint)
	}
	return this
}

func (l *LevelManager) GetLevel(level int) *entity.LevelEntity {
	if level < 0 || level >= len(l.levels) {
		return l.levels[0]
	}
	return l.levels[level]
}

func (l *LevelManager) GetMaxExp() int64 {
	return l.levelToMaxExp[len(l.levels)-1]
}

func (l *LevelManager) GetLevelAndRemainExp(exp int64) (int, int64) {
	for i := 0; i < len(l.levels); i++ {
		if exp < l.levelToMaxExp[i] {
			if i > 0 {
				return i - 1, exp - l.levelToMaxExp[i-1]
			}
		}
	}
	return len(l.levels) - 1, 0
}

func (l *LevelManager) GetLevelAndRemainExp1(exp int64) (int, int64) {
	var num = exp
	var i = 0
	for i = 0; i < len(l.levels) && num >= l.levels[i].Exps; num, i = num-l.levels[i].Exps, i+1 {
	}
	return i, num
}

func (l *LevelManager) TotalSPoint(level int) int16 {
	if level < 0 || level >= len(l.levels) {
		return 0
	}
	return l.levelToSPoint[level]
}

func (l *LevelManager) TotalPPoint(level int) int16 {
	if level < 0 || level >= len(l.levels) {
		return 0
	}
	return l.levelToPPoint[level]
}

func (l *LevelManager) MaxLevel(level int) int64 {
	if level < 0 || level >= len(l.levels) {
		return 0
	}
	return l.levelToMaxExp[level]
}
