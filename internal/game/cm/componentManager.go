package cm

import (
	"log"
	"reflect"
)

type Layer string

const (
	LayerPlayer      Layer = "player"
	LayerOtherPlayer Layer = "otherPlayer"

	LayerDefault Layer = "default"
)

type Manager struct {
	objects []*GameObject
	lastId  uint16
}

func NewManager() *Manager {
	return &Manager{
		objects: make([]*GameObject, 0),
	}
}

func (m *Manager) CreateGameObject() *GameObject {
	o := &GameObject{}

	m.objects = append(m.objects, o)

	m.lastId++

	o.ID = m.lastId

	// to not get errors
	o.components = make([]Component, 0)
	o.Layer = LayerDefault // default layer

	return o
}

func (m *Manager) DeleteGameObject(obj *GameObject) {
	for i, o := range m.objects {
		if o == obj {
			m.objects = append(m.objects[:i], m.objects[i+1:]...)
			return
		}
	}
}

func (m *Manager) DeleteGameObjectByID(id uint16) {
	for i, o := range m.objects {
		if o.ID == id {
			m.objects = append(m.objects[:i], m.objects[i+1:]...)
			return
		}
	}
}

func (m *Manager) Update() {
	for _, o := range m.objects {
		o.Update()
	}
}

func (m *Manager) Render() {
	// render objects in order of layers
	for _, layer := range []Layer{
		LayerPlayer,
		LayerOtherPlayer, // here need to be all layers in order we want to render them

		LayerDefault,
	} {
		for _, o := range m.objects {
			if o.Layer == layer {
				o.Render()
			}
		}
	}
}

type GameObject struct {
	ID    uint16
	Layer Layer // public field to set layer

	components []Component
}

func (o *GameObject) AddComponent(component Component) {
	log.Printf("add component: %v", component)

	component.Init(o)
	o.components = append(o.components, component)
}

func (o *GameObject) GetComponent(componentType Component) Component {
	for _, c := range o.components {
		if reflect.TypeOf(c) == reflect.TypeOf(componentType) {
			return c
		}
	}
	return nil
}

func (o *GameObject) HasComponent(componentType Component) bool {
	for _, c := range o.components {
		if reflect.TypeOf(c) == reflect.TypeOf(componentType) {
			return true
		}
	}
	return false
}

func (o *GameObject) Update() {
	for _, c := range o.components {
		c.Update()
	}
}

func (o *GameObject) Render() {
	for _, c := range o.components {
		c.Render()
	}
}

type Component interface {
	Init(obj *GameObject)
	Update()
	Render()
}
