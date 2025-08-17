package main

func main() {
	for i, a := 1, len(os.Args)-1; a > 0; i, a = i+1, a-1 {
		n := i - 1
		if n == 0 {
			z01.PrintRune('0')
		} else {
			d := []int{}
			for x := n; x > 0; x /= 10 {
				d = append([]int{x % 10}, d...)
			}
			for _, v := range d {
				z01.PrintRune(rune(v + '0'))
			}
		}
		z01.PrintRune(':')
		z01.PrintRune(' ')
		for _, r := range os.Args[len(os.Args)-i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
