package leetcode

func canConstruct(ransomNote string, magazine string) bool {
    hash := make([]int, 256)
    for i :=0;i<len(magazine);i++ {
        hash[int(magazine[i])]++
    }

    for i :=0;i<len(ransomNote);i++ {
        if hash[int(ransomNote[i])] <1  {
            return false
        } else {
            hash[int(ransomNote[i])]--
        }
    }
    return true
}