// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"net/url"
	"sync"

	"github.com/ChimeraCoder/anaconda"
)

// TweetPosterMock is a mock implementation of proc.TweetPoster.
//
//	func TestSomethingThatUsesTweetPoster(t *testing.T) {
//
//		// make and configure a mocked proc.TweetPoster
//		mockedTweetPoster := &TweetPosterMock{
//			PostTweetFunc: func(msg string, v url.Values) (anaconda.Tweet, error) {
//				panic("mock out the PostTweet method")
//			},
//		}
//
//		// use mockedTweetPoster in code that requires proc.TweetPoster
//		// and then make assertions.
//
//	}
type TweetPosterMock struct {
	// PostTweetFunc mocks the PostTweet method.
	PostTweetFunc func(msg string, v url.Values) (anaconda.Tweet, error)

	// calls tracks calls to the methods.
	calls struct {
		// PostTweet holds details about calls to the PostTweet method.
		PostTweet []struct {
			// Msg is the msg argument value.
			Msg string
			// V is the v argument value.
			V url.Values
		}
	}
	lockPostTweet sync.RWMutex
}

// PostTweet calls PostTweetFunc.
func (mock *TweetPosterMock) PostTweet(msg string, v url.Values) (anaconda.Tweet, error) {
	if mock.PostTweetFunc == nil {
		panic("TweetPosterMock.PostTweetFunc: method is nil but TweetPoster.PostTweet was just called")
	}
	callInfo := struct {
		Msg string
		V   url.Values
	}{
		Msg: msg,
		V:   v,
	}
	mock.lockPostTweet.Lock()
	mock.calls.PostTweet = append(mock.calls.PostTweet, callInfo)
	mock.lockPostTweet.Unlock()
	return mock.PostTweetFunc(msg, v)
}

// PostTweetCalls gets all the calls that were made to PostTweet.
// Check the length with:
//
//	len(mockedTweetPoster.PostTweetCalls())
func (mock *TweetPosterMock) PostTweetCalls() []struct {
	Msg string
	V   url.Values
} {
	var calls []struct {
		Msg string
		V   url.Values
	}
	mock.lockPostTweet.RLock()
	calls = mock.calls.PostTweet
	mock.lockPostTweet.RUnlock()
	return calls
}
