# Components

<!-- TOC -->
* [Components](#components)
  * [Template for component:](#template-for-component-)
  * [Important](#important)
  * [Comments for components](#comments-for-components)
    * [Basic Template](#basic-template)
    * [If any method is empty](#if-any-method-is-empty)
    * [If Component have very important info to tell](#if-component-have-very-important-info-to-tell)
<!-- TOC -->

Components always implement Component interface (internal/cm/componentManager.go)
```go
type Component interface {
    Init(obj *GameObject)
    Update()
    Render()
}
```

## Template for component: 
```go
type TestComponent struct {
    // some info that component store

    obj *cm.GameObject // pointer to object that has this component
}

func (t *TestComponent) Init(obj *cm.GameObject) {
	t.obj = obj // set pointer to object that has this component
}

func (t *TestComponent) Update() { 
/*
	Called in every frame by managers Update func.
	Calculate, read input, etc.
	Mustn't render anything 
*/
}

func (t *TestComponent) Render() {
/*
    Called in every frame by managers Render func.
    If component has some graphic it needs to be rendered here.
*/
}
```
## Important
1. Make sure to save pointer to obj that will be given in Init. This is important to get other components of object.
2. Do **NOT** use *Update* for render 
3. Do **NOT** use *Render* for calculations not related to graphics.
4. Make sure that structs **pointer** implement methods 
5. Create comments for component. [How to write comments for components](#comments-for-components)

## Comments for components
### Basic Template
```go
// TestComponent stores value about Data1 and Data2.
//
// Init: info about Init
// Update: info about Update
// Render: info about Render
type TestComponent struct {
    // write about exported fields in "stores value ..."
    Data1 Type1 
    Data2 Type2

    // DON'T write about not exported fields in "stores value ..."
    notExportedData Type3
    
	
/*
    don't write anything about obj because it must be stored in every component, 
    so everything related to it is already known
*/
    obj *cm.GameObject // pointer to object that has this component
}
```
### If any method is empty
Use **none** if func is empty
```go
// TestComponent stores value about ... .
//
// Init: info about Init
// Update: info about Update
// Render: none
type TestComponent struct {
    obj *cm.GameObject // pointer to object that has this component
}
```
### If Component have very important info to tell
```go
// TestComponent stores value about ... .
//
// Init: info about Init
// Update: info about Update
// Render: info about Render
//
// # Very important info
```
```go
// TestComponent stores value about ... .
//
// Init: info about Init
// Update: info about Update
// Render: info about Render
//
// # Cant be used without TestComponent2
```