package notification

import (
	"model"
	"time"
	"testing"
	"gopkg.in/mgo.v2/bson"
	"github.com/stretchr/testify/assert"
	"github.com/satori/go.uuid"
)

type (
	notifTestCase struct {
		positive	bool
		input		notifInput
		expected	notifExpected
	}

	notifInput struct {
		in			interface{}
	}

	notifExpected struct {
		expect		interface{}
	}
)

func TestFlushNotif(t *testing.T) {
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	const timestampForm = "Jan 2, 2006 at 3:04pm (MST)"
	ts := time.Time{}
	ns := []model.Notification{}

	ts, _ = time.Parse(timestampForm, "Nov 2, 2017 at 3:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: ts, Type: model.FollowType, Detail: model.FollowNotif{Followee: "testFlushNotif", Follower: "AAA"}})
	ts, _ = time.Parse(timestampForm, "Nov 3, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: ts, Type: model.FollowType, Detail: model.FollowNotif{Followee: "testFlushNotif", Follower: "BBB"}})
	ts, _ = time.Parse(timestampForm, "Dec 25, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: ts, Type: model.FollowType, Detail: model.FollowNotif{Followee: "testFlushNotif", Follower: "CCC"}})

	c := client{username: "testFlushNotif", incoming: make(chan model.Notification)}

	// test success
	tc := notifTestCase{true, notifInput{&c}, notifExpected{ns}}

	if tc.positive {
		go c.testFlushNotifTool(t, tc.expected.expect.([]model.Notification))
		h.Manager.FlushNotif(tc.input.in.(*client))
	}
	
	h.Shutdown()
}

func (c *client) testFlushNotifTool(t *testing.T, ns []model.Notification) {
	for _, n := range ns {
		r := <- c.incoming
		assert.Equal(t, n.Type, r.Type, c.username)
		assert.Equal(t, n.Detail, r.Detail, c.username)
	}
	close(c.incoming)
}

func TestForwardNewTweetNotif(t *testing.T) {
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	c1 := client{username: "testForwardNewTweetNotif_1", incoming: make(chan model.Notification)}
	h.Manager.clients[c1.username] = make(map[uuid.UUID]*client)
	h.Manager.clients[c1.username][c1.id] = &c1
	c2 := client{username: "testForwardNewTweetNotif_2", incoming: make(chan model.Notification)}
	h.Manager.clients[c2.username] = make(map[uuid.UUID]*client)
	h.Manager.clients[c2.username][c2.id] = &c2

	// test success
	tc := notifTestCase{true, notifInput{model.Notification{Type: model.NewTweetType, Detail: model.NewTweetNotif{Publisher: "XXX"}}}, 
		notifExpected{model.Notification{Type: model.NewTweetType, Detail: model.NewTweetNotif{Publisher: "XXX"}}}}

	if tc.positive {
		go c1.testForwardNewTweetNotifTool(t, tc.expected.expect.(model.Notification))
		go c2.testForwardNewTweetNotifTool(t, tc.expected.expect.(model.Notification))
		h.Manager.forwardNewTweetNotif(tc.input.in.(model.Notification))
	}
	
	h.Manager.db.Close()
}

func (c *client) testForwardNewTweetNotifTool(t *testing.T, n model.Notification) {
	r := <- c.incoming
	assert.Equal(t, n.Type, r.Type, c.username)
	assert.Equal(t, n.Detail, r.Detail, c.username)
	close(c.incoming)
}

func TestForwardFollowNotif(t *testing.T) {
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	c := client{username: "testForwardFollowNotif", incoming: make(chan model.Notification)}
	h.Manager.clients[c.username] = make(map[uuid.UUID]*client)
	h.Manager.clients[c.username][c.id] = &c

	// test success
	tc := notifTestCase{true, notifInput{model.Notification{Type: model.FollowType, Detail: model.FollowNotif{Followee: "testForwardFollowNotif", Follower: "XXX"}}}, 
		notifExpected{model.Notification{Type: model.FollowType, Detail: model.FollowNotif{Followee: "testForwardFollowNotif", Follower: "XXX"}}}}

	if tc.positive {
		go c.testForwardFollowNotifTool(t, tc.expected.expect.(model.Notification))
		h.Manager.forwardFollowNotif(tc.input.in.(model.Notification))
	}
	
	h.Manager.db.Close()
}

func (c *client) testForwardFollowNotifTool(t *testing.T, n model.Notification) {
	r := <- c.incoming
	assert.Equal(t, n.Type, r.Type, c.username)
	assert.Equal(t, n.Detail, r.Detail, c.username)
	close(c.incoming)
}

func TestClearNotif(t *testing.T) {
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	// test success
	tc := notifTestCase{true, notifInput{&client{username: "testClearNotif"}}, notifExpected{[]model.Notification{}}}

	h.Manager.ClearNotif(tc.input.in.(*client).username)
	if tc.positive {
		r := model.Individual{}
		h.Manager.db.DB(h.Manager.dbName).C(model.NotificationCollection).Find(bson.M{"username": tc.input.in.(*client).username}).One(&r)
		assert.Equal(t, tc.expected.expect.([]model.Notification), r.Notifications, tc.input.in.(*client).username)
	}

	const timestampForm = "Jan 2, 2006 at 3:04pm (MST)"
	ts := time.Time{}
	ns := []model.Notification{}

	ts, _ = time.Parse(timestampForm, "Nov 2, 2017 at 3:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: ts, Type: model.FollowType, Detail: model.FollowNotif{}})
	ts, _ = time.Parse(timestampForm, "Nov 3, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: ts, Type: model.FollowType, Detail: model.FollowNotif{}})
	ts, _ = time.Parse(timestampForm, "Dec 25, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: ts, Type: model.FollowType, Detail: model.FollowNotif{}})

	for _, n := range ns {
		h.Manager.db.DB(h.Manager.dbName).C(model.NotificationCollection).Update(bson.M{"username": tc.input.in.(*client).username}, bson.M{"$addToSet": bson.M{"notifications": n}})
	}
	
	h.Shutdown()
}
