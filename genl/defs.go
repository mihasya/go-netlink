package genl

/*
#include <linux/taskstats.h>
*/
import "C"

type TaskStats C.struct_taskstats

// Constants
const (
	TASKSTATS_VERSION                      = C.TASKSTATS_VERSION                     
	TS_COMM_LEN                            = C.TS_COMM_LEN                           
	TASKSTATS_CMD_UNSPEC                   = C.TASKSTATS_CMD_UNSPEC                  
	TASKSTATS_CMD_GET                      = C.TASKSTATS_CMD_GET                     
	TASKSTATS_CMD_NEW                      = C.TASKSTATS_CMD_NEW                     
	__TASKSTATS_CMD_MAX                    = C.__TASKSTATS_CMD_MAX                   
	TASKSTATS_TYPE_UNSPEC                  = C.TASKSTATS_TYPE_UNSPEC                 
	TASKSTATS_TYPE_PID                     = C.TASKSTATS_TYPE_PID                    
	TASKSTATS_TYPE_TGID                    = C.TASKSTATS_TYPE_TGID                   
	TASKSTATS_TYPE_STATS                   = C.TASKSTATS_TYPE_STATS                  
	TASKSTATS_TYPE_AGGR_PID                = C.TASKSTATS_TYPE_AGGR_PID               
	TASKSTATS_TYPE_AGGR_TGID               = C.TASKSTATS_TYPE_AGGR_TGID              
	TASKSTATS_TYPE_NULL                    = C.TASKSTATS_TYPE_NULL                   
	__TASKSTATS_TYPE_MAX                   = C.__TASKSTATS_TYPE_MAX                  
	TASKSTATS_CMD_ATTR_UNSPEC              = C.TASKSTATS_CMD_ATTR_UNSPEC             
	TASKSTATS_CMD_ATTR_PID                 = C.TASKSTATS_CMD_ATTR_PID                
	TASKSTATS_CMD_ATTR_TGID                = C.TASKSTATS_CMD_ATTR_TGID               
	TASKSTATS_CMD_ATTR_REGISTER_CPUMASK    = C.TASKSTATS_CMD_ATTR_REGISTER_CPUMASK   
	TASKSTATS_CMD_ATTR_DEREGISTER_CPUMASK  = C.TASKSTATS_CMD_ATTR_DEREGISTER_CPUMASK 
	__TASKSTATS_CMD_ATTR_MAX               = C.__TASKSTATS_CMD_ATTR_MAX              
)


