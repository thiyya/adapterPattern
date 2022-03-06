package main

type IkidenUceCeviriciAdaptor struct {
	IkiGirisliPriz *IkiGirisliPriz
}

func (ikidenUceCeviriciAdaptor *IkidenUceCeviriciAdaptor) UcGirisliPrizeTak(ucGirisliPriz *UcGirisliPriz) {
	ucGirisliPriz.takiliAletler = append(ucGirisliPriz.takiliAletler, ikidenUceCeviriciAdaptor.IkiGirisliPriz.takiliAletler...)
}
