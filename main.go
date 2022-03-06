package main

func main() {

	// amerikan tarzı 3 girişli prizi olustur.
	// bu prizde nelerin takılı oldugunu goster.
	ucGirisliPriz := UcGirisliPriz{}
	ucGirisliPriz.TakiliAletleriGetir()

	// avrupa tarzı 2 girişli prizi olustur.
	ikiGirisliPriz := IkiGirisliPriz{}

	// amerikan tarzı 3 girişli prize uygun bir buzdolabi olustur.
	// avrupa tarzı 2 girişli prize uygun bir tost makinası olustur.
	buzdolabi := Buzdolabi{
		prizSekli: 3,
	}
	tostMakinasi := TostMakinasi{
		prizSekli: 2,
	}

	// buzdolabinin uc girisli priz için gerekli metodu oldugundan sorun yok.
	ucGirisliPriz.UcGirisliPriziCalistir(&buzdolabi)
	//Ancak tost makinasının uc girisli priz için gerekli metodu bulunmuyor. Bu nedenle alttaki kodu calıstıramazsın!
	//ucGirisliPriz.UcGirisliPriziCalistir(&tostMakinasi)

	// cözum olarak bir adaptor olustur.
	ikidenUceCeviriciAdaptor := IkidenUceCeviriciAdaptor{
		IkiGirisliPriz: &ikiGirisliPriz,
	}

	// iki girisli prize tost makinasını tak.
	ikiGirisliPriz.IkiGirisliPriziCalistir(&tostMakinasi)

	// uc girisli prize adaptoru takabilirsin cunku onun uygun metodu var.
	// Bu metodun ıcınde tost makinası fise takılınca ne olacaksa o kısımları cagırabılırsın.
	ucGirisliPriz.UcGirisliPriziCalistir(&ikidenUceCeviriciAdaptor)

	ucGirisliPriz.TakiliAletleriGetir()
	ikiGirisliPriz.TakiliAletleriGetir()
}
