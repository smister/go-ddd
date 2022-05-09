package vars

import (
	"github.com/jeffcail/jsnowflake"
	"github.com/smister/go-ddd/demo2/common/pkg/db"
	"github.com/smister/go-ddd/demo2/common/pkg/event"
)

var DatabaseSetting *db.DatabaseSettingS
var Snowflake *jsnowflake.Machine
var EventPublisher *event.Event
