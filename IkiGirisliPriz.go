package main

import "fmt"

type IkiGirisliPriz struct {
	takiliAletler []interface{}
}

func (ikiGirisliPriz *IkiGirisliPriz) IkiGirisliPriziCalistir(alet IkiGiriseUygunAlet) {
	fmt.Println("alet iki girişli prize takılıyor !")
	alet.IkiGirisliPrizeTak(ikiGirisliPriz)
}

func (ikiGirisliPriz *IkiGirisliPriz) TakiliAletleriGetir() {
	fmt.Printf("2 girişlide Takili Aletler : %s\n", ikiGirisliPriz.takiliAletler)
}
