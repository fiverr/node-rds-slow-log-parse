package internal

var patterns = []string{
	`User@Host: (\w*)\[\w*] @ ([\w\d.-]*)\s?\[([\d.]*)]`,
	`Id: (\d*) #`,
	`Query_time: ([\d.]*)`,
	`Lock_time: ([\d.]*)`,
	`Rows_sent: (\d*)`,
	`Rows_examined: (\d*)`,
	`(?:use )?(\w*)?;?`,
	`SET timestamp=(\d*);`,
	`(.*);`,
}
