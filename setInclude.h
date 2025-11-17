#pragma once

#include <iostream>
#include <vector>

using namespace std;

enum class Color{
    RED,
    BLACK
};

class RBNode{
    public:
        string data;
        Color color;
        RBNode* parent;
        RBNode* left;
        RBNode* right;
        
        RBNode(const string& value) : data(value), color(Color::RED), parent(nullptr), left(nullptr), right(nullptr) {}
};

// Предварительные объявления функций красно-черного дерева
class RBTree;
void addRBNode(RBTree* tree, const string& value);
void delRBNode(RBTree* tree, const string& value);
RBNode* getRBNode(const RBTree* tree, const string& value);

class RBTree{
    public:
        RBNode* root;
        RBTree() : root(nullptr) {}
        
        // Итератор для обхода дерева
        class Iterator {
        private:
            RBNode* current;
            RBNode* findMin(RBNode* node) {
                while (node && node->left) node = node->left;
                return node;
            }
            
            RBNode* findSuccessor(RBNode* node) {
                if (node->right) return findMin(node->right);
                
                RBNode* parent = node->parent;
                while (parent && node == parent->right) {
                    node = parent;
                    parent = parent->parent;
                }
                return parent;
            }
            
        public:
            Iterator(RBNode* node) : current(node) {}
            
            string& operator*() { return current->data; }
            string* operator->() { return &current->data; }
            
            Iterator& operator++() {
                current = findSuccessor(current);
                return *this;
            }
            
            Iterator operator++(int) {
                Iterator temp = *this;
                ++(*this);
                return temp;
            }
            
            bool operator==(const Iterator& other) const { return current == other.current; }
            bool operator!=(const Iterator& other) const { return current != other.current; }
        };
        
        Iterator begin() const {
            return Iterator(findMin(root));
        }
        
        Iterator end() const {
            return Iterator(nullptr);
        }
        
    private:
        RBNode* findMin(RBNode* node) const {
            while (node && node->left) node = node->left;
            return node;
        }
};

class StringSet {
private:
    RBTree tree;

public:
    // Конструкторы
    StringSet() = default;
    
    StringSet(const vector<string>& elements) {
        for (const auto& elem : elements) {
            add(elem);
        }
    }
    
    // Базовые операции множества
    void add(const string& value) {
        //if (contains(value)) return;
        ::addRBNode(&tree, value);
    }
    
    void remove(const string& value) {
        ::delRBNode(&tree, value);
    }
    
    bool contains(const string& value) const {
        return ::getRBNode(&tree, value) != nullptr;
    }
    
    size_t size() const {
        size_t count = 0;
        for (auto it = begin(); it != end(); ++it) count++;
        return count;
    }
    
    bool empty() const {
        return tree.root == nullptr;
    }
    
    void clear() {
        clearTree(tree.root);
        tree.root = nullptr;
    }
    
    // Операции множества
    StringSet unionWith(const StringSet& other) const {
        StringSet result = *this;
        for (const auto& elem : other) {
            result.add(elem);
        }
        return result;
    }
    
    StringSet intersectionWith(const StringSet& other) const {
        StringSet result;
        for (const auto& elem : *this) {
            if (other.contains(elem)) {
                result.add(elem);
            }
        }
        return result;
    }
    
    StringSet differenceWith(const StringSet& other) const {
        StringSet result;
        for (const auto& elem : *this) {
            if (!other.contains(elem)) {
                result.add(elem);
            }
        }
        return result;
    }
    
    // Проверка принадлежности
    bool operator==(const StringSet& other) const {
        if (size() != other.size()) return false;
        for (const auto& elem : *this) {
            if (!other.contains(elem)) return false;
        }
        return true;
    }
    
    bool operator!=(const StringSet& other) const {
        return !(*this == other);
    }
    
    // Итераторы
    RBTree::Iterator begin() const { return tree.begin(); }
    RBTree::Iterator end() const { return tree.end(); }
    
    // Вывод
    void print() {
        for (RBTree::Iterator i = (*this).begin(); i != (*this).end(); i++) cout << *i << " ";
        cout << endl;
    }

private:
    void clearTree(RBNode* node) {
        if (node == nullptr) return;
        clearTree(node->left);
        clearTree(node->right);
        delete node;
    }
};