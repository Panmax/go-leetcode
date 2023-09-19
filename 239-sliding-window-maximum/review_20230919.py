import heapq


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
       ans = []
       pq = [(-nums[i], i) for i in range(k)]
       heapq.heapify(pq)
       ans.append(-pq[0][0])
       for i in range(k, len(nums)):
           heapq.heappush(pq, (-nums[i],i))
           while pq[0][1] <= i-k:
               heapq.heappop(pq)
           ans.append(-pq[0][0])

       return ans