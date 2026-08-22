func wordBreak(s string, wordDict []string) bool {
    memo := make(map[int]bool)
    memo[len(s)] = true

    var dfs func(int) bool
    dfs = func(i int) bool {
        if val, found := memo[i]; found {
            return val
        }

        for _, w := range wordDict {
            if i+len(w) <= len(s) && s[i:i+len(w)] == w {
                if dfs(i + len(w)) {
                    memo[i] = true
                    return true
                }
            }
        }

        memo[i] = false
        return false
    }

    return dfs(0)
}