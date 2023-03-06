/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package notifier

import (
	"errors"
)

// ReleaseMonitor is the concrete Publisher.
type ReleaseMonitor struct {
	observerList []Observer
	current      string
	enabled      *bool
	lastCheck    string
	available    bool
}

// NewReleaseMonitor creates a new ReleaseMonitor object.
func NewReleaseMonitor(enabled *bool, current, lastCheck string) *ReleaseMonitor {
	return &ReleaseMonitor{
		enabled:   enabled,
		current:   current,
		lastCheck: lastCheck,
	}
}

// Notify defines the logic to issue notifications.
func (r *ReleaseMonitor) Notify() (bool, error) {
	if *r.enabled {
		r.available = true
		return r.notifyAll()
	}

	return true, nil
}

// Register attaches the Observer to the ReleaseMonitor.
func (r *ReleaseMonitor) Register(o Observer) (bool, error) {
	for _, observer := range r.observerList {
		if observer == o {
			return false, errors.New("Observer already exists")
		}
	}

	r.observerList = append(r.observerList, o)
	return true, nil
}

// Deregister detaches the Observer to the ReleaseMonitor.
func (r *ReleaseMonitor) Deregister(o Observer) (bool, error) {
	for i, observer := range r.observerList {
		if observer == o {
			r.observerList = append(r.observerList[:i], r.observerList[i+1:]...)
			return true, nil
		}
	}

	return false, errors.New("Observer not found")
}

func (r *ReleaseMonitor) notifyAll() (bool, error) {
	var results []bool
	var res bool
	var err error

	for _, observer := range r.observerList {
		res, err = observer.Update(r.current, r.lastCheck)
		results = append(results, res)
	}

	for _, result := range results {
		if !result {
			return false, err
		}
	}

	return true, err
}
