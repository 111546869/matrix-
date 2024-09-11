// LinkedList.h
#ifndef LINKEDLIST_H
#define LINKEDLIST_H

class LinkedList {
public:
    struct Node {
        int data;
        Node* next;
        Node(int val) : data(val), next(nullptr) {}
    };
    
    LinkedList() : head(nullptr) {}
    ~LinkedList() { clear(); }
    
    void insert(int data);
    bool search(int data);
    void remove(int data);
    void clear();
    
private:
    Node* head;
};

#endif // LINKEDLIST_H
