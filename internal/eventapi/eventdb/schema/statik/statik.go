// Code generated by statik. DO NOT EDIT.

package statik

import (
	"github.com/rakyll/statik/fs"
)

const EventapiSql = "eventapi/sql" // static asset namespace

func init() {
	data := "PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x16\x00	\x00001_initial_schema.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE jobset\n(\n    id        bigserial PRIMARY KEY,\n    queue     text NOT NULL,\n    jobset    text NOT NULL,\n    created   timestamp DEFAULT NOW(),\n    UNIQUE (queue, jobset)\n);\nCREATE UNIQUE INDEX idx_queue_jobset ON jobset(queue, jobset);\nCREATE INDEX idx_jobset_created ON jobset(created);\n\nCREATE TABLE latest_seqno\n(\n    jobset_id   bigint PRIMARY KEY,\n    seqno       bigint NOT NULL,\n    update_time timestamp\n);\nCREATE INDEX latest_seqno_update_time ON latest_seqno(update_time);\n\nCREATE TABLE event\n(\n    jobset_id   bigint NOT NULL,\n    seqno       bigint NOT NULL,\n    event       bytea NOT NULL,\n    PRIMARY KEY (jobset_id, seqno)\n);PK\x07\x08\xb3\x00:r\x8d\x02\x00\x00\x8d\x02\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\xb3\x00:r\x8d\x02\x00\x00\x8d\x02\x00\x00\x16\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00001_initial_schema.sqlUT\x05\x00\x01\x80Cm8PK\x05\x06\x00\x00\x00\x00\x01\x00\x01\x00M\x00\x00\x00\xda\x02\x00\x00\x00\x00"
	fs.RegisterWithNamespace("eventapi/sql", data)
}
