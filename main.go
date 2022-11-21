package main

import (
	"fmt"
	"test_qoin/lib"
)

func PrintDice(dices []lib.Dice) {
	for i, dice := range dices {
		fmt.Printf("dice : %d -> ", i+1)

		fmt.Println(dice.GetNumber())
	}
}

func main() {
	var (
		PLAYER_COUNT int
		DICE_COUNT   int
	)

	fmt.Print("Masukkan Jumlah Pemain : ")
	fmt.Scan(&PLAYER_COUNT)
	fmt.Print("Masukkan Jumlah Dadu : ")
	fmt.Scan(&DICE_COUNT)

	fmt.Println("======================")

	var (
		round int = 1
	)

	players := []lib.Player{}

	for i := 0; i < PLAYER_COUNT; i++ {
		players = append(players, lib.NewPlayer(DICE_COUNT))
	}

	for true {
		activePlayer := 0
		activePlayerIndex := -1

		fmt.Printf("\nGiliran %d lempar dadu!\n", round)
		for i, player := range players {
			if player.IsFinished() {
				fmt.Printf("\t Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", i+1, player.GetPoint())
			} else {
				player.Play()
				fmt.Printf("\t Pemain #%d (%d): %s\n", i+1, player.GetPoint(), player.DiceToString())
			}
		}
		fmt.Println("Setelah Evaluasi:")
		// evaluate
		for i, player := range players {
			dices := player.GetDices()

			j := 0
			for _, dice := range dices {

				if dice.GetNumber() == 6 {
					player.AddPoint(1)
					player.RemoveDice(j)
					j--
				} else if dice.GetNumber() == 1 {
					if i+1 == PLAYER_COUNT {
						players[0].AddBonus(dice)
					} else {
						players[i+1].AddBonus(dice)
					}
					player.RemoveDice(j)
					j--
				}
				j++
			}
		}

		for i, player := range players {
			if player.IsFinished() {
				fmt.Printf("\t Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", i+1, player.GetPoint())
			} else {
				player.MergeBonus()
				fmt.Printf("\t Pemain #%d (%d): %s\n", i+1, player.GetPoint(), player.DiceToString())
				activePlayer++
				activePlayerIndex = i
			}
		}

		if activePlayer == 1 {
			fmt.Printf("\nGame berakhir karena hanya pemain #%d yang memiliki dadu\n", activePlayerIndex+1)
			break
		} else if activePlayer == 0 {
			fmt.Printf("\nGame berakhir karena tidak ada pemain yang memiliki dadu\n")
			break
		}
		round++
	}

	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", GetWinnerIndex(&players)+1)

}

func GetWinnerIndex(players *[]lib.Player) int {
	index := -1
	point := 0
	for i, player := range *players {
		if player.GetPoint() > point {
			point = player.GetPoint()
			index = i
		}
	}
	return index
}
