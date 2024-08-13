package main

import (
	"fmt"
	"os"
)

type Person struct {
	Id        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	argsRaw := os.Args
	if len(argsRaw) > 1 {
		arg := argsRaw[1]
		FindUser(arg)
	} else {
		fmt.Println("Tolong Masukkan Nama Atau Nomor Absen")
		fmt.Println("Contoh : 'go run . reza' atau 'go run . 2'")
	}
}
func FindUser(arg string) {
	person := []Person{
		{Id: "1", Nama: "fahreza", Alamat: "bondowoso", Pekerjaan: "Website Developer", Alasan: "Grinding Ilmuu"},
		{Id: "2", Nama: "maulana", Alamat: "bondowoso", Pekerjaan: "Website Developer", Alasan: "Ngeliat golang beda dari php tentang memory, concurency dimana hal tsb ga ada di php"},
		{Id: "3", Nama: "wisnu", Alamat: "bondowoso", Pekerjaan: "Website Developer", Alasan: "Pengen Ngerasain Bikin project dari golang"},
		{Id: "4", Nama: "wardana", Alamat: "bondowoso", Pekerjaan: "Website Developer", Alasan: "Pengen coba bahasa baru"},
	}
	for _, v := range person {
		if v.Nama == arg || v.Id == arg {
			fmt.Println("ID : ", v.Id)
			fmt.Println("Nama : ", v.Nama)
			fmt.Println("Alamat : ", v.Alamat)
			fmt.Println("Pekerjaan : ", v.Pekerjaan)
			fmt.Println("Alasan : ", v.Alasan)
			return
		}
	}
	fmt.Println("Data dengan nama/absen tsb tidak tersedia")
}
