from heapq import heappush, heappop


def topKFrequent_20220925(nums, k):
    m = {}
    for num in nums:
        m[num] = m.get(num, 0) + 1
    pq = []
    for num, count in m.items():
        heappush(pq, (count, num))
        if len(pq) > k:
            heappop(pq)
    res = []
    while pq:
        _, num = heappop(pq)
        res.append(num)
    return res
