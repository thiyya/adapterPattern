package main

type Buzdolabi struct{
	prizSekli int
}

func (buzdolabi *Buzdolabi)UcGirisliPrizeTak(ucGirisliPriz *UcGirisliPriz){
	ucGirisliPriz.takiliAletler = append(ucGirisliPriz.takiliAletler, buzdolabi)
}