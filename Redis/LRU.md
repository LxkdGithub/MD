### Redis的内存不是无限制的，所以要有驱逐策略eviction
1. 可以再redis.conf中设置maxmemory
1. 如果maxmemory=0或者说是noeviction时, 64位机器将是无限内存，32位将是3G大小。
2. policy有很多，总的来说由两部分构成，  
    (1). 一是evict的对象是谁有两种策略，分别是volatile(设置有过期时间的key) 和 allkeys (所有key)  
    (2). 二是evict的选择方式，lru(最近最少使用)和random  
    所以组成了allkeys-lru, volatile-lru, allkeys-random, volaile-random
    还有一种是volatile-ttl(对于设置了)过期时间的key选择未来存活时间最短是。
    
   
但是Redis并没有使用以上策略，而是使用近似lru的策略
1. 设置pool，大小默认是16吧，大概是
2. 然后使用随机原则取N个key(N可以自己设置，默认是5)
  找到pool和N个keys中的lru最大值对应的
  key，删掉该key后
3. 使用挑战然后替换的原则将N个key中4或者5个替换掉pool里
lru值比自己小的key，**简单来说**就是pool和N个keys删除一个最大
之后重新排序前16个最大是继续放进pool里，下一次使用

首先为什么这样?
加入直接使用lru，数据量太大，若要维护一个lru链表在信息量太大，不值得，但是lru信息不能不利用，
于是每个key都包含24位的最后访问时间信息，使用随机的方法，就不需要那么麻烦
但是这样做效果不太好，而且前边随机取过的信息都浪费了，
所以就放进pool里利用起来，将局部比较慢慢变为全局化的比较，
不仅不增加多少工作量还可以达到很好的效果。   
    