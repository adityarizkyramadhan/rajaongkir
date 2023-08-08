# Proyek Sedang Berlangsung: RajaOngkir Socket Golang

## Deskripsi Proyek
Proyek ini berfokus pada pembuatan library rajaongkir dalam bahasa go dalam rangka mencapai penggunaan API rajaongkir yang simpel pada bahasa pemrograman golang.

## Tim Proyek
- **Developer**: Aditya Rizky Ramadhan

## Tanggal
- **Mulai Proyek**: Sabtu, 5 Agustus 2023

## Tujuan Utama
- [x] Tujuan 1: Account Starter
- [ ] Tujuan 2: Account Basic
- [ ] Tujuan 3: Account Pro

## Tugas yang Sedang Dikerjakan
### [Account Starter]
- [x] Langkah 1: Mendapatkan data provinsi
- [x] Langkah 2: Mendapatkan data kota
- [x] Langkah 3: Menghitung biaya pengiriman dari kurir tertentu

### [Account Basic]
- [ ] Langkah 1: Mendapatkan data provinsi
- [ ] Langkah 2: Mendapatkan data kota
- [ ] Langkah 3: Menghitung biaya pengiriman dari kurir tertentu

### [Account Pro]
- [ ] Langkah 1: Mendapatkan data provinsi
- [ ] Langkah 2: Mendapatkan data kota
- [ ] Langkah 3: Mendapatkan data kecamatan
- [ ] Langkah 4: Menghitung biaya pengiriman dari kurir tertentu
- [ ] Langkah 5: Melacak status dan lokasi paket

## Kontak
Untuk informasi lebih lanjut, silakan hubungi Aditya Rizky Ramadhan [adityarizky1020@gmail.com]

## Example

### Account Starter
```go
package main

import (
	"fmt"
	"os"

	"github.com/adityarizkyramadhan/rajaongkir"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err.Error())
		return
	}
	roStarter := rajaongkir.New(os.Getenv("RAJA_ONGKIR_KEY")).Starter()
	provinces, err := roStarter.GetProvince()
	fmt.Println("Province : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(provinces)
	cities, err := roStarter.GetCityByIDProvince(provinces[10].ProvinceID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("City : ")
	fmt.Println(cities)

	costs, err := roStarter.GetCost(rajaongkir.CostInput{
		Origin:      "179",
		Destination: "343",
		Weight:      1700,
		Courier:     "jne",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(costs)

}
```
_Sidoarjo, 06 Agustus 2023_
