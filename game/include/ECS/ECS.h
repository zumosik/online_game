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
class Manager;

using ComponentID = std::size_t;
using Group = std::size_t;

inline ComponentID getNewComponentTypeID()  {
    static ComponentID lastID = 0;
    return lastID++;
}

template <typename T> inline ComponentID  getComponentTypeID() noexcept {
    static ComponentID  typeID = getNewComponentTypeID();
    return typeID;
}

constexpr  std::size_t  maxComponents = 32;
constexpr  std::size_t  maxGroups = 32;

using ComponentBitSet = std::bitset<maxComponents>;
using GroupBitset = std::bitset<maxGroups>;
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
class Entity
{
private:
    Manager& manager;
    bool active = true;
    std::vector<std::unique_ptr<Component>> components;

    ComponentArray componentArray;
    ComponentBitSet componentBitset;
    GroupBitset groupBitset;

public:
    Entity(Manager& mManager) : manager(mManager) {}

    void update()
    {
        for (auto& c : components) c->update();
    }
    void draw()
    {
        for (auto& c : components) c->draw();
    }

    bool isActive() const { return active; }
    void destroy() { active = false; }

    bool hasGroup(Group mGroup)
    {
        return groupBitset[mGroup];
    }

    void addGroup(Group mGroup);
    void delGroup(Group mGroup)
    {
        groupBitset[mGroup] = false;
    }

    template <typename T> bool hasComponent() const
    {
        return componentBitset[getComponentTypeID<T>()];
    }

    template <typename T, typename... TArgs>
    T& addComponent(TArgs&&... mArgs)
    {
        T* c(new T(std::forward<TArgs>(mArgs)...));
        c->entity = this;
        std::unique_ptr<Component>uPtr { c };
        components.emplace_back(std::move(uPtr));

        componentArray[getComponentTypeID<T>()] = c;
        componentBitset[getComponentTypeID<T>()] = true;

        c->init();
        return *c;
    }

    template<typename T> T& getComponent() const
    {
        auto ptr(componentArray[getComponentTypeID<T>()]);
        return *static_cast<T*>(ptr);
    }
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
        for(auto i(0u); i < maxGroups; i++) {
            auto & v(groupedEntities[i]);
            v.erase(
                    std::remove_if(std::begin(v), std::end(v),
                                   [i](Entity* mEntity)
                                   {
                                       return !mEntity->isActive() || !mEntity->hasGroup(i);
                                   }),
                    std::end(v));

        }

        entites.erase(std::remove_if(std::begin(entites), std::end(entites),
                                     [](const std::unique_ptr<Entity> &mEntity)
                                     {
                                         return !mEntity->isActive();
                                     }),std::end(entites));
    }

    void AddToGroup(Entity * p_entity, Group p_group) {
        groupedEntities[p_group].emplace_back(p_entity);
    }

    std::vector<Entity*>& getGroup(Group p_group) {
        return groupedEntities[p_group];
    }

    Entity&  addEntity()
    {
        Entity* e = new Entity(*this);
        std::unique_ptr<Entity> uniquePtr{e};
        entites.emplace_back(std::move(uniquePtr));
        return  *e;
    }

private:
    std::vector<std::unique_ptr<Entity>> entites;
    std::array<std::vector<Entity*>, maxGroups> groupedEntities;
};

#endif //GAME_SDL2_ECS_H
