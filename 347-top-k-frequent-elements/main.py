from heapq import heappop, heappush


def topKFrequent(nums, k):
    res = []

    map = {}
    for num in nums:
        map[num] = map.get(num, 0) + 1

    pq = []
    for key, freq in map.items():
        heappush(pq, (freq, key))
        if len(pq) > k:
            heappop(pq)

    while pq and len(res) < k:
        _, item = heappop(pq)
        res.append(item)
    return res


if __name__ == '__main__':
    print(topKFrequent([1, 1, 3, 3, 3, 2, 2, 2, 2], 2))
