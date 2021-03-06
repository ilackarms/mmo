package main

import (
	"github.com/faiface/pixel"
	"time"
)

type (
	update struct {
		notifyPlayerMoved        *notifyPlayerMoved
		notifyPlayerSpoke        *notifyPlayerSpoke
		notifyWorldState         *notifyWorldState
		notifyPlayerDisconnected *notifyPlayerDisconnected
	}

	notifyPlayerMoved struct {
		id          string
		newPosition pixel.Vec
		requestTime time.Time
	}

	notifyPlayerSpoke struct {
		id   string
		text string
	}

	notifyWorldState struct {
		targetID string
	}

	notifyPlayerDisconnected struct {
		id string
	}
)
