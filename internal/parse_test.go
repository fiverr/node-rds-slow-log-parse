package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseRaw(t *testing.T) {
	type args struct {
		record string
	}
	tests := []struct {
		name string
		args args
		want RawRecord
	}{
		{
			args: args{
				record: `# Time: 190227 9:46:00 # User@Host: rdsadmin[rdsadmin] @ localhost [] Id: 4069049 # Query_time: 0.000202 Lock_time: 0.000077 Rows_sent: 1 Rows_examined: 1 SET timestamp=1551260760; SELECT count(*) from mysql.rds_replication_status WHERE action = "reset slave" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1;`,
			},
			want: RawRecord{
				User:         "rdsadmin",
				Host:         "localhost",
				IP:           "",
				ID:           "4069049",
				Duration:     "0.000202",
				LockWait:     "0.000077",
				RowsSent:     "1",
				RowsExamined: "1",
				Timestamp:    "1551260760",
				Query:        `SELECT count(*) from mysql.rds_replication_status WHERE action = "reset slave" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1`,
			},
		},
		{
			args: args{
				record: "# Time: 190227 9:46:00 # User@Host: rdsadmin[rdsadmin] @ localhost [172.16.0.12] Id: 4069049 # Query_time: 0.000202 Lock_time: 0.000077 Rows_sent: 1 Rows_examined: 1 SET timestamp=1551260760; SELECT count(*) from mysql.rds_replication_status WHERE action = 'reset slave' and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1;",
			},
			want: RawRecord{
				User:         "rdsadmin",
				Host:         "localhost",
				IP:           "172.16.0.12",
				ID:           "4069049",
				Duration:     "0.000202",
				LockWait:     "0.000077",
				RowsSent:     "1",
				RowsExamined: "1",
				Timestamp:    "1551260760",
				Query:        `SELECT count(*) from mysql.rds_replication_status WHERE action = 'reset slave' and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1`,
			},
		},
		{
			args: args{
				record: "# Time: 190418 8:55:34 # User@Host: root[root] @ [172.30.0.229] Id: 19348 # Query_time: 3.000551 Lock_time: 0.000000 Rows_sent: 1 Rows_examined: 0 SET timestamp=1555577734; select sleep(3);",
			},
			want: RawRecord{
				User:         "root",
				Host:         "",
				IP:           "172.30.0.229",
				ID:           "19348",
				Duration:     "3.000551",
				LockWait:     "0.000000",
				RowsSent:     "1",
				RowsExamined: "0",
				Timestamp:    "1555577734",
				Query:        "select sleep(3)",
			},
		},
		{
			args: args{
				record: `# Time: 190418 8:55:34
# User@Host: root[root] @ [172.30.0.229] Id: 19348
# Query_time: 3.000551 Lock_time: 0.000000 Rows_sent: 1 Rows_examined: 0
SET timestamp=1555577734;
select sleep(3);`,
			},
			want: RawRecord{
				User:         "root",
				Host:         "",
				IP:           "172.30.0.229",
				ID:           "19348",
				Duration:     "3.000551",
				LockWait:     "0.000000",
				RowsSent:     "1",
				RowsExamined: "0",
				Timestamp:    "1555577734",
				Query:        "select sleep(3)",
			},
		},
		{
			args: args{
				record: "# Time:  190418 8:55:34\n# User@Host:  root[root] @   [172.30.0.229] Id: 19348       \n# Query_time:  3.000551 Lock_time: 0.000000 Rows_sent: 1 Rows_examined: 0\n\n\nSET timestamp=1555577734;        \nselect sleep(3);",
			},
			want: RawRecord{
				User:         "root",
				Host:         "",
				IP:           "172.30.0.229",
				ID:           "19348",
				Duration:     "3.000551",
				LockWait:     "0.000000",
				RowsSent:     "1",
				RowsExamined: "0",
				Timestamp:    "1555577734",
				Query:        "select sleep(3)",
			},
		},
		{
			args: args{
				record: "2019-04-18T08:55:36.001Z   465aa33d-83d4-4ac3-8198-d94e19993b25   # Time: 190418 12:29:49\n# User@Host: root[root] @ [172.30.0.229] Id: 19348\n# Query_time: 1.099867 Lock_time: 0.000000 Rows_sent: 1 Rows_examined: 0\nSET timestamp=1555590589;\nselect sleep(1.1);",
			},
			want: RawRecord{
				User:         "root",
				Host:         "",
				IP:           "172.30.0.229",
				ID:           "19348",
				Duration:     "1.099867",
				LockWait:     "0.000000",
				RowsSent:     "1",
				RowsExamined: "0",
				Timestamp:    "1555590589",
				Query:        "select sleep(1.1)",
			},
		},
		{
			args: args{
				record: "# User@Host: submitter[submitter] @ [172.16.235.69] Id: 3434185\n# Query_time: 0.003231 Lock_time: 0.000053 Rows_sent: 218 Rows_examined: 218\nuse submiterr_production;\nSET timestamp=1555754404;\nSELECT `jobs`.* FROM `jobs`;",
			},
			want: RawRecord{
				User:         "submitter",
				Host:         "",
				IP:           "172.16.235.69",
				ID:           "3434185",
				Duration:     "0.003231",
				LockWait:     "0.000053",
				RowsSent:     "218",
				RowsExamined: "218",
				Context:      "submiterr_production",
				Timestamp:    "1555754404",
				Query:        "SELECT `jobs`.* FROM `jobs`",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParseRaw(tt.args.record), "ParseRaw(%v)", tt.args.record)
		})
	}
}
