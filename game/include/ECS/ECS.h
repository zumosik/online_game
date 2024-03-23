// E - entity
// C - component
// S - system

#ifndef GAME_SDL2_ECS_H
#define GAME_SDL2_ECS_H


#include "vector"
#include "memory"
#include "algorithm"
#include "bitset"
#include "array"

class Component;
class Entity;

using ComponentID = std::size_t;

inline ComponentID getComponentTypeID()  {
    static ComponentID lastID = 0;
    return lastID++;
}

template <typename T> inline ComponentID  getComponentTypeID() noexcept {
    static ComponentID  typeID = getComponentTypeID();
    return typeID;
}

constexpr  std::size_t  maxComponents = 32;

using ComponentBitSet = std::bitset<maxComponents>;
using ComponentArray  = std::array<Component*, maxComponents>;

class Component
{
public:
    Entity* entity;

    virtual void init(){}
    virtual void update() {}
    virtual void draw(){}

    virtual ~Component() {}
};

class Entity {
public:
    void update() {
        for (auto& c: components) c->update();
    }
    void draw() {
        for (auto& c: components) c->draw();
    }
    bool isActive() const {return active;}
    void destroy() {active = false;}

    template<typename T> bool hasComponent() const {
        return componentBitSet[getComponentTypeID<T>];
    }

    template<typename T, typename... TArgs>
    T& addComponent(TArgs&&... mArgs) {
        T* c(new T(std::forward<TArgs>(mArgs)...));
        c->entity = this;
        std::unique_ptr<Component> uniquePtr{c};
        components.emplace_back(std::move(uniquePtr));

        componentArray[getComponentTypeID<T>()] = c;
        componentBitSet[getComponentTypeID<T>()] = true;

        c->init();
        return *c;
    }

    template<typename T> T& getComponent() const {
        return * static_cast<T*>(componentArray[getComponentTypeID<T>()]);
    }

private:
    bool active = true;
    std::vector<std::unique_ptr<Component>> components;

    ComponentArray  componentArray;
    ComponentBitSet  componentBitSet;
};

class Manager {
public:
    void update() {
        for (auto& entity: entites) entity->update();
    }

    void draw() {
        for (auto& entity: entites) entity->draw();
    }

    void refresh()
    {
        entites.erase(std::remove_if(std::begin(entites), std::end(entites),
                                     [](const std::unique_ptr<Entity> &mEntity)
                                     {
                                         return !mEntity->isActive();
                                     }),std::end(entites));
    }

    Entity&  addEntity()
    {
        Entity* e = new Entity();
        std::unique_ptr<Entity> uniquePtr{e};
        entites.emplace_back(std::move(uniquePtr));
        return  *e;
    }

private:
    std::vector<std::unique_ptr<Entity>> entites;
};

#endif //GAME_SDL2_ECS_H
