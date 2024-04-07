package cm_test

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"online_game/internal/game/cm"
	"online_game/internal/game/cm/components"
	"reflect"
	"testing"
)

// TODO: add tests

func TestGameObject_Components(t *testing.T) {
	manager := cm.NewManager()

	t.Run("check", func(t *testing.T) {
		obj := manager.CreateGameObject()

		cmp := &components.TransformComponent{
			Pos: rl.Vector2{
				X: 2, Y: 3,
			},
			Size: rl.Vector2{
				X: 100, Y: 200.001,
			},
			Rotation: 35.447,
			Scale:    rl.Vector2{X: 2, Y: 3},
		}

		obj.AddComponent(cmp)

		if !obj.HasComponent(&components.TransformComponent{}) {
			t.Error("has component doesnt work")
		} else {
			t.Log("has component - ok")
		}

		if !reflect.DeepEqual(
			obj.GetComponent(&components.TransformComponent{}), cmp) {
			t.Error("get component doesnt work")
		} else {
			t.Log("get component - ok")
		}
	})

}
