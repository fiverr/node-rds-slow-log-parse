package internal

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatterns(t *testing.T) {
	var (
		userhost      = regexp.MustCompile(patterns[0])
		id            = regexp.MustCompile(patterns[1])
		duration      = regexp.MustCompile(patterns[2])
		lock_wait     = regexp.MustCompile(patterns[3])
		rows_sent     = regexp.MustCompile(patterns[4])
		rows_examined = regexp.MustCompile(patterns[5])
		context       = regexp.MustCompile(patterns[6])
		timestamp     = regexp.MustCompile(patterns[7])
		query         = regexp.MustCompile(patterns[8])
	)

	t.Run("userhost", func(t *testing.T) {
		t.Run("Should extract username and hostname", func(t *testing.T) {
			result := userhost.FindStringSubmatch(`User@Host: admin[wheel] @ localhost [127.0.0.1]`)
			assert.Equal(t, "admin", result[1])
			assert.Equal(t, "localhost", result[2])
			assert.Equal(t, "127.0.0.1", result[3])
		})
		t.Run("Should extract username and hostname", func(t *testing.T) {
			result := userhost.FindStringSubmatch(`User@Host: admin[wheel] @ [127.0.0.1]`)
			assert.Equal(t, "admin", result[1])
			assert.Equal(t, "", result[2])
			assert.Equal(t, "127.0.0.1", result[3])
		})
		t.Run("Should extract username and hostname", func(t *testing.T) {
			result := userhost.FindStringSubmatch(`User@Host: admin2[wheel] @ some-domain.com []`)
			assert.Equal(t, "admin2", result[1])
			assert.Equal(t, "some-domain.com", result[2])
			assert.Equal(t, "", result[3])
		})
	})

	t.Run("id", func(t *testing.T) {
		t.Run("Should extract id", func(t *testing.T) {
			result := id.FindStringSubmatch(`Id: 123456 #`)
			assert.Equal(t, "123456", result[1])
		})
	})
	t.Run("duration", func(t *testing.T) {
		t.Run("Should extract duration", func(t *testing.T) {
			result := duration.FindStringSubmatch(`Query_time: 0.000202`)
			assert.Equal(t, "0.000202", result[1])
		})
	})
	t.Run("lock_wait", func(t *testing.T) {
		t.Run("Should extract lock_wait", func(t *testing.T) {
			result := lock_wait.FindStringSubmatch(`Lock_time: 0.000077`)
			assert.Equal(t, "0.000077", result[1])
		})
	})
	t.Run("rows_sent", func(t *testing.T) {
		t.Run("Should extract rows_sent", func(t *testing.T) {
			result := rows_sent.FindStringSubmatch(`Rows_sent: 12`)
			assert.Equal(t, "12", result[1])
		})
	})
	t.Run("rows_examined", func(t *testing.T) {
		t.Run("Should extract rows_examined", func(t *testing.T) {
			result := rows_examined.FindStringSubmatch(`Rows_examined: 12`)
			assert.Equal(t, "12", result[1])
		})
	})
	t.Run("context", func(t *testing.T) {
		t.Run("Should extract SQL context (table name)", func(t *testing.T) {
			result := context.FindStringSubmatch(`use submiterr12_production;`)
			assert.Equal(t, "submiterr12_production", result[1])
		})
	})
	t.Run("timestamp", func(t *testing.T) {
		t.Run("Should extract timestamp", func(t *testing.T) {
			result := timestamp.FindStringSubmatch(`SET timestamp=1551260760;`)
			assert.Equal(t, "1551260760", result[1])
		})
	})

	t.Run("query", func(t *testing.T) {
		t.Run("Should extract query with single quotes", func(t *testing.T) {
			result := query.FindStringSubmatch(`SELECT count(*) from mysql.rds_replication_status WHERE action = 'reset slave' and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1;`)
			assert.Equal(t, "SELECT count(*) from mysql.rds_replication_status WHERE action = 'reset slave' and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1", result[1])
		})
		t.Run("Should extract query with double quotes", func(t *testing.T) {
			result := query.FindStringSubmatch(`SELECT count(*) from mysql.rds_replication_status WHERE action = "reset slave" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1;`)
			assert.Equal(t, "SELECT count(*) from mysql.rds_replication_status WHERE action = \"reset slave\" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1", result[1])
		})
	})
}
