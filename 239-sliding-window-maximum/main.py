import heapq


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        n = len(nums)
        pq = []
        for i in range(k):
            pq.append((-nums[i], i))
        heapq.heapify(pq)
        ans = []
        ans.append(-pq[0][0])
        for i in range(k, n):
            heapq.heappush(pq, (-nums[i], i))
            while pq[0][1] <= i - k:
                heapq.heappop(pq)
            ans.append(-pq[0][0])
        return ans
