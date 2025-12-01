# PS-dn5
Paket redovalnica omogoča upravljanje z ocenami študentov.
Nudi funkcionalnosti za dodajanje ocen, izpis redovalnice in določanje končnega uspeha študentov.

Primer uporabe:
```go
studenti := map[string]Student{
    "123": {Ime: "Ana", Priimek: "Kovač"},
}
redovalnica.DodajOceno(studenti, "123", 9, 1, 10)
redovalnica.IzpisRedovalnice(studenti)
redovalnica.IzpisiKoncniUspeh(studenti, 3)
```