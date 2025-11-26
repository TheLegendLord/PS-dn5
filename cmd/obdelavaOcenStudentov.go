package main

import (
	"fmt"
    "os"
    
    "github.com/TheLegendLord/PS-dn5/redovalnica"
    "github.com/urfave/cli/v3"
)

func main() {
    students := map[string]redovalnica.Student{
		"63190000": {Ime: "Anton", Priimek: "Abeldaber", Ocene: []int{7, 9, 10, 9, 10}},
		"63190001": {Ime: "Bine", Priimek: "Baltazaros", Ocene: []int{6, 6, 8, 6, 7, 4}},
		"63190002": {Ime: "Cecilija", Priimek: "Carazaros", Ocene: []int{9, 8, 8, 10, 9}},
	}

    cmd := &cli.App {
        Name: "obdelavaOcenStudentov",
        Usage: "hrani ocene študentam in izračuna njihov končni uspeh",
        Flags: []cli.Flag {
            &cli.IntFlag {
                Name: "stOcen",
                Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
                Value: 6,
            },
            &cli.IntFlag {
                Name: "minOcena",
                Usage: "najmanjša možna ocena",
                Value: 1,
            },
            &cli.IntFlag {
                Name: "maxOcena",
                Usage: "največja možna ocena",
                Value: 10,
            },
        },
        Action: func(c *cli.Context) error {
            stOcen := c.Int("stOcen")
            minOcena := c.Int("minOcena")
            maxOcena := c.Int("maxOcena")

            redovalnica.DodajOceno(slovarStudentov, "63190000", 10, minOcena, maxOcena)
            redovalnica.IzpisRedovalnice(slovarStudentov)
            redovalnica.IzpisiKoncniUspeh(slovarStudentov, stOcen)
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        fmt.Println(err)
    }
}