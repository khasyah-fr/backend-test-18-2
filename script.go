package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var numPlayers, numDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scanln(&numPlayers)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scanln(&numDice)

	players := make([][]int, numPlayers)
	for i := range players {
		players[i] = make([]int, numDice)
	}
	scores := make([]int, numPlayers)

	// Perulangan setiap giliran lempar dadu sampai hanya tersisa 1 pemain
	for turn := 1; numPlayers > 1; turn++ {
		fmt.Println("==================")
		fmt.Printf("Giliran %d lempar dadu\n", turn)

		// Loop untuk setiap pemain
		for i := range players {
			// Lewati pemain yang sudah tidak memiliki dadu
			if len(players[i]) == 0 {
				continue
			}

			// Roll dadu
			for j := range players[i] {
				players[i][j] = rand.Intn(6) + 1
			}
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, scores[i], players[i])
		}

		fmt.Println("Setelah evaluasi: ")

		// Loop untuk setiap pemain
		for i := range players {
			// Evaluasi dadu
			for j := 0; j < len(players[i]); j++ {
				switch players[i][j] {
				case 1:
					// Berikan dadu angka 1 ke pemain berikutnya
					if i == numPlayers-1 {
						players[(i+1)%numPlayers] = append(players[(i+1)%numPlayers], 1)
					} else {
						players[i+1] = append(players[i+1], 1)
					}
					players[i] = append(players[i][:j], players[i][j+1:]...)
					j--
				case 6:
					// Hilangkan dadu angka 6 dan tambahkan skor
					players[i] = append(players[i][:j], players[i][j+1:]...)
					j--
					scores[i]++
				default:
					// Biarkan dadu
				}
			}
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, scores[i], players[i])
		}

		// Hitung jumlah pemain yang masih memiliki dadu
		numPlayers = 0
		for i := range players {
			if len(players[i]) > 0 {
				numPlayers++
			}
		}
		fmt.Println("==================")
	}

	// Cari pemain dengan skor tertinggi
	var winner int
	for i := range scores {
		if scores[i] > scores[winner] {
			winner = i
		}
	}

	// Cetak hasil
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", winner+1)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki skor lebih banyak dari pemain lainnya.\n", winner+1)
}
