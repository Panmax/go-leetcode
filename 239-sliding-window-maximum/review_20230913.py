import heapq


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        n = len(nums)
        qp = [(-nums[i], i) for i in range(k)]
        heapq.heapify(qp)
        ans = [-qp[0][0]]
        for i in range(k, n):
            heapq.heappush(qp, (-nums[i], i))
            while qp[0][1] <= i - k:
                heapq.heappop(qp)
            ans.append(-qp[0][0])

        return ans
