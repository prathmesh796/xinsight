package xapi

type UserResponse struct {
	Data User `json:"data"`
}

type User struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Username      string        `json:"username"`
	Description   string        `json:"description"`
	CreatedAt     string        `json:"created_at"`
	PublicMetrics PublicMetrics `json:"public_metrics"`
	Verified      bool          `json:"verified"`
	Protected     bool          `json:"protected"`
}

type PublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}
