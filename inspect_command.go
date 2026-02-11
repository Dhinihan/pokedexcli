package main

import "fmt"

func commandInspect(c config, args []string) error {
	if len(args) < 1 {
		fmt.Println("É preciso informar o pokémon, 'inspect <pokemon>'")
	}

	pName := args[0]
	pokemon, ok := c.pokeDex.Get(pName)
	if !ok {
		fmt.Println("You must catch the Pokémon first!")
		return nil
	}

	fmt.Printf("Name  : %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	stats := pokemon.ParseStats()

	fmt.Println("Stats :")
	fmt.Printf("%14s: %d\n", "Hp", stats.Hp)
	fmt.Printf("%14s: %d\n", "Attack", stats.Attack)
	fmt.Printf("%14s: %d\n", "Deffense", stats.Defense)
	fmt.Printf("%14s: %d\n", "SpecialAttack", stats.SpecialAttack)
	fmt.Printf("%14s: %d\n", "SpecialDefense", stats.SpecialDefense)
	fmt.Printf("%14s: %d\n", "Hp", stats.Hp)

	fmt.Println("Types :")
	for _, ty := range pokemon.Types {
		fmt.Println("- " + ty.Type.Name)
	}

	return nil
}
