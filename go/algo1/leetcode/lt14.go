package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    minLen := len(strs[0])
    for i :=1;i<len(strs);i++ {
        if len(strs[i])  < minLen {
            minLen = len(strs[i])
        }
    }
    var ans strings.Builder
    for i := 0; i < minLen; i++ {
        r := strs[0][i]
        for j:=1;j<len(strs);j++ {
            if len(strs[j]) < i {
                return ans.String()
            }
            if strs[j][i] != r {
                return ans.String()
            }
        }
        ans.WriteByte(r)
    }
    return ans.String()
}