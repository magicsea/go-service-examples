package birth

// 场景中 球出生点 辅助类

import (
	bmath "base/math"
	"math"
	"math/rand"
	"roomserver/conf"
	"usercmd"
)

type BirthPoints struct {
	birthPoints []*BirthPoint // 球出生点
}

func (b *BirthPoints) AddBirthPoint(point *BirthPoint) {
	b.birthPoints = append(b.birthPoints, point)
}

func (b *BirthPoints) RefreshBirthPoint(d float64, scene IScene) {
	for _, birth := range b.birthPoints {
		birth.Refresh(d, scene)
	}
}

// 生成食物 、 动态障碍物
func (b *BirthPoints) CreateAllBirthPoint(scene IScene) {
	items := conf.ConfigMgr_GetMe().GetXmlFoodItems(scene.SceneID())
	for _, item := range items.Items {
		ftype := item.FoodType
		fid := item.FoodId
		birthTime := item.BirthTime
		foodnum := conf.ConfigMgr_GetMe().GetFoodMapNum(scene.SceneID(), fid)
		size := float64(conf.ConfigMgr_GetMe().GetFoodSize(scene.SceneID(), fid))
		if foodnum > 0 {
			for i := 0; i < int(foodnum); i++ {
				x, y := scene.GetRandPos()
				point := NewBirthPoint(scene.GenBallID(), float32(x), float32(y), float32(size), float32(size), fid, ftype, birthTime, 1, scene)
				b.AddBirthPoint(point)
			}
		}
	}
}

// 不同的食物，初始位置会做调整。如 食物(普通) 根据输入x,y 随机附近的值； 如食物(锤子) 根据输入x,y 对齐到地图上对应格子的中心 等等
func BallFood_InitPos(pos *bmath.Vector2, t usercmd.BallType, birthRadiusMin, birthRadiusMax float32) *bmath.Vector2 {
	switch t {
	case usercmd.BallType_FoodNormal:
		x := math.Floor(float64(pos.X)) + rand.Float64()*0.5
		y := math.Floor(float64(pos.Y)) + rand.Float64()*0.5
		posNew := &bmath.Vector2{float32(x), float32(y)}
		return posNew

	default:
		x := math.Floor(float64(pos.X)) + 0.25
		y := math.Floor(float64(pos.Y)) + 0.25
		posNew := &bmath.Vector2{float32(x), float32(y)}
		return posNew
	}
}
