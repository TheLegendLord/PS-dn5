// Paket redovalnica omogoča upravljanje z ocenami študentov.
// Nudi funkcionalnosti za dodajanje ocen, izračun povprečja ocen, izpis redovalnice in
// določanja končnega uspeha študentov.
//
// Primer uporabe:
//
//   studenti := map[string]Student{
//       "123": {ime: "Ana", priimek: "Kovač"},
//   }
//   redovalnica.DodajOceno(studenti, "123", 9)
//   avg := redovalnica.PovprecjeGetter(studenti, "123")
package redovalnica

import (
    "fmt"
)

// Student predstavlja posameznega študenta z imenom, priimkom in seznamom ocen.
type Student struct {
    Ime     string
    Priimek string
    Ocene   []int
}

// DodajOceno doda oceno (1-10) študentu z dano vpisno številko
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) {
	if ocena < minOcena || ocena > maxOcena {
        fmt.Println("Ocena je izven sprejetega intervala")
        return
    }
    
    student, ok := studenti[vpisnaStevilka]
    if !ok {
        fmt.Println("Študenta ni v seznamu")
        return
    }

    student.Ocene = append(student.Ocene, ocena)
    studenti[vpisnaStevilka] = student
    return
}

// povprecje izračuna povprečno oceno študenta, funkcija ni izvožena(je skrita).
func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
    student, ok := studenti[vpisnaStevilka]
    if !ok {
        fmt.Println("Študenta ni v seznamu")
        return -1.0
    }

    if len(student.Ocene) < stOcen {
        return 0.0
    }

    var povprecjeOcen float64 = 0
    for _, v := range student.Ocene {
        povprecjeOcen += float64(v)
    }
    povprecjeOcen /= float64(len(student.Ocene))

    return povprecjeOcen
}

// IzpisRedovalnice izpiše vse študente v podanem slovarju skupaj z njihovim imenom, priimkom in
// seznamom ocen.
func IzpisRedovalnice(studenti map[string]Student) {
    for key, value := range studenti {
        fmt.Printf("%s - %s %s: %v\n", key, value.Ime, value.Priimek, value.Ocene)
    }
}

// IzpisiKoncniUspeh izpiše končni uspeh vsakega študenta glede na njegovo povprečno oceno.
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int) {
    for key, value := range studenti {
        povprecnaOcena := povprecje(studenti, key, stOcen)

        if povprecnaOcena >= 9 {
            fmt.Printf("%s %s: povprečna ocena %.2f -> Odličen študent!\n", value.Ime, value.Priimek, povprecnaOcena)
        } else if povprecnaOcena >= 6 {
            fmt.Printf("%s %s: povprečna ocena %.2f -> Povprečen študent\n", value.Ime, value.Priimek, povprecnaOcena)
        } else {
            fmt.Printf("%s %s: povprečna ocena %.2f -> Neuspešen študent\n", value.Ime, value.Priimek, povprecnaOcena)
        }
    }
}