// Paket main ki sprejme ukaze iz ukazne vrstice in tako upravlja z redovalnico ki se nahaja v paketu redovalnica
package main

import (
	"GO_5/redovalnica"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v3"
)

func main() {
	// Test podatki ki se vnesejo v redovalnico s pomocjo metod
	redovalnica.StOcen = 3
	redovalnica.MinOcena = 5
	redovalnica.MaxOcena = 10
	redovalnica.DodajStudent("1111", "K", "B")
	redovalnica.DodajStudent("2222", "V", "G")
	redovalnica.DodajStudent("3333", "Z", "S")
	redovalnica.DodajOceno("1111", 8)
	redovalnica.DodajOceno("1111", 10)
	redovalnica.DodajOceno("1111", 9)
	redovalnica.DodajOceno("2222", 6)
	redovalnica.DodajOceno("2222", 7)
	redovalnica.DodajOceno("2222", 6)
	redovalnica.DodajOceno("3333", 5)
	redovalnica.DodajOceno("3333", 6)
	redovalnica.DodajOceno("3333", 5)

	// Komanda
	cmd := &cli.Command{
		// Informacije o programu
		Name:  "Redovalnica",
		Usage: "Redovalnica ocen, dodajanje studentov, dodajanje ocen, izpis ocen, izpis povprecja",
		// Stikala
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "stOcen",
				Value:       3,
				Usage:       "Min stevilo ocen ki jih student potrebuje",
				Destination: &redovalnica.StOcen,
			},
			&cli.IntFlag{
				Name:        "minOcena",
				Value:       5,
				Usage:       "Min ocena ki jo lahko pridobi student",
				Destination: &redovalnica.MinOcena,
			},
			&cli.IntFlag{
				Name:        "maxOcena",
				Value:       10,
				Usage:       "Max ocena ki jo lahko pridobi student",
				Destination: &redovalnica.MaxOcena,
			},
		},
		// Komande ki jih lahko klicemo
		Commands: []*cli.Command{
			{
				Name:    "DodajStudent",
				Aliases: []string{"DS"},
				Usage:   "Doda studenta v redovalnico, args req* (id* ime* priimek*)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					redovalnica.DodajStudent(cmd.Args().First(), cmd.Args().Get(1), cmd.Args().Get(2))
					fmt.Println("Dodan Student", cmd.Args().First(), cmd.Args().Get(1), cmd.Args().Get(2))
					return nil
				},
			},
			{
				Name:    "DodajOceno",
				Aliases: []string{"DO"},
				Usage:   "Doda dolocenemu studentu oceno, args req* (id* ocena*)",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					num, _ := strconv.Atoi(cmd.Args().Get(1))
					redovalnica.DodajOceno(cmd.Args().First(), num)
					return nil
				},
			},
			{
				Name:    "IzpisVsehOcen",
				Aliases: []string{"IVO"},
				Usage:   "Izpise ocene vseh v redovalnici",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					redovalnica.IzpisVsehOcen()
					return nil
				},
			},
			{
				Name:    "IzpisiKoncniUspeh",
				Aliases: []string{"IKU"},
				Usage:   "Izpise koncni uspeh vseh studentov",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					redovalnica.IzpisiKoncniUspeh()
					return nil
				},
			},
		},
	}

	// Izpise napako ce pride to napake
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
