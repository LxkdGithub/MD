## 切片
1. 在切片数组多次append同一个底层数组值发生变化的切片时不能直接append,
要新开辟内存,然后copy data  
    ```
    (1). new = make([]int, len)
    (2). copy(new, old)
    (3). all = append(all, new)
    然后一直循环修改old再append
    ```
2. copy切片可以对共享底层数组的切片进行操作
    即使有重叠(其实就是一个个单纯的向前覆盖即可)
   
3. 不用说，切片cap>len时append不会改变底层数组地址，
    但是超出cap就会另外开辟内存空间
    这点虽然简单，但是也要看情况，若
    b = [][]int{}
    a = []int{}
    b = append(b, a)且随着a动态变化，那就不行