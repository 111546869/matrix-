// LinkedList.cpp
#include "LinkedList.h"
#include <iostream>

void LinkedList::insert(int data) {
    Node* newNode = new Node(data);
    if (!head) {
        head = newNode;
    } else {
        Node* temp = head;
        while (temp->next) {
            temp = temp->next;
        }
        temp->next = newNode;
    }
}

bool LinkedList::search(int data) {
    Node* temp = head;
    while (temp) {
        if (temp->data == data) return true;
        temp = temp->next;
    }
    return false;
}

void LinkedList::remove(int data) {
    if (!head) return;
    
    if (head->data == data) {
        Node* toDelete = head;
        head = head->next;
        delete toDelete;
        return;
    }

    Node* temp = head;
    while (temp->next && temp->next->data != data) {
        temp = temp->next;
    }
    
    if (temp->next) {
        Node* toDelete = temp->next;
        temp->next = temp->next->next;
        delete toDelete;
    }
}

void LinkedList::clear() {
    while (head) {
        Node* temp = head;
        head = head->next;
        delete temp;
    }
}
