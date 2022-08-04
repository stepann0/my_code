# nums1 = [4, 6, 3, 7]
# nums2 = [1, 5, 6, 3]
nums1 = [1,2,2,1]
nums2 = [2,2]

def intersection(nums1, nums2):
    res = []
    if len(nums1) >= len(nums2):
        for i in nums1:
            if i in nums2:
                res.append(i)
    return res
print(intersection(nums1, nums2))