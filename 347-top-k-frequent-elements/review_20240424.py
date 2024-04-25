import heapq
def topKFrequent_20230331(nums, k):
    m = {}
    for n in nums:
        m[n] = m.get(n, 0)+1

    pq = []
    for n, count in m.items():
        heapq.heappush(pq, (count, n))
        if len(pq) > k:
            heapq.heappop(pq)

    res = []
    while pq:
        _, n = heapq.heappop(pq)
        res.append(n)

    return res