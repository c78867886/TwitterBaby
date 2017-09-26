# **package model**
    import "model"

### TYPES

```go
* type Tweet struct {
    ID bson.ObjectId `json:"id" bson:"_id"`
    //Owner         bson.ObjectId   `json:"owner" bson:"owner"`
    Owner bson.ObjectId `json:"owner,omitempty" bson:"owner,omitempty"`
    //From      bson.ObjectId   `json:"from,omitempty" bson:"from,omitempty"`
    From      string    `json:"from,omitempty" bson:"from,omitempty"`
    Message   string    `json:"message" bson:"message"`
    Timestamp time.Time `json:"timestamp" bson:"timstamp"`
}
```
> Tweet : Data structure that holds a single tweet.

```go
* type User struct {
    ID        bson.ObjectId `json:"id" bson:"_id"`
    FirstName string        `json:"firstname" bson:"firstname"`
    LastName  string        `json:"lastname,omitempty" bson:"lastname,omitempty"`
    Password  string        `json:"password,omitempty" bson:"password"`
    Email     string        `json:"email" bson:"email"`
    Followers []string      `json:"followers,omitempty" bson:"followers,omitempty"`
    Followed  []string      `json:"followed,omitempty" bson:"followed,omitempty"`
    Bio       string        `json:"bio,omitempty" bson:"bio,omitempty"`
    Token     string        `json:"token,omitempty" bson:"-"`

    UserIDdev string `json:"useriddev,omitempty" bson:"useriddev,omitempty"`
}
```
> User : Data structure that holds a single user.

# **package handler**
    import "handler"

### TYPES

```go
* type Handler struct {
    DB *mgo.Session
}
```
> Handler : Data structure that holds a handler for a session.

### Functions

```go
* func (h *Handler) DeleteTweet(c echo.Context) (err error)
```
> DeleteTweet : Delete a specific tweet.

```go
* func (h *Handler) FetchOwnTweets(c echo.Context) (err error)
```
> FetchOwnTweets : Handle requests asking for a list of tweets posted by a specific user, and respond with that list along with some user info.

 ```go
* func (h *Handler) FetchUserInfo(c echo.Context) (err error)
```

```go
* func (h *Handler) NewTweet(c echo.Context) (err error)
```
> NewTweet : Add one tweet for a specific user.