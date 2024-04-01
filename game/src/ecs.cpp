#include "ECS/ECS.h"

void Entity::addGroup(Group p_group) {
    groupBitset[p_group] = true;
    manager.AddToGroup(this, p_group);
}