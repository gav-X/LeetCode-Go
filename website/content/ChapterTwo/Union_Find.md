---
title: Union Find
type: docs
---

# Union Find

![](https://img.halfrost.com/Leetcode/Union_Find.png)

- 灵活使用并查集的思想，熟练掌握并查集的[模板]({{< relref "/ChapterThree/UnionFind.md" >}})，模板中有两种并查集的实现方式，一种是路径压缩 + 秩优化的版本，另外一种是计算每个集合中元素的个数 + 最大集合元素个数的版本，这两种版本都有各自使用的地方。能使用第一类并查集模板的题目有：第 128 题，第 130 题，第 547 题，第 684 题，第 721 题，第 765 题，第 778 题，第 839 题，第 924 题，第 928 题，第 947 题，第 952 题，第 959 题，第 990 题。能使用第二类并查集模板的题目有：第 803 题，第 952 题。第 803 题秩优化和统计集合个数这些地方会卡时间，如果不优化，会 TLE。
- 并查集是一种思想，有些题需要灵活使用这种思想，而不是死套模板，如第 399 题，这一题是 stringUnionFind，利用并查集思想实现的。这里每个节点是基于字符串和 map 的，而不是单纯的用 int 节点编号实现的。
- 有些题死套模板反而做不出来，比如第 685 题，这一题不能路径压缩和秩优化，因为题目中涉及到有向图，需要知道节点的前驱节点，如果路径压缩了，这一题就没法做了。这一题不需要路径压缩和秩优化。
- 灵活的抽象题目给的信息，将给定的信息合理的编号，使用并查集解题，并用 map 降低时间复杂度，如第 721 题，第 959 题。
- 关于地图，砖块，网格的题目，可以新建一个特殊节点，将四周边缘的砖块或者网格都 union() 到这个特殊节点上。第 130 题，第 803 题。
- 能用并查集的题目，一般也可以用 DFS 和 BFS 解答，只不过时间复杂度会高一点。



| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0128|Longest Consecutive Sequence|[Go]({{< relref "/ChapterFour/0128.Longest-Consecutive-Sequence.md" >}})|Hard| O(n)| O(n)|❤️|46.0%|
|0130|Surrounded Regions|[Go]({{< relref "/ChapterFour/0130.Surrounded-Regions.md" >}})|Medium| O(m\*n)| O(m\*n)||29.2%|
|0200|Number of Islands|[Go]({{< relref "/ChapterFour/0200.Number-of-Islands.md" >}})|Medium| O(m\*n)| O(m\*n)||48.6%|
|0399|Evaluate Division|[Go]({{< relref "/ChapterFour/0399.Evaluate-Division.md" >}})|Medium| O(n)| O(n)||54.0%|
|0547|Number of Provinces|[Go]({{< relref "/ChapterFour/0547.Number-of-Provinces.md" >}})|Medium| O(n^2)| O(n)||60.1%|
|0684|Redundant Connection|[Go]({{< relref "/ChapterFour/0684.Redundant-Connection.md" >}})|Medium| O(n)| O(n)||58.7%|
|0685|Redundant Connection II|[Go]({{< relref "/ChapterFour/0685.Redundant-Connection-II.md" >}})|Hard| O(n)| O(n)||32.9%|
|0721|Accounts Merge|[Go]({{< relref "/ChapterFour/0721.Accounts-Merge.md" >}})|Medium| O(n)| O(n)|❤️|51.2%|
|0765|Couples Holding Hands|[Go]({{< relref "/ChapterFour/0765.Couples-Holding-Hands.md" >}})|Hard| O(n)| O(n)|❤️|55.3%|
|0778|Swim in Rising Water|[Go]({{< relref "/ChapterFour/0778.Swim-in-Rising-Water.md" >}})|Hard| O(n^2)| O(n)|❤️|54.4%|
|0803|Bricks Falling When Hit|[Go]({{< relref "/ChapterFour/0803.Bricks-Falling-When-Hit.md" >}})|Hard| O(n^2)| O(n)|❤️|31.2%|
|0839|Similar String Groups|[Go]({{< relref "/ChapterFour/0839.Similar-String-Groups.md" >}})|Hard| O(n^2)| O(n)||40.0%|
|0924|Minimize Malware Spread|[Go]({{< relref "/ChapterFour/0924.Minimize-Malware-Spread.md" >}})|Hard| O(m\*n)| O(n)||41.8%|
|0928|Minimize Malware Spread II|[Go]({{< relref "/ChapterFour/0928.Minimize-Malware-Spread-II.md" >}})|Hard| O(m\*n)| O(n)|❤️|41.2%|
|0947|Most Stones Removed with Same Row or Column|[Go]({{< relref "/ChapterFour/0947.Most-Stones-Removed-with-Same-Row-or-Column.md" >}})|Medium| O(n)| O(n)||55.4%|
|0952|Largest Component Size by Common Factor|[Go]({{< relref "/ChapterFour/0952.Largest-Component-Size-by-Common-Factor.md" >}})|Hard| O(n)| O(n)|❤️|36.1%|
|0959|Regions Cut By Slashes|[Go]({{< relref "/ChapterFour/0959.Regions-Cut-By-Slashes.md" >}})|Medium| O(n^2)| O(n^2)|❤️|66.8%|
|0990|Satisfiability of Equality Equations|[Go]({{< relref "/ChapterFour/0990.Satisfiability-of-Equality-Equations.md" >}})|Medium| O(n)| O(n)||46.4%|
|1202|Smallest String With Swaps|[Go]({{< relref "/ChapterFour/1202.Smallest-String-With-Swaps.md" >}})|Medium||||48.3%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Bit_Manipulation/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterOne/Sliding_Window/">下一页➡️</a></p>
</div>
