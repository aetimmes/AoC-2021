from os import posix_fadvise
from aocd import data

NINF = -99999999999

def main(input): 
    x_start, x_end, y_start, y_end = parse(input)
    result = solve_y(y_start, y_end)
    print(result)
    return result
    
def parse(input):
    tokens = input.split(" ")
    x_tokens = tokens[2].split("..")
    x_start = int(x_tokens[0].split("=")[1])
    x_end = int(x_tokens[1].split(",")[0])
    y_tokens = tokens[3].split("..")
    y_start = int(y_tokens[0].split("=")[1])
    y_end = int(y_tokens[1].split(",")[0])
    return x_start, x_end, y_start, y_end
     
def solve_y(y_start, y_end):
    result = 0 
    for i_v in range(0, 1000):
        m = sim_y(i_v, y_start, y_end)
        result = max(m, result)
        i_v += 1
    return result


def sim_y(i_v, y_start, y_end):
    result = 0
    m = 0
    t = 0
    position = 0
    velocity = i_v
    while position > y_end :
        t+=1
        position += velocity
        velocity -= 1
        m = max(m, position)
        if y_start <= position <= y_end:
            result = m
    return result


if __name__ == "__main__":
    main(data)
