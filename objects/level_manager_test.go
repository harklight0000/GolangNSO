package objects

import (
	"nso/database"
	"nso/entity"
	"testing"
)

var exps = []int64{
	57575935,
	80500,
	80500,
	80500,
	80500,
	141826167711,
	45684610755,
	142402,
	1146576195936,
	120472,
	120472,
	80500,
	169036225876,
	187403497900,
	159442471070,
	201190030458,
	2119507171240,
	4338095503,
	587844021716,
	80500,
	24192137663,
	9128015221,
	51110745370,
	38027351133,
	14255451671,
	5736463915,
	80503,
	5285284457,
	898915197112,
	39659560932,
	673846027927,
	81730,
	149376,
	232494210,
	204956254833,
	125332,
	80500,
	40552575932,
	1760261013312,
	80500,
	70345671919,
	200321,
	3074650458,
}

var lvs = []int{
	36,
	10,
	10,
	10,
	10,
	76,
	69,
	11,
	95,
	11,
	11,
	10,
	77,
	78,
	76,
	78,
	103,
	50,
	87,
	10,
	64,
	55,
	69,
	67,
	60,
	51,
	10,
	51,
	92,
	68,
	89,
	10,
	11,
	41,
	78,
	11,
	10,
	68,
	101,
	10,
	71,
	12,
	48,
}

func TestLevelManager(t *testing.T) {
	db := database.NewSQLDatabaseAdapter(database.InitSQLDB())
	var levels []*entity.LevelEntity
	defer db.Close()
	err := db.FindAll(&levels)
	if err != nil {
		t.Error(err)
	}
	if len(levels) == 0 {
		t.Error("No levels loaded")
	}
	levelManager := NewLevelManager(levels)
	level, remain := levelManager.GetLevelAndRemainExp(23743462466)
	if level != 64 {
		t.Error("Expected level 64, got", level)
	}
	if remain <= 0 {
		t.Error("Expected remain exp > 0, got", remain)
	}
	for i, exp := range exps {
		level1, remain1 := levelManager.GetLevelAndRemainExp1(exp)
		level2, remain2 := levelManager.GetLevelAndRemainExp(exp)
		if level1 != level2 {
			t.Error("Expected level", lvs[i], "got", level2)
		}
		if remain1 != remain2 {
			t.Error("Expected remain exp", remain1, "got", remain2)
		}
		if level1 != lvs[i] {
			t.Error("Expected level", lvs[i], "got", level)
		}
	}
	lv, exp := levelManager.GetLevelAndRemainExp(0)
	if lv != 1 {
		t.Error("Expected level 1, got", lv)
	}
	if exp != 0 {
		t.Error("Expected remain exp 0, got", exp)
	}
	lv, exp = levelManager.GetLevelAndRemainExp(199)
	if lv != 1 {
		t.Error("Expected level 1, got", lv)
	}
	if exp != 199 {
		t.Error("Expected remain exp 199, got", exp)
	}
}

func BenchmarkLevelTest_V1(t *testing.B) {
	db := database.NewSQLDatabaseAdapter(database.InitSQLDB())
	var levels []*entity.LevelEntity
	err := db.FindAll(&levels)
	if err != nil {
		t.Error(err)
	}
	if len(levels) == 0 {
		t.Error("No levels loaded")
	}
	levelManager := NewLevelManager(levels)
	for i := 0; i < t.N; i++ {
		levelManager.GetLevelAndRemainExp(int64(i % len(levelManager.levels)))
	}
}

func BenchmarkLevelTest_V2(t *testing.B) {
	db := database.NewSQLDatabaseAdapter(database.InitSQLDB())
	var levels []*entity.LevelEntity
	err := db.FindAll(&levels)
	if err != nil {
		t.Error(err)
	}
	if len(levels) == 0 {
		t.Error("No levels loaded")
	}
	levelManager := NewLevelManager(levels)
	for i := 0; i < t.N; i++ {
		levelManager.GetLevelAndRemainExp1(int64(i % len(levelManager.levels)))
	}
}
