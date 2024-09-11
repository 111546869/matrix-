// test_linkedlist.cpp
#include "gtest/gtest.h"
#include "LinkedList.h"

// 测试插入节点功能
TEST(LinkedListTest, InsertTest) {
    LinkedList list;
    list.insert(5);
    EXPECT_TRUE(list.search(5));  // 确认5插入成功
}

// 测试查找功能
TEST(LinkedListTest, SearchTest) {
    LinkedList list;
    list.insert(10);
    EXPECT_TRUE(list.search(10));  // 查找已插入元素
    EXPECT_FALSE(list.search(20)); // 查找不存在的元素
}

// 测试删除节点功能
TEST(LinkedListTest, RemoveTest) {
    LinkedList list;
    list.insert(5);
    list.insert(10);
    list.remove(5);
    EXPECT_FALSE(list.search(5));  // 确认5被删除
    EXPECT_TRUE(list.search(10));  // 10应该还存在
}

// 测试链表清空
TEST(LinkedListTest, ClearTest) {
    LinkedList list;
    list.insert(1);
    list.insert(2);
    list.clear();
    EXPECT_FALSE(list.search(1));  // 清空后应该找不到元素
    EXPECT_FALSE(list.search(2));
}
