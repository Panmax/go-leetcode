class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        m = len(nums1)
        n = len(nums2)
        pq = [(nums1[i]+nums2[0], i, 0) for i in range(min(k, m))]
        res = []
        while pq and len(res) < k:
            _, i, j = heappop(pq)
            res.append([nums1[i], nums2[j]])
            if j+1<n:
                heappush(pq, (nums1[i]+nums2[j+1], i, j+1))
        return res