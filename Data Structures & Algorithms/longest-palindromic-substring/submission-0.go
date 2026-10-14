func longestPalindrome(s string) string {
    n := len(s)
	dp := make([][]bool,n)
	for i:=0; i<n; i++{
		dp[i]=make([]bool, n)
	}
	resIdx, resLen := 0, 0
	for i := n-1; i>=0; i--{
		for j := i; j<n; j++{
			if s[i] == s[j] && (j-i<=2 || dp[i+1][j-1]){
				dp[i][j]=true
				if j-i+1>resLen{
					resLen = j-i+1
					resIdx = i
				}
			}
		}
	}
	return s[resIdx:resIdx+resLen]

}
