import numpy as np

lines = []
with open("input.txt", "r") as f:
    for line in f:
        lines.append(line.strip())
# lines = [l for l in lines if l != '']
seeds = [int(s) for s in lines[0].split()[1:]]

maps = []
for line in lines[2:]:
    if 'map' in line:
        maps.append([])
    elif line != '':
        maps[-1].append([int(l) for l in line.split()])

def resolve_map(map, index):
    for entry in map:
        if index >= entry[1] and index < (entry[1] + entry[2]):
            return index - entry[1] + entry[0]
    return index 

def resolve_map_reverse(map, index):
    for entry in map:
        if index >= entry[0] and index < (entry[0] + entry[2]):
            return index - entry[0] + entry[1]
    return index

result = []
for seed in seeds:
    for map in maps:
        seed = resolve_map(map, seed)
    result.append(seed)
print(f"Part 1: {np.min(result)}")


from tqdm import trange

maps.reverse()
for i in trange(0, 100000000, 1000):
    seed = i
    for map in maps:
        seed = resolve_map_reverse(map, seed)
    for seeed, rang in np.array(seeds).reshape((-1, 2)):
        if seed >= seeed and seed < seeed + rang:
            iter_1 = i
            print(f"Part 2: {i} iteration 1")
            break
    else: continue
    break

for i in trange(iter_1 - 1000, iter_1 + 1):
    seed = i
    for map in maps:
        seed = resolve_map_reverse(map, seed)
    for seeed, rang in np.array(seeds).reshape((-1, 2)):
        if seed >= seeed and seed < seeed + rang:
            print(f"Part 2: {i}")
            break
    else: continue
    break
