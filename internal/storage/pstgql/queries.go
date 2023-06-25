package pstgql

import (
	"fmt"
)

// Get all user data with data_types by user id, or only info data
//
// return: id, data_type, meta_data, hash, updated_at, created_at
//
// WithData
// return: id, data, data_type, meta_data, hash, updated_at, created_at
//
// sql args:
// 1 - user_id.
func queryGetUserData(withData bool, dataTypes ...string) string {
	fields := `id,data_type,meta_data,hash,updated_at,created_at`
	query := ""
	if withData {
		fields = `id,data,data_type,meta_data,hash,updated_at,created_at`
	}
	if len(dataTypes) == 0 {
		query = fmt.Sprintf(`--sql
			SELECT %s
			FROM manager_data 
			WHERE user_id=$1
		`, fields)
		return query
	}
	types := "'" + dataTypes[0] + "'"
	for i, t := range dataTypes {
		if i == 0 {
			continue
		}
		types += ",'" + t + "'"
	}
	query = fmt.Sprintf(`--sql
		SELECT %s
		FROM manager_data 
		WHERE data_type in (%s)  and user_id=$1`, fields, types)
	return query
}

// Get user by email
//
// return: id, email, private_key
//
// sql agrs:
// 1 - email.
func queryUserByEmail() string {
	return `--sql
		SELECT id,email,private_key 
		FROM users
		WHERE email=$1
	`
}

// Insert user
// sql agrs:
// 1 - @email.
// 2 - @private_key.
func queryInsertUser() string {
	return `--sql
		insert into users (email, private_key) values (@email, @privateKey) returning id,email,private_key
	`
}
