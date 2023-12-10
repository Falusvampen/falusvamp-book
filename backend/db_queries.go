package main

import "log"

func statementsCreation() {
	for key, query := range map[string]string{
		"addUser":            `INSERT INTO users (email, password, first_name, last_name, dob, avatar, nickname, about_me, privacy) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		"getAllUsers":        `SELECT id, email, first_name, last_name, nickname from users;`,
		"getUserProfile":     `SELECT id, email, first_name, last_name, dob, avatar, nickname, about_me, privacy from users WHERE id = ?;`,
		"getUserbyID":        `SELECT email, first_name, last_name, nickname FROM users WHERE id = ?;`,
		"getUserIDByEmail":   `SELECT id FROM users WHERE email = ?;`,
		"getUserPrivacy":     `SELECT privacy FROM users WHERE id = ?;`,
		"getUserCredentials": `SELECT email, password FROM users WHERE email = ?;`,
		"updateUserPrivacy":  `UPDATE users SET privacy = ? WHERE id = ?;`,
		"getEmailByID":       `SELECT email FROM users WHERE id = ?;`,

		"addSession":    `INSERT INTO session (uuid, user_id) VALUES (?, ?);`,
		"getSession":    `SELECT * FROM session WHERE uuid = ?;`,
		"getIDbyUUID":   `SELECT id FROM session INNER JOIN users ON users.id=user_id WHERE uuid = ?;`,
		"getIDbyEmail":  `SELECT id FROM users WHERE email = ?;`,
		"removeSession": `DELETE FROM session WHERE uuid = ?;`,

		"addAlmostPrivate": `INSERT INTO almost_private (user_id, post_id) VALUES (?, ?);`,
		"getAlmostPrivate": `SELECT post_id FROM almost_private WHERE user_id = ? AND post_id = ?;`,

		"addPost":          `INSERT INTO post (user_id, title, content, privacy, picture, created_at) VALUES (?, ?, ?, ?, ?, ?);`,
		"getPostsByUserID": `SELECT post.id, title, content, picture, first_name, last_name, post.privacy, email, created_at FROM post INNER JOIN users ON user_id=? ORDER BY created_at DESC;`,
		"addComment":       `INSERT INTO comment (user_id, post_id, content, picture, created_at) VALUES (?, ?, ?, ?, ?);`,
		"getComments":      `SELECT first_name, last_name, content, picture FROM comment INNER JOIN users ON user_id = users.id WHERE post_id = ? ORDER BY comment.id DESC;`,
		"addMessage":       `INSERT INTO message (from_id, to_id, content, created_at) VALUES (?, ?, ?, ?);`,
		"getMessages":      `SELECT from_id, to_id, content, created_at FROM message WHERE (from_id = ? AND to_id = ?) OR (from_id = ? AND to_id = ?) ORDER BY created_at DESC;`,

		"addGroup":  `INSERT INTO groups (name, description, creator, created_at, privacy) VALUES (?, ?, ?, ?, ?);`,
		"getGroups": `SELECT id, name, description, creator, created_at, privacy FROM groups ORDER BY created_at DESC;`,
		"getGroup":  `SELECT id, name, description, creator, created_at, privacy FROM groups WHERE id = ?;`,

		"addGroupMember":           `INSERT INTO group_members (group_id, member_id) VALUES (?, ?);`,
		"getGroupMembers":          `SELECT member_id FROM group_members WHERE group_id = ?;`,
		"getGroupMembersInfo":      `SELECT nickname, first_name, last_name FROM users WHERE id = ?;`,
		"getGroupPendingMembers":   `SELECT member_id FROM group_pending_members WHERE group_id = ?;`,
		"addGroupPendingMember":    `INSERT INTO group_pending_members (group_id, member_id) VALUES (?, ?);`,
		"removeGroupPendingMember": `DELETE FROM group_pending_members WHERE group_id = ? AND member_id = ?;`,
		"removeGroupMember":        `DELETE FROM group_members WHERE group_id = ? AND member_id = ?;`,

		"addGroupInvitedUser":    `INSERT INTO group_invited_users (user_id, group_id, inviter_id, created_at) VALUES (?, ?, ?, ?);`,
		"getGroupInvitedUsers":   `SELECT user_id, group_id, inviter_id, created_at FROM group_invited_users WHERE group_id = ?;`,
		"removeGroupInvitedUser": `DELETE FROM group_invited_users WHERE user_id = ? AND group_id = ?;`,

		"addGroupPost":           `INSERT INTO group_post (user_id, title, content, picture, created_at) VALUES (?, ?, ?, ?, ?);`,
		"addGroupPostMembership": `INSERT INTO group_post_membership (group_id, group_post_id) VALUES (?, ?);`,
		"getGroupPosts":          `SELECT group_post.id, title, content, first_name, last_name, email, created_at, picture FROM group_post JOIN group_post_membership ON group_post.id = group_post_membership.group_post_id JOIN users ON group_post.user_id = users.id ORDER BY created_at DESC;`,

		"addGroupComment":  `INSERT INTO group_comment (user_id, group_post_id, content, picture, created_at) VALUES (?, ?, ?, ?, ?);`,
		"getGroupComments": `SELECT email, first_name, last_name, content, picture, created_at FROM group_comment INNER JOIN users ON users.id = user_id WHERE group_post_id = ? ORDER BY group_comment.id DESC;`,

		"addEvent":                  `INSERT INTO events (group_id, event_name, event_description, event_date, event_location, creator_id, created_at) VALUES (?, ?, ?, ?, ?, ?, ?);`,
		"addEventParticipant":       `INSERT INTO event_participants (event_id, user_id, status, status_updated_at) VALUES (?, ?, ?, ?);`,
		"getEvents":                 `SELECT event_id, group_id, event_name, event_description, event_date, event_location, creator_id, created_at FROM events WHERE group_id = ? ORDER BY created_at DESC;`,
		"getEvent":                  `SELECT event_id, group_id, event_name, event_description, event_date, event_location, creator_id, created_at FROM events WHERE event_id = ? LIMIT 1;`,
		"getEventParticipants":      `SELECT user_id, first_name, last_name, email, avatar, status, status_updated_at FROM event_participants INNER JOIN users ON users.id = user_id WHERE event_id = ? ORDER BY status_updated_at DESC;`,
		"updateEventParticipant":    `UPDATE event_participants SET status = ?, status_updated_at = ? WHERE event_id = ? AND user_id = ?;`,
		"getEventParticipantStatus": `SELECT status FROM event_participants WHERE event_id = ? AND user_id = ? LIMIT 1;`,
		"getUserIDwithEventCount":   `SELECT COUNT(*) FROM event_participants WHERE event_id = ? AND user_id = ?;`,

		"getFollowers":          `SELECT follower_id FROM followers WHERE user_id = ?;`,
		"getFollowersPending":   `SELECT follower_id FROM followers_pending WHERE user_id = ?;`,
		"addFollower":           `INSERT INTO followers (user_id, follower_id) VALUES (?, ?);`,
		"addFollowerPending":    `INSERT INTO followers_pending (user_id, follower_id) VALUES (?, ?);`,
		"removeFollower":        `DELETE FROM followers WHERE user_id = ? AND follower_id = ?;`,
		"removeFollowerPending": `DELETE FROM followers_pending WHERE user_id = ? AND follower_id = ?;`,
		"getFollowing":          `SELECT user_id FROM followers WHERE follower_id = ?;`,
		"doesSecondFollowFirst": `SELECT * FROM followers WHERE user_id = ? AND follower_id = ? LIMIT 1;`,
	} {
		err := error(nil)
		statements[key], err = db.Prepare(query)
		if err != nil {
			log.Print("Error preparing query: " + key)
			log.Fatal(err.Error())
		}
	}
}
