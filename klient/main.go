package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strings"
	"time"

	"github.com/slaraz/turniej/gra_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	IP_ADDR               = "localhost:50051"
	NOWY_MECZ_TIMEOUT     = time.Second * 5
	DOLACZ_DO_GRY_TIMEOUT = time.Second * 1000
	RUCH_GRACZA_TIMEOUT   = time.Second * 1000
)

var (
	addr  = flag.String("addr", IP_ADDR, "adres serwera gry")
	nazwa = flag.String("nazwa", "Ziutek", "nazwa gracza")
	nowa  = flag.Bool("nowa", false, "tworzy nową grę na serwerze")
	graID = flag.String("gra", "", "dołącza do gry o podanym id")
	lg    = flag.Int("lg", 2, "określa liczbę graczy")
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Koniec.")

	flag.Parse()

	// Utowrzenie połączenia z serwerem gry.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial: %v", err)
	}
	defer conn.Close()
	c := proto.NewGraClient(conn)

	conn.GetState()
	// Jeśli podano opcję -nowa, to utwórz nową grę.
	if *nowa {
		ctx, cancel := context.WithTimeout(context.Background(), NOWY_MECZ_TIMEOUT)
		defer cancel()
		nowaGraInfo, err := c.NowyMecz(ctx, &proto.KonfiguracjaGry{LiczbaGraczy: int32(*lg)})
		if err != nil {
			log.Fatalf("c.NowyMecz: %v", err)
		}
		log.Printf("Nowa gra %q\n", nowaGraInfo.GraID)

		*graID = nowaGraInfo.GraID
	}

	// Jeśli nie utworzono -nowa,
	// ani nie podano opcji -gra, to kończymy.
	if *graID == "" {
		flag.Usage()
		return
	}

	var (
		kartyDlaKtorychTrzebaPodacKolor = map[proto.Karta]bool{
			proto.Karta_L1:  true,
			proto.Karta_L2:  true,
			proto.Karta_A1:  true,
			proto.Karta_A1B: true,
		}
		karta proto.Karta
		kolor proto.KolorZolwia
	)

	// przebieg gry

	// dołączamy do gry graID
	stanGry := dolaczDoGry(c, *graID, *nazwa)
	for {
		// wypisuję stan gry na ekranie
		drukujStatus(stanGry)
		if stanGry.CzyKoniec {
			return
		}
		dwaBledneRuchy := 0
		for {
			if dwaBledneRuchy > 1 {
				karta = botLosowy(stanGry)
			} else {
				karta, kolor = botRuch(stanGry)
			}
			// gracz podaje kartę na konsoli

			if _, ok := kartyDlaKtorychTrzebaPodacKolor[karta]; ok {
				if kolor == 0 {
					kolor = botKolor()
				}
			} else {
				kolor = proto.KolorZolwia_XXX
			}

			// wysyłam ruch do serwera
			nowyStan, err := wyslijRuch(c, &proto.RuchGracza{
				GraID:        stanGry.GraID,
				GraczID:      stanGry.GraczID,
				ZagranaKarta: karta,
				KolorWybrany: kolor,
			})
			if err != nil && status.Code(err) == codes.InvalidArgument {
				dwaBledneRuchy++
				// zły ruch
				// fmt.Printf("Błąd ruchu: %v\n", err)
				continue
			} else if err != nil {
				// inny błąd, np. połączenie z serwerem
				log.Fatalf("wyslijRuch: status: %v, err: %v", status.Code(err), err)
			}
			dwaBledneRuchy = 0
			// ruch ok
			stanGry = nowyStan
			break
		}
	}
}

func botLosowy(stanGry *proto.StanGry) proto.Karta {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	n := r.Intn(len(stanGry.TwojeKarty))
	return stanGry.TwojeKarty[n]
}

func botRuch(stanGry *proto.StanGry) (proto.Karta, proto.KolorZolwia) {
	// for _, karta := range stanGry.TwojeKarty {
	// 	kolor, funkcja := getKolorFunkcja(karta)
	// 	if kolor == int32(stanGry.TwojKolor-1) && (funkcja == 0 || funkcja == 1) {
	// 		return karta, 0
	// 	}
	// }

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	n := r.Intn(len(stanGry.TwojeKarty))

	koloryZolwia, inneKolory, ostatnie, dowolne := rozpoznajKarty(stanGry.TwojKolor, stanGry.TwojeKarty)
	polePrzed, poleZolw, polePo := rozponajPola(stanGry)
	czyOstatni, _ := czyNaszZolwOstatni(stanGry)

	zolwiePod := znajdzZolwiePodNaszym(poleZolw, stanGry.TwojKolor)
	zolwieNad := znajdzZolwieNadNaszym(poleZolw, stanGry.TwojKolor)

	for _, zolw := range zolwiePod {
		koloryZolwia, _, _, _ := rozpoznajKarty(zolw, stanGry.TwojeKarty)

		if czyOstatni && len(ostatnie) > 0 {
			return slices.Max(ostatnie), zolw
		}

		if czySaKarty(koloryZolwia) {
			karty := znajdzKarteDoPrzodu(koloryZolwia)
			if len(karty) > 0 {
				return karty[0], 0
			}
		}
		for _, karta := range dowolne {
			if int32(karta) == 18 {
				return karta, zolw
			}
		}
	}
	for _, zolw := range zolwieNad {
		koloryZolwia, _, _, _ := rozpoznajKarty(zolw, stanGry.TwojeKarty)

		if czySaKarty(koloryZolwia) {
			karty := znajdzKarteDoTylu(koloryZolwia)
			if len(karty) > 0 {
				return karty[0], 0
			}
		}
		for _, karta := range dowolne {
			if int32(karta) == 19 {
				return karta, zolw
			}
		}
	}

	if czyOstatni && len(ostatnie) > 0 {
		return slices.Max(ostatnie), stanGry.TwojKolor
	}

	if czySaKarty(koloryZolwia) {
		karty := znajdzKarteDoPrzodu(koloryZolwia)
		if len(karty) > 0 {
			return karty[0], 0
		}
	}
	for _, karta := range dowolne {
		if int32(karta) == 18 {
			return karta, stanGry.TwojKolor
		}
	}

	if czySaKarty(inneKolory) {
		karty := znajdzKarteDoTylu(koloryZolwia)
		if len(karty) > 0 {
			return karty[0], 0
		}
	}
	for _, karta := range dowolne {
		if int32(karta) == 19 {
			if len(polePo) > 0 {
				return karta, polePo[0]
			}
			if len(polePrzed) > 0 {
				return karta, polePrzed[0]
			}
		}
	}

	return stanGry.TwojeKarty[n], 0
}

func czySaKarty(karty []proto.Karta) bool {
	return len(karty) != 0
}

func znajdzKarteDoPrzodu(karty []proto.Karta) []proto.Karta {
	kartyDoPrzodu := []proto.Karta{}
	for _, karta := range karty {
		if czyKartaPoruszaDoPrzodu(karta) {
			kartyDoPrzodu = append(kartyDoPrzodu, karta)
		}
	}
	if len(kartyDoPrzodu) > 0 {
		return []proto.Karta{slices.Max(kartyDoPrzodu)}
	} else {
		return []proto.Karta{}
	}
}

func czyKartaPoruszaDoPrzodu(karta proto.Karta) bool {
	return int32(karta)%3 == 1 || int32(karta)%3 == 2
}

func znajdzKarteDoTylu(karty []proto.Karta) []proto.Karta {
	kartyDoTylu := []proto.Karta{}
	for _, karta := range karty {
		if czyKartaPoruszaDoTylu(karta) {
			kartyDoTylu = append(kartyDoTylu, karta)
		}
	}

	if len(kartyDoTylu) > 0 {
		return []proto.Karta{slices.Max(kartyDoTylu)}
	} else {
		return []proto.Karta{}
	}
}

func czyKartaPoruszaDoTylu(karta proto.Karta) bool {
	return int32(karta)%3 == 0
}

func rozpoznajKarty(kolor proto.KolorZolwia, karty []proto.Karta) (kartyKolor []proto.Karta, inneKolory []proto.Karta, ostatnie []proto.Karta, dowolne []proto.Karta) {
	for _, karta := range karty {
		if int32(karta) > 17 {
			dowolne = append(dowolne, karta)
		} else if int32(karta) > 15 {
			ostatnie = append(ostatnie, karta)
		} else if int32(kolor)-1 == int32(karta)/3 {
			kartyKolor = append(kartyKolor, karta)
		} else {
			inneKolory = append(inneKolory, karta)
		}
	}

	return
}

func rozponajPola(stanGry *proto.StanGry) (polePrzed []proto.KolorZolwia, poleZolw []proto.KolorZolwia, polePo []proto.KolorZolwia) {
	for i, pole := range stanGry.Plansza {
		for _, zolw := range pole.Zolwie {
			if zolw == stanGry.TwojKolor {
				if i > 0 {
					polePrzed = stanGry.Plansza[i-1].Zolwie
				}
				poleZolw = pole.Zolwie
				if i < 9 {
					polePo = stanGry.Plansza[i+1].Zolwie
				}
			}
		}

	}
	return
}

func czyNaszZolwOstatni(stanGry *proto.StanGry) (czyOstatni bool, ostatnie []proto.KolorZolwia) {
	zebrane := make(map[proto.KolorZolwia]bool)
	for _, pole := range stanGry.Plansza {
		for _, zolw := range pole.Zolwie {
			zebrane[zolw] = true
		}
	}
	if len(zebrane) < 5 {
		if _, ok := zebrane[stanGry.TwojKolor]; !ok {
			czyOstatni = true
		}
		if _, ok := zebrane[1]; !ok {
			ostatnie = append(ostatnie, 1)
		}
		if _, ok := zebrane[2]; !ok {
			ostatnie = append(ostatnie, 2)
		}
		if _, ok := zebrane[3]; !ok {
			ostatnie = append(ostatnie, 3)
		}
		if _, ok := zebrane[4]; !ok {
			ostatnie = append(ostatnie, 4)
		}
		if _, ok := zebrane[5]; !ok {
			ostatnie = append(ostatnie, 5)
		}
		return
	}

	for _, pole := range stanGry.Plansza {
		for _, zolw := range pole.Zolwie {
			if zolw == stanGry.TwojKolor {
				czyOstatni = true
			}
			ostatnie = append(ostatnie, zolw)
		}
		if len(ostatnie) > 0 {
			return
		}

	}
	return
}

func znajdzZolwiePodNaszym(pole []proto.KolorZolwia, naszZolw proto.KolorZolwia) (zolwie []proto.KolorZolwia) {
	for _, zolw := range pole {
		if zolw == naszZolw {
			return
		}
		zolwie = append(zolwie, zolw)
	}
	return
}
func znajdzZolwieNadNaszym(pole []proto.KolorZolwia, naszZolw proto.KolorZolwia) (zolwie []proto.KolorZolwia) {
	zacznijDodawac := false
	for _, zolw := range pole {
		if zacznijDodawac {
			zolwie = append(zolwie, zolw)
		}
		if zolw == naszZolw {
			zacznijDodawac = true
		}
	}
	return
}

func botKolor() proto.KolorZolwia {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	n := r.Intn(5)

	return proto.KolorZolwia(n + 1)
}

func getKolorFunkcja(karta proto.Karta) (quotient, remainder int32) {
	quotient = int32(karta-1) / 3 // integer division, decimals are truncated
	remainder = int32(karta-1) % 3
	return
}

func wczytajKolor() proto.KolorZolwia {
	fmt.Print("Wybierz kolor\n> ")
	var kolor string
	_, err := fmt.Scanln(&kolor)
	if err != nil {
		log.Fatalf("Błąd wczytywania koloru: %v", err)
	}
	k, ok := proto.KolorZolwia_value[strings.ToUpper(kolor)]
	if !ok {
		log.Fatalf("Niepoprawny kolor: %q", kolor)
	}
	return proto.KolorZolwia(k)
}

func wczytajKarte() proto.Karta {
	var karta proto.Karta
	for {
		fmt.Print("Wybierz kartę do zagrania:\n> ")
		var kartatxt string
		_, err := fmt.Scanln(&kartatxt)
		if err != nil {
			log.Fatalf("Błąd wczytywania karty: %v", err)
		}
		k, ok := proto.Karta_value[strings.ToUpper(kartatxt)]
		if !ok {
			fmt.Printf("Niepoprawna karta: %q\n", kartatxt)
			continue
		}
		karta = proto.Karta(k)
		break
	}
	return karta
}

func dolaczDoGry(c proto.GraClient, graID, nazwa string) *proto.StanGry {
	log.Printf("Gracz %s dołącza do gry %q", nazwa, graID)
	ctx, cancel := context.WithTimeout(context.Background(), DOLACZ_DO_GRY_TIMEOUT)
	defer cancel()
	log.Println("Czekam na odpowiedź od serwera...")
	stanGry, err := c.DolaczDoGry(ctx, &proto.Dolaczanie{
		GraID:       graID,
		NazwaGracza: nazwa,
	})
	if err != nil {
		log.Fatalf("c.Dolacz: %v", err)
	}
	return stanGry
}

func wyslijRuch(c proto.GraClient, ruch *proto.RuchGracza) (*proto.StanGry, error) {
	log.Printf("Gracz %s-%s zagrywa kartę: %v", ruch.GraID, ruch.GraczID, ruch.ZagranaKarta)
	ctx, cancel := context.WithTimeout(context.Background(), RUCH_GRACZA_TIMEOUT)
	defer cancel()
	// log.Println("Czekam na odpowiedź od serwera (ruch przeciwnika)...")

	return c.MojRuch(ctx, ruch)
}

func drukujStatus(stanGry *proto.StanGry) {
	if stanGry.CzyKoniec {
		fmt.Println("Koniec gry, wygrał gracz nr", stanGry.KtoWygral)
	} else {
		fmt.Printf("Twój kolor: %v, Pola:", stanGry.TwojKolor)
		for _, pole := range stanGry.Plansza {
			fmt.Printf(" %v", pole.Zolwie)
		}
		fmt.Printf(", Twoje karty: %v\n", stanGry.TwojeKarty)
	}
}
