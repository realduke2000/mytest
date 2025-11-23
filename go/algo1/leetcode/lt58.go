package leetcode

func lengthOfLastWord(s string) int {
    pos := len(s)-1
    for i:=len(s)-1;i>=0;i-- {
        if s[i] == ' ' {
            pos--
        } else {
            break
        }
    }

    ans := 0
    for i := pos;i>=0;i-- {
        if s[i] != ' ' {
            ans++
        } else {
            break
        }
    }
    return ans
}