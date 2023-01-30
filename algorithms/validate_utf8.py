class Solution:
    def validFollowingByte(self, byte: int) -> bool:
        return byte >= 0b10_000000 and byte <= 0b10_111111

    def validUtf8(self, data: list[int]) -> bool:
        i = 0
        while i < len(data):
            byte = data[i]
            if byte < 0b10000000: # 1-byte char
                i += 1
            elif byte >= 0b110_00000 and byte <= 0b110_11111: # 2-byte char
                if i == len(data) - 1: # can't read next byte
                    return False
                i += 1
                if self.validFollowingByte(data[i]):
                    i += 1
            elif byte >= 0b1110_0000 and byte <= 0b1110_1111: # 3-byte char
                if i + 2 >= len(data): # can't read next 2 bytes
                    return False
                for _ in range(2):
                    i += 1
                    if not self.validFollowingByte(data[i]):
                        return False
                i += 1
            elif byte >= 0b11110_000 and byte <= 0b11110_111: # 4-byte char
                if i + 3 >= len(data): # can't read next 3 bytes
                    return False
                for _ in range(3):
                    i += 1
                    if not self.validFollowingByte(data[i]):
                        return False
                i += 1
            else:
                return False
        return True


s = Solution()
s.validUtf8([72, 101, 108, 108, 111, 44, 32, 228, 184, 150, 231, 149, 140])
