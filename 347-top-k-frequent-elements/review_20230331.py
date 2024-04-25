import heapq


def topKFrequent_20230331(nums, k):
    m = {}
    for num in nums:
        m[num] = m.get(num, 0) + 1

    pq = []
    for num, count in m.items():
        heapq.heappush(pq, (count, num))
        if len(pq) > k:
            heapq.heappop(pq)

    res = []
    while pq:
        _, num = heapq.heappop(pq)
        res.append(num)

    return res