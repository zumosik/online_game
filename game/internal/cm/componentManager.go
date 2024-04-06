package cm

import (
	"log"
	"reflect"
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

	o.id = m.lastId

	// to not get errors
	o.components = make([]Component, 0)
	o.layers = make([]string, 0)

	return o
}

func (m *Manager) Update() {
	for _, o := range m.objects {
		o.Update()
	}
}

func (m *Manager) Render() {
	for _, o := range m.objects {
		o.Render()
	}
}

type GameObject struct {
	id     uint16
	layers []string // TODO

	components []Component
}

func (o *GameObject) AddComponent(component Component) {
	log.Printf("add component: %v", component)

	component.SetGameObject(o)
	component.Init()
	o.components = append(o.components, component)
}

func (o *GameObject) GetComponent(componentType interface{}) interface{} {
	for _, c := range o.components {
		if reflect.TypeOf(c) == reflect.TypeOf(componentType) {
			log.Printf("get component: %v", c)
			return c
		}
	}
	return nil
}

func (o *GameObject) HasComponent(componentType interface{}) bool {
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
	Init()
	Update()
	Render()
	SetGameObject(obj *GameObject)
}
