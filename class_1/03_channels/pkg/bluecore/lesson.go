package bluecore

import "fmt"

func Lesson(cfg Config) error {
	fmt.Println("Lesson 03 Channels")

	times := 3
	data := make([]ChanData, times)
	for i := 0; i < times; i++ {
		rem := times - i

		fmt.Printf("scanning %d of %d times\n", rem, times)
		if _, err := fmt.Scanf("%s %d %s", &data[i].A, &data[i].B, &data[i].C); err != nil {
			return fmt.Errorf("scan error: %v", err)
		}
	}
	return nil
}

func ReceiveData(cd ChanData) {
	cd.Print()
}

type ChanData struct {
	A string
	B int
	C []byte
}

func (c ChanData) Print() {
	fmt.Printf("Data: A(%s) B(%d) C(%s)\n", c.A, c.B, c.C)
}

type CData chan ChanData

// Reference JSON
// ../../data/journey_campaign.json
