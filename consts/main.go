package consts

const (
	NodeBits  uint8 = 10                    // 节点 ID 的位数
	StepBits  uint8 = 12                    // 序列号的位数
	NodeMax   int64 = -1 ^ (-1 << NodeBits) // 节点 ID 的最大值，用于检测溢出
	StepMax   int64 = -1 ^ (-1 << StepBits) // 序列号的最大值，用于检测溢出
	TimeShift uint8 = NodeBits + StepBits   // 时间戳向左的偏移量
	NodeShift uint8 = StepBits              // 节点 ID 向左的偏移量
	Epoch     int64 = 1288834974657         // 时间 2006
)
