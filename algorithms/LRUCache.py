class Counter:
    "Infinite counter. Starting from start, increasing += 1."

    def __init__(self, start: int=0) -> None:
        self.i = start

    def __iter__(self):
        return self

    def __next__(self) -> int:
        i = self.i
        self.i += 1
        return i

class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.time = {}
        self.cache = {}
        self.last_time = iter(Counter(0))

    def remove(self, key: int) -> None:
        try:
            del self.cache[key]
            del self.time[key]
        except KeyError:
            return

    def update_val(self, key: int, value: int) -> None:
        self.cache[key] = value
        self.time[key] = next(self.last_time)

    def get(self, key: int) -> int:
        try:
            res = self.cache[key]
        except KeyError:
            return -1
        self.time[key] = next(self.last_time)
        return res
        
    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            self.update_val(key, value)
            return
        if len(self.cache) == self.capacity:
            self.remove_oldest()
            self.update_val(key, value)
        
    def remove_oldest(self) -> None:
        if len(self.time) == 0:
            return
        self.remove(min(self.time, key=self.time.get))

    def __str__(self) -> str:
        res = ""
        for k in self.cache:
            res += f"{k}:{self.cache[k]}({self.time[k]})\n"
        return res


cache = LRUCache(2);
cache.put(2, 6)
cache.get(2)
print(cache)

cache.get(1)
cache.put(1, 5)
cache.put(1, 2)
print(cache)

cache.get(1)
cache.get(2)
print(cache)