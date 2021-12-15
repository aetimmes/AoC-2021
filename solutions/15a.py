import heapq
from aocd import data

dirs = [
    (1,0),
    (0,-1),
    (-1,0),
    (0,1)
]

def main(input): 
    grid = [[int(c) for c in line] for line in input.split()]
    rows, cols = len(grid), len(grid[0])
    best = [[-1 for i in range(cols)] for j in range(rows)]
    heap = []
    # Tuples are risk, x, y
    heapq.heappush(heap, (0, 0, 0,))
    while True:
        r, x, y = heapq.heappop(heap)
        if x == cols-1 and y == rows-1:
            print(r)
            return
        for x1, y1 in dirs:
            x2 = x+x1
            y2 = y+y1
            if bounds_check(x2, y2, rows, cols):
                r2 = r + grid[y2][x2]
                if best[y2][x2] == -1 or r2 < best[y2][x2]:
                    print(f"Pushing ({r2}, {x2}, {y2})") 
                    best[y2][x2] = r2
                    heapq.heappush(heap, (r2,x2,y2))


def bounds_check(x, y, r, c):
    return 0 <= x < c and 0 <= y < r


if __name__ == "__main__":
    main(data)
