package main

type TostMakinasi struct{
	prizSekli int
}

func (tostMakinasi *TostMakinasi)IkiGirisliPrizeTak(ikiGirisliPriz *IkiGirisliPriz){
	ikiGirisliPriz.takiliAletler = append(ikiGirisliPriz.takiliAletler, tostMakinasi)
}