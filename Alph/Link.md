## 静态链表
    struct node {
        Type data
        int cur 
    }
## 动态链表
    struct node {
        Type data
        struct node * next
    }

### 区别就在于两点:
1. 动态链表需要动态分配内存(malloc && free)，静态链表数组提前分配好
2. 动态链表还是使用指针，静态使用cur数组位置

静态链表为了表示哪些位置可用，还要有一个空链表存储可用位置

---

稀疏矩阵采用三元组链表表示  
head->(row_num, col_num, data)->...
