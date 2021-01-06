## Hashmap的存储
hashmap是数组加链表的形式，即其冲突的解决方式是使用链地址法，
首先说明几点
1. 数组大小默认是2的幂此，为什么呢? 以为其从key到数组index的对应是
这样计算的， index = hash(key)^(length-1) 加入长度是16就是hash值于15即可，只是用后四位对应
index，这就是为什么数组长度一般取2的幂次。
2. 填充因子 = key的数目/数组长度，超过固定数值比如0.7-0.8
就要rehash(rehash就要重新计算对应位置了)
当然可能会出现很多冲突，比如说极端情况，所有的key对应到一个位置
这时也要rehash，以为rehash大概率会解决这个冲突率太高的情况。

3. 但是链地址并不是简单链表的问题，当链表太长时轮训也是很长时间的，
所以要转为红黑树
4. key又是正面覆盖呢  
首先数组对应index没有值直接写入，
(1).如果有key的存在要比较的，若是空则覆盖，
(2).若相等也是覆盖，
若不相等
(3).是Node也就是链表时插入链表
(4).红黑树则是插入红黑树

5. 为什么是红黑树，而不是AVL树?
具体的细节也就不在这里赘述，不知不觉已经写了这么多了，直接说结论吧。AVL树的查找速度更快，但是相应的插入和修改的速度较慢。而红黑树则在插入和修改操作较为密集的时候表现更好。
而总结我们日常的HashMap使用，大多数情况下插入和修改应该是比查找更频繁一些的
而在这种情况下，红黑树的综合表现会更好一些。
5. 为什么链表的插入方式改成了尾部插入而不是头部插入?
假设头部插入按顺序来有B->A ,B后面来，B的next只想A
rehash之后，B先rehash，就变成了A->B,但是B的next还指向A(鬼知道为什么)
这样就会有循环出现，循环不是怕你查不到数据，**而是查找那些hash值和
A ，B一样但实际上不存在的key的时候回有死循环出现**

7. 这里先放一个数据结构
static class Node<K, V> implements Entry<K, V> {
        final int hash;
        final K key;
        V value;
        Node next;
        Node(int hash, K key, V value, Node next) {this.hash = hash;this.key = key;this.value = value;this.next = next;
        }
     ......public final V getValue() {return value;
        }public final V setValue(V newValue) {
            V oldValue = value;
            value = newValue;return oldValue;
        }
     ......
    }
可以看到就是hash，key，value，next指针组成的对象

8. LinkedHashMap
LinkedHashmap的entry就是比Hashmap多了Node * before和Node * after用于
构造一个隐式的链表来标明插入顺序的
