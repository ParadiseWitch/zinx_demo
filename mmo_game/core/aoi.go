package core

import "fmt"

type AOIManager struct {
	MinX  int           //区域左边界坐标
	MaxX  int           //区域右边界坐标
	CntsX int           //x方向格子的数量
	MinY  int           //区域上边界坐标
	MaxY  int           //区域下边界坐标
	CntsY int           //y方向的格子数量
	grids map[int]*Grid //当前区域中都有哪些格子，key=格子ID， value=格子对象
}

func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
	}
	//给AOI初始化区域中所有的格子
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			//计算格子ID
			//格子编号：id = idy *nx + idx  (利用格子坐标得到格子编号)
			gid := y*cntsX + x

			//初始化一个格子放在AOI中的map里，key是当前格子的ID
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridLength(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridLength())
		}
	}
	return aoiMgr
}

func (aoiMgr *AOIManager) gridWidth() int {
	return (aoiMgr.MaxX - aoiMgr.MinX) / aoiMgr.CntsX
}

func (aoiMgr *AOIManager) gridLength() int {
	return (aoiMgr.MaxY - aoiMgr.MinY) / aoiMgr.CntsY
}

//根据格子的gID得到当前周边的九宫格信息
func (aoiMgr *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gID是否存在
	if _, ok := aoiMgr.grids[gID]; !ok {
		return
	}
	//将当前gid添加到九宫格中
	grids = append(grids, aoiMgr.grids[gID])

	//根据gid得到当前格子所在的X轴编号
	idx := gID % aoiMgr.CntsX

	//判断当前idx左边是否还有格子
	if idx > 0 {
		grids = append(grids, aoiMgr.grids[gID-1])
	}
	//判断当前的idx右边是否还有格子
	if idx < aoiMgr.CntsX-1 {
		grids = append(grids, aoiMgr.grids[gID+1])
	}

	//将x轴当前的格子都取出，进行遍历，再分别得到每个格子的上下是否有格子

	//得到当前x轴的格子id集合
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	//遍历x轴格子
	for _, v := range gidsX {
		//计算该格子处于第几列
		idy := v / aoiMgr.CntsX

		//判断当前的idy上边是否还有格子
		if idy > 0 {
			grids = append(grids, aoiMgr.grids[v-aoiMgr.CntsX])
		}
		//判断当前的idy下边是否还有格子
		if idy < aoiMgr.CntsY-1 {
			grids = append(grids, aoiMgr.grids[v+aoiMgr.CntsX])
		}
	}

	return
}

//通过横纵坐标获取对应的格子ID
func (aoiMgr *AOIManager) GetGIDByPos(x, y float32) int {
	gx := (int(x) - aoiMgr.MinX) / aoiMgr.gridWidth()
	gy := (int(x) - aoiMgr.MinY) / aoiMgr.gridLength()

	return gy*aoiMgr.CntsX + gx
}

//通过横纵坐标得到周边九宫格内的全部PlayerIDs
func (aoiMgr *AOIManager) GetPIDsByPos(x, y float32) (playerIDs []int) {
	//根据坐标得到对应的格子ID
	gid := aoiMgr.GetGIDByPos(x, y)

	//根据格子ID得到周边九宫格内的所有格子
	grids := aoiMgr.GetSurroundGridsByGid(gid)

	//遍历所有的格子，得到所有的PlayerIDs
	for _, v := range grids {
		playerIDs = append(playerIDs, v.GetPlyerIDs()...)
		fmt.Printf("===> grid ID : %d, pids : %v  ====", v.GID, v.GetPlyerIDs())
	}

	return
}

func (aoiMgr *AOIManager) String() string {
	s := fmt.Sprintf("AOIManagr:\nminX:%d, maxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\n Grids in AOI Manager:\n",
		aoiMgr.MinX, aoiMgr.MaxX, aoiMgr.CntsX, aoiMgr.MinY, aoiMgr.MaxY, aoiMgr.CntsY)
	for _, grid := range aoiMgr.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}
