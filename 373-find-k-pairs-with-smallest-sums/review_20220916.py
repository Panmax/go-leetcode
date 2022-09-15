class Solution(object):
    def kSmallestPairs(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: List[List[int]]
        """
        m, n = len(nums1), len(nums2)
        res = []
        pq = [(nums1[i]+nums2[0], i, 0) for i in range(min(k, m))]
        while pq and len(res)<k:
            _, i, j = heappop(pq)
            res.append([nums1[i],nums2[j]])
            if j+1<n:
                heappush(pq, (nums1[i]+nums2[j+1], i, j+1))
        return res