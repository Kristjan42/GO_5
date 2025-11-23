// Paket redovalnica vsebuje funkcije za interakcijo z redovalnico
//
// Funkcije ki jih vsebuje:
//   - DodajStudent
//   - DodajOceno
//   - povprecje
//   - IzpisVsehOcen
//   - IzpisiKoncniUspeh
package redovalnica

import (
	"fmt"
)

// Min stevilo ocen ki jih potrebuje student
var StOcen int

// Min ocena ki jo lahko student pridobi
var MinOcena int

// MaX ocena ki jo lahko student pridobi
var MaxOcena int

// Map struktov Student
var Studenti = make(map[string]Student)

// Struct Student
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// Funkcija doda vnesenega studenta v redovalnico
func DodajStudent(vpisnaStevilka string, ime string, priimek string) {
	Studenti[vpisnaStevilka] = Student{Ime: ime, Priimek: priimek, Ocene: make([]int, 0, 100)}
}

// Funkcija doda dolocenemu studentu oceno v redovalnico
func DodajOceno(vpisnaStevilka string, ocena int) {
	if MinOcena <= ocena && ocena <= MaxOcena {
		_, ok := Studenti[vpisnaStevilka]
		switch ok {
		case true:
			{
				var copyStudent = Studenti[vpisnaStevilka]
				copyStudent.Ocene = append(copyStudent.Ocene, ocena)
				Studenti[vpisnaStevilka] = copyStudent
			}
		case false:
			fmt.Printf("Student z vpisno stevliko %s ne obstaja\n\n", vpisnaStevilka)
		}
	} else {
		fmt.Printf("Ocena ni v ustreznem obmocju %d-%d, (vnesena ocena: %d )\n\n", MinOcena, MaxOcena, ocena)
	}
}

// Funkcija povprecje izracuna povprecno oceno studenta
func povprecje(vpisnaStevilka string) float64 {
	_, ok := Studenti[vpisnaStevilka]
	switch ok {
	case true:
		{
			length := len(Studenti[vpisnaStevilka].Ocene)
			if length <= 0 {
				return 0.0
			}
			var sum int = 0
			for _, v := range Studenti[vpisnaStevilka].Ocene {
				sum += v
			}
			avg := (float64(sum) / float64(length))
			if avg < float64(MinOcena) || length < StOcen {
				return 0.0
			} else {
				return avg
			}
		}
	case false:
		return -1.0
	}
	return -1.0

}

// Funkcija izpise redovalnico
func IzpisVsehOcen() {
	fmt.Println("REDOVALNICA")
	for k, v := range Studenti {
		fmt.Printf("%s - %s %s: %v\n", k, v.Ime, v.Priimek, v.Ocene)
	}
	fmt.Println()
}

// Funkcija izpise vse studente in njihove povprecne ocene
func IzpisiKoncniUspeh() {
	uspesnost := func(avg float64) string {
		switch {
		case avg < float64(MinOcena)+1:
			return "Neuspesen student"
		case float64(MinOcena) <= avg && avg < float64(MaxOcena)-1:
			return "Povprecen student"
		case avg >= float64(MaxOcena)-1:
			return "Odlicen student!"
		}
		return "err"
	}
	for k, v := range Studenti {
		fmt.Printf("%s %s: povprecna ocena %f -> %s\n", v.Ime, v.Priimek, povprecje(k), uspesnost(povprecje(k)))
	}
	fmt.Println()

}

