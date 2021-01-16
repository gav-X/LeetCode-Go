---
title: Sliding Window
type: docs
---

# Sliding Window

![](https://img.halfrost.com/Leetcode/Sliding_Window.png)

- 双指针滑动窗口的经典写法。右指针不断往右移，移动到不能往右移动为止(具体条件根据题目而定)。当右指针到最右边以后，开始挪动左指针，释放窗口左边界。第 3 题，第 76 题，第 209 题，第 424 题，第 438 题，第 567 题，第 713 题，第 763 题，第 845 题，第 881 题，第 904 题，第 978 题，第 992 题，第 1004 题，第 1040 题，第 1052 题。

```c
	left, right := 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
```
- 滑动窗口经典题。第 239 题，第 480 题。


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0003|Longest Substring Without Repeating Characters|[Go]({{< relref "/ChapterFour/0003.Longest-Substring-Without-Repeating-Characters.md" >}})|Medium| O(n)| O(1)|❤️|31.3%|
|0076|Minimum Window Substring|[Go]({{< relref "/ChapterFour/0076.Minimum-Window-Substring.md" >}})|Hard| O(n)| O(n)|❤️|35.7%|
|0239|Sliding Window Maximum|[Go]({{< relref "/ChapterFour/0239.Sliding-Window-Maximum.md" >}})|Hard| O(n * k)| O(n)|❤️|44.5%|
|0424|Longest Repeating Character Replacement|[Go]({{< relref "/ChapterFour/0424.Longest-Repeating-Character-Replacement.md" >}})|Medium| O(n)| O(1) ||48.0%|
|0480|Sliding Window Median|[Go]({{< relref "/ChapterFour/0480.Sliding-Window-Median.md" >}})|Hard| O(n * log k)| O(k)|❤️|38.4%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.6%|
|0978|Longest Turbulent Subarray|[Go]({{< relref "/ChapterFour/0978.Longest-Turbulent-Subarray.md" >}})|Medium| O(n)| O(1)|❤️|46.6%|
|0992|Subarrays with K Different Integers|[Go]({{< relref "/ChapterFour/0992.Subarrays-with-K-Different-Integers.md" >}})|Hard| O(n)| O(n)|❤️|50.4%|
|0995|Minimum Number of K Consecutive Bit Flips|[Go]({{< relref "/ChapterFour/0995.Minimum-Number-of-K-Consecutive-Bit-Flips.md" >}})|Hard| O(n)| O(1)|❤️|49.6%|
|1004|Max Consecutive Ones III|[Go]({{< relref "/ChapterFour/1004.Max-Consecutive-Ones-III.md" >}})|Medium| O(n)| O(1) ||60.5%|
|1040|Moving Stones Until Consecutive II|[Go]({{< relref "/ChapterFour/1040.Moving-Stones-Until-Consecutive-II.md" >}})|Medium| O(n log n)| O(1) |❤️|54.0%|
|1052|Grumpy Bookstore Owner|[Go]({{< relref "/ChapterFour/1052.Grumpy-Bookstore-Owner.md" >}})|Medium| O(n log n)| O(1) ||55.6%|
|1074|Number of Submatrices That Sum to Target|[Go]({{< relref "/ChapterFour/1074.Number-of-Submatrices-That-Sum-to-Target.md" >}})|Hard| O(n^3)| O(n) |❤️|61.4%|
|1208|Get Equal Substrings Within Budget|[Go]({{< relref "/ChapterFour/1208.Get-Equal-Substrings-Within-Budget.md" >}})|Medium||||43.7%|
|1658|Minimum Operations to Reduce X to Zero|[Go]({{< relref "/ChapterFour/1658.Minimum-Operations-to-Reduce-X-to-Zero.md" >}})|Medium||||33.6%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterOne/Segment_Tree/">下一页➡️</a></p>
</div>
