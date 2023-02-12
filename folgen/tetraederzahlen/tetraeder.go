package tetraederzahlen

import "github.com/reinerhuechting-courses/go.aufgaben-rekursion/folgen/dreieckszahlen"

// Tetraeder berechnet für ein gegebenes n die n-te Tetraederzahl.
// Für die konkreten Zahlenwerte: siehe Tests.
func Tetraeder(n int) int {
	if n == 1 {
		return 1
	}
	return dreieckszahlen.Dreieck(n) + Tetraeder(n-1)
}

// Anmerkung:
// Diese Zahlen sind die Fortsetzung der Dreieckszahlen für die dritte Dimenstion.
// Statt ein Dreieck zu legen, baut man einen Tetraeder aus Kugeln.
// Die Tetraederzahlen geben an, wie viele Kugeln man jeweils braucht.
