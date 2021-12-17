from typing import DefaultDict
from aocd import data

def main(input): 
    x_start, x_end, y_start, y_end= parse(input)
    x_possibilities = solve_x(x_start, x_end)
    x_max = min(x_possibilities["INF"])
    y_possibilities = solve_y(y_start, y_end)
    result = set()
    for t in y_possibilities:
        if t <= x_max:
            for x in x_possibilities[t]:
                for y in y_possibilities[t]:
                    result.add((x,y,))
        else:
            for x in x_possibilities["INF"]:
                for y in y_possibilities[t]:
                    result.add((x,y,))
    print(len(result))

def parse(input):
    tokens = input.split(" ")
    x_tokens = tokens[2].split("..")
    x_start = int(x_tokens[0].split("=")[1])
    x_end = int(x_tokens[1].split(",")[0])
    y_tokens = tokens[3].split("..")
    y_start = int(y_tokens[0].split("=")[1])
    y_end = int(y_tokens[1].split(",")[0])
    return x_start, x_end, y_start, y_end

def solve_x(x_start, x_end):
    result = DefaultDict(list)
    for i in range(x_end+1):
        time = 0
        position = 0
        velocity = i
        while velocity != 0 and position <= x_end:
            time += 1
            position += velocity
            if velocity > 0:
                velocity -= 1
            else:
                velocity += 1
            if x_start <= position <= x_end:
                result[time].append(i)
                if velocity == 0:
                    result["INF"].append(i)
    return result

def solve_y(y_start, y_end):
    result = DefaultDict(list)
    i_v = -1000
    while i_v <= 100:
        ts = sim_y(i_v, y_start, y_end)
        if len(ts) > 0:
            for t in ts:
                result[t].append(i_v)
        i_v+=1
    return result

def sim_y(i_v, y_start, y_end):
    result = []
    t = 0
    position = 0
    velocity = i_v
    while position >= y_start :
        t+=1
        position += velocity
        velocity -= 1
        if y_start <= position <= y_end:
            result.append(t)
    return result

if __name__ == "__main__":
    main(data)
