package model

type Ranking struct {
	Games []GameResults `json:"games"`
}

func (r *Ranking) AddGame(gm Game) {
	gmRst := gm.ToResults()
	gmRst.Number = uint32(len(r.Games) + 1)
	r.Games = append(r.Games, gmRst)
}
