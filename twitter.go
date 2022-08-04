package main

import "fmt"

// Posts
type Post struct {
	Id   int
	Text string
}

func NewPost(id int, text string) *Post {
	return &Post{
		Id:   id,
		Text: text,
	}
}

func (p Post) String() string {
	return fmt.Sprintf("post_%v: %s", p.Id, p.Text)
}

// User
type User struct {
	Id        int
	Followers []*User
	Posts     []*Post
}

func NewUser(id int) *User {
	return &User{
		Id:        id,
		Followers: []*User{},
		Posts:     []*Post{},
	}
}

func (u User) String() string {
	followers := make([]string, len(u.Followers))
	for i, f := range u.Followers {
		followers[i] = fmt.Sprintf("user_%v", f.Id)
	}
	return fmt.Sprintf("User_%v:\nfollowers: %v\nposts: %v\n\n", u.Id, followers, u.Posts)
}

func (u *User) HasFollower(id int) bool {
	for _, f := range u.Followers {
		if f.Id == id {
			return true
		}
	}
	return false
}

// Server
type Twitter struct {
	Users      []*User
	FreeUserId int
	FreePostId int
}

func NewTwitter() Twitter {
	return Twitter{
		Users:      []*User{},
		FreeUserId: 0,
		FreePostId: 0,
	}
}

func (this *Twitter) UUserId() int {
	uuid := this.FreeUserId
	this.FreeUserId++
	return uuid
}

func (this *Twitter) AddUser() {
	// if this.IsUniqId(id) == false {
	// 	fmt.Println("Twitter.AddUser(): id isn't uniq.")
	// 	return
	// }
	new_user := &User{
		Id:        this.UUserId(),
		Followers: []*User{},
		Posts:     []*Post{},
	}
	this.Users = append(this.Users, new_user)
}

func (this *Twitter) IsUniqId(id int) bool {
	for i := 0; i < len(this.Users); i++ {
		if this.Users[i].Id == id {
			return false
		}
	}
	return true
}

func (this *Twitter) GetUserById(id int) *User {
	for _, u := range this.Users {
		if u.Id == id {
			return u
		}
	}
	return nil
}

func (this *Twitter) UPostId() int {
	upid := this.FreePostId
	this.FreePostId++
	return upid
}

func (this *Twitter) PostTweet(userId int, tweetText string) {
	user := this.GetUserById(userId)
	if user == nil {
		fmt.Printf("Twitter.PostTweet(): cant find user_%v.\n", userId)
		return
	}
	post := NewPost(this.UPostId(), tweetText)
	user.Posts = append(user.Posts, post)
}

func (this *Twitter) GetNewsFeed(userId int) []*Post {
	user := this.GetUserById(userId)
	if user == nil {
		return []*Post{}
	}
	if len(user.Posts) < 10 {
		return user.Posts
	}
	return user.Posts[:10]
}

// followerId --> followeeId
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		fmt.Println("Twitter.Follow(): followerId and followeeId can't be the same.")
		return
	}
	followee := this.GetUserById(followeeId)
	follower := this.GetUserById(followerId)
	if followee == nil || follower == nil {
		fmt.Println("Twitter.Follow(): cant find user.")
		return
	}
	if !followee.HasFollower(followerId) {
		followee.Followers = append(followee.Followers, follower)
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followee := this.GetUserById(followeeId)
	if followee == nil {
		return
	}
	for i, follower := range followee.Followers {
		if follower.Id == followerId {
			followee.Followers = append(followee.Followers[:i], followee.Followers[i+1:]...)
		}
	}
}

// Utils
func has[T any](arr []T, val T, equal func(a, b T) bool) bool {
	for i := 0; i < len(arr); i++ {
		if equal(arr[i], val) {
			return true
		}
	}
	return false
}

func main() {
	twitter := NewTwitter()

	twitter.AddUser()
	twitter.AddUser()
	twitter.AddUser()

	twitter.PostTweet(0, "Hi!")
	twitter.PostTweet(0, "My name is Peter.")
	twitter.Follow(0, 1)
	twitter.Follow(0, 2)
	fmt.Println(twitter.GetUserById(0))

	twitter.PostTweet(1, "My name is Sonya.")
	twitter.PostTweet(1, "I'm a programmer.")
	twitter.Follow(1, 0)
	twitter.Follow(1, 2)
	fmt.Println(twitter.GetUserById(1))

	twitter.PostTweet(2, "My name is Alex.")
	twitter.PostTweet(2, "I'm a driver.")
	twitter.Follow(2, 0)
	fmt.Println(twitter.GetUserById(2))

	fmt.Println("unfollow")
	twitter.Unfollow(2, 1)
	fmt.Println(twitter.GetUserById(1))
	twitter.Unfollow(2, 0)
	fmt.Println(twitter.GetUserById(0))
}
