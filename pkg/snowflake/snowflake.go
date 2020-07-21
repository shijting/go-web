package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

type Obj struct {
	StartTime time.Time
	MachineID int64
	Node      *snowflake.Node
}

// startTime
// machineID
func NewSnowflake(st time.Time, machineID int64) (sf *Obj, err error) {
	//var st time.Time
	//st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	var node *snowflake.Node
	node, err = snowflake.NewNode(machineID)
	sf = &Obj{
		StartTime: st,
		MachineID: machineID,
		Node:      node,
	}
	return
}
func (this *Obj) GenID() int64 {
	return this.Node.Generate().Int64()
}
