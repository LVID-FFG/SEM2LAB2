#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <algorithm>
#include <map>

#include "setInclude.h"

using namespace std;

// Функции красно-черного дерева
void leftRotate(RBTree* tree, RBNode* x){
    RBNode* y = x->right;
    x->right = y->left;
    if (y->left != nullptr){
        y->left->parent = x;
    }
    y->parent = x->parent;
    if (x->parent == nullptr){
        tree->root = y;
    }else{
        if(x == x->parent->left) x->parent->left = y;
        else x->parent->right = y;
    }
    y->left = x;
    x->parent = y;
}

void rightRotate(RBTree* tree, RBNode* x){
    RBNode* y = x->left;
    x->left = y->right;
    if (y->right != nullptr){
        y->right->parent = x;
    }
    y->parent = x->parent;
    if (x->parent == nullptr){
        tree->root = y;
    }else{
        if(x == x->parent->left) x->parent->left = y;
        else x->parent->right = y;
    }
    y->right = x;
    x->parent = y;
}

void fixAdd(RBTree* tree, RBNode* z){
    while (z->parent != nullptr && z->parent->color == Color::RED){
        
        if (z->parent->parent == nullptr) {
            break;
        }

        if(z->parent == z->parent->parent->left){
            RBNode* y = z->parent->parent->right;
            
            if (y != nullptr && y->color == Color::RED){
                z->parent->color = Color::BLACK;
                y->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                z = z->parent->parent;
            }
            else{
                if (z == z->parent->right){
                    z = z->parent;
                    leftRotate(tree, z);
                }
                z->parent->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                rightRotate(tree, z->parent->parent);
            }
        }
        else{
            RBNode* y = z->parent->parent->left;
            
            if (y != nullptr && y->color == Color::RED){
                z->parent->color = Color::BLACK;
                y->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                z = z->parent->parent;
            }
            else{
                if (z == z->parent->left){
                    z = z->parent;
                    rightRotate(tree, z);
                }
                z->parent->color = Color::BLACK;
                z->parent->parent->color = Color::RED;
                leftRotate(tree, z->parent->parent);
            }
        }
        
        if (z == tree->root) {
            break;
        }
    }
    tree->root->color = Color::BLACK;
}

void addRBNode(RBTree* tree, const string& value){
    RBNode* new_RBNode = new RBNode(value);
    if (tree->root == nullptr){
        new_RBNode->color = Color::BLACK;
        tree->root = new_RBNode;
        return;
    }
    RBNode* address = tree->root;
    while(true){
        if(value < address->data){
            if (address->left == nullptr){
                new_RBNode->parent = address;
                address->left = new_RBNode;
                fixAdd(tree, new_RBNode);
                return;
            }else{
                address = address->left;
            }
        }
        else{
            if (address->right == nullptr){
                new_RBNode->parent = address;
                address->right = new_RBNode;
                fixAdd(tree, new_RBNode);
                return;
            }else{
                address = address->right;
            }
        }
    }  
}

RBNode* getRBNode(const RBTree* tree, const string& value) {
    if (tree->root == nullptr) {
        return nullptr;
    }
    
    RBNode* address = tree->root;
    while (address != nullptr) {
        if (value == address->data) {
            return address;
        } else if (value < address->data) {
            address = address->left;
        } else {
            address = address->right;
        }
    }
    
    return nullptr;
}

RBNode* treeMinimum(RBTree* tree, RBNode* node) {
    if (node == nullptr) return nullptr;
    RBNode* address = node;
    while (address->left != nullptr) {
        address = address->left;
    }
    return address;
}

void deleteFix(RBTree* tree, RBNode* x) {
    // Если x nullptr, значит дерево пустое или это последний узел
    if (x == nullptr) return;
    
    while (x != tree->root && x->color == Color::BLACK) {
        if (x == x->parent->left) {
            RBNode* w = x->parent->right;
            if (w == nullptr) break;
            
            if (w->color == Color::RED) {
                w->color = Color::BLACK;
                x->parent->color = Color::RED;
                leftRotate(tree, x->parent);
                w = x->parent->right;
                if (w == nullptr) break;
            }
            
            bool leftBlack = (w->left == nullptr) || (w->left->color == Color::BLACK);
            bool rightBlack = (w->right == nullptr) || (w->right->color == Color::BLACK);
            
            if (leftBlack && rightBlack) {
                w->color = Color::RED;
                x = x->parent;
            } else {
                if (w->right == nullptr || w->right->color == Color::BLACK) {
                    if (w->left != nullptr) {
                        w->left->color = Color::BLACK;
                    }
                    w->color = Color::RED;
                    rightRotate(tree, w);
                    w = x->parent->right;
                    if (w == nullptr) break;
                }
                
                w->color = x->parent->color;
                x->parent->color = Color::BLACK;
                if (w->right != nullptr) {
                    w->right->color = Color::BLACK;
                }
                leftRotate(tree, x->parent);
                x = tree->root;
            }
        } else {
            RBNode* w = x->parent->left;
            if (w == nullptr) break;
            
            if (w->color == Color::RED) {
                w->color = Color::BLACK;
                x->parent->color = Color::RED;
                rightRotate(tree, x->parent);
                w = x->parent->left;
                if (w == nullptr) break;
            }
            
            bool leftBlack = (w->left == nullptr) || (w->left->color == Color::BLACK);
            bool rightBlack = (w->right == nullptr) || (w->right->color == Color::BLACK);
            
            if (leftBlack && rightBlack) {
                w->color = Color::RED;
                x = x->parent;
            } else {
                if (w->left == nullptr || w->left->color == Color::BLACK) {
                    if (w->right != nullptr) {
                        w->right->color = Color::BLACK;
                    }
                    w->color = Color::RED;
                    leftRotate(tree, w);
                    w = x->parent->left;
                    if (w == nullptr) break;
                }
                
                w->color = x->parent->color;
                x->parent->color = Color::BLACK;
                if (w->left != nullptr) {
                    w->left->color = Color::BLACK;
                }
                rightRotate(tree, x->parent);
                x = tree->root;
            }
        }
    }
    
    if (x != nullptr) {
        x->color = Color::BLACK;
    }
}

void transplant(RBTree* tree, RBNode* u, RBNode* v) {
    if (u->parent == nullptr) {
        tree->root = v;
    } else if (u == u->parent->left) {
        u->parent->left = v;
    } else {
        u->parent->right = v;
    }
    
    if (v != nullptr) {
        v->parent = u->parent;
    }
}

void delRBNode(RBTree* tree, const string& value) {
    RBNode* z = getRBNode(tree, value);
    if (z == nullptr) return;
    
    RBNode* y = z;
    RBNode* x = nullptr;
    Color y_original_color = y->color;
    
    if (z->left == nullptr) {
        x = z->right;
        transplant(tree, z, x);
    } else if (z->right == nullptr) {
        x = z->left;
        transplant(tree, z, x);
    } else {
        y = treeMinimum(tree, z->right);
        y_original_color = y->color;
        x = y->right;
        
        if (y->parent != z) {
            transplant(tree, y, x);
            y->right = z->right;
            if (y->right != nullptr) {
                y->right->parent = y;
            }
        } else {
            if (x != nullptr) {
                x->parent = y;
            }
        }
        
        transplant(tree, z, y);
        y->left = z->left;
        if (y->left != nullptr) {
            y->left->parent = y;
        }
        y->color = z->color;
    }
    
    delete z;
    
    if (y_original_color == Color::BLACK) {
        if (x != nullptr) {
            deleteFix(tree, x);
        }
        // Если x nullptr, значит удалили последний узел
    }
}