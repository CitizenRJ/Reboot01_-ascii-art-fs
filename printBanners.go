package asciiArtFs

import "fmt"

// Print the full outcome
func PrintBanners(banners, arr []string) {
	num := 0
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
			fmt.Println()
			continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				fmt.Print(arr[int(n)+i])

			}
				fmt.Println()

		}
		fmt.Println()

	}
}
