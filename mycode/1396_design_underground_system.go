package mycode

type UndergroundSystemEntryRecord struct {
	StationName string
	Time        int
}

type UndergroundSystemStatics struct {
	Total int
	Count int
}

type UndergroundSystem struct {
	Entries map[int]*UndergroundSystemEntryRecord
	Statics map[string]map[string]*UndergroundSystemStatics
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Entries: make(map[int]*UndergroundSystemEntryRecord),
		Statics: make(map[string]map[string]*UndergroundSystemStatics),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if _, ok := this.Entries[id]; ok {
		return
	} else {
		this.Entries[id] = &UndergroundSystemEntryRecord{stationName, t}
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if _, ok := this.Entries[id]; !ok {
		return
	} else {
		entryRecord := this.Entries[id]
		delete(this.Entries, id)
		from, to := entryRecord.StationName, stationName
		if _, ok1 := this.Statics[from]; !ok1 {
			this.Statics[from] = make(map[string]*UndergroundSystemStatics)
		}
		if _, ok1 := this.Statics[from][to]; !ok1 {
			this.Statics[from][to] = &UndergroundSystemStatics{0, 0}
		}
		this.Statics[from][to].Total += t - entryRecord.Time
		this.Statics[from][to].Count++
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if m, ok := this.Statics[startStation]; ok {
		if s, ok1 := m[endStation]; ok1 {
			return float64(s.Total) / float64(s.Count)
		}
	}
	return 0.0
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
