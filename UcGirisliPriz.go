package main

import "fmt"

type UcGirisliPriz struct {
	takiliAletler []interface{}
}

func (ucGirisliPriz *UcGirisliPriz) UcGirisliPriziCalistir(alet UcGiriseUygunAlet) {
	fmt.Println("alet uc girişli prize takılıyor !")
	alet.UcGirisliPrizeTak(ucGirisliPriz)
}

func (ucGirisliPriz *UcGirisliPriz) TakiliAletleriGetir() {
	fmt.Printf("3 girişlide Takili Aletler : %s\n", ucGirisliPriz.takiliAletler)
}
