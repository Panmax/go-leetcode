from heapq import heappush, heappop


def topKFrequent(nums, k):
    m = {}
    for num in nums:
        m[num] = m.get(num, 0) + 1
    pq = []
    for key, count in m.items():
        heappush(pq, (count, key))
        if len(pq) > k:
            heappop(pq)
    res = []
    while pq and len(res) < k:
        _, item = heappop(pq)
        res.append(item)
    return res

