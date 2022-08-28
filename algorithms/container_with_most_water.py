
def most_water(height: list[int]) -> int:
    l = 0
    r = len(height)-1
    max_area = 0
    while l < r:
        area = min(height[l], height[r]) * (r-l)
        if area > max_area:
            max_area = area
        if height[r] > height[l]:
            l += 1
        else:
            r -= 1
    return max_area

height = [3, 6, 2, 9, 6, 7]
print(most_water(height))