package main

import "fmt"

//instruksi: menampilkan angka 1-100, untuk angka hasil yang x3 tulis "Fizz"
//untuk x5 "Buzz" lalu jika kedua nya maka tulis "FizzBuzz"

// answer: untuk masalah ini kita bisa memanfaatkan sisa bagi agar mengetahui
// apakah suatu angka merupakan kelipatan dari yang kita mau (x3 atau x5)
// Dalam masalah ini, kita ingin memeriksa apakah suatu angka
// habis dibagi 3 atau 5. Operasi sisa bagi (%) memberikan sisa
// dari pembagian dua bilangan bulat. Jika sisa pembagian suatu angka dengan
// 3 atau 5 adalah 0, maka angka tersebut habis dibagi 3 atau 5.
func main() {

	//menggunakan perulangan dari 1-100
	for number := 1; number <= 100; number++ {
		//kemudian kita harus mencari angka yang habis dibagi 3 & 5 terlebih dahulu,
		//karena sifat pengkondisian itu sendiri. Jika kita mendahulukan sisa bagi 3 atau 5
		//maka angka yang habis di bagi 3 & 5 akan masuk ke salah satu pengkodisian tersebut, dan kita
		//tidak mau itu terjadi. Oleh karena itu kita perlu menaruhnya diawal.
		if number%3 == 0 && number%5 == 0 {
			fmt.Println("FizzBuzz")

			//jika angka tidak habis dibagi 3 & 5 maka kita cek apakah angka habis dibagi 3?
			//jika iya maka tulis "Fizz", Jika tidak maka lanjut ke else if selanjutnya
		} else if number%3 == 0 {
			fmt.Println("Fizz")

			//Apakah angka habis dibagi 5?
			//jika iya maka "Buzz"
		} else if number%5 == 0 {
			fmt.Println("Buzz")
			//lalu, jika semua kondisi ternyata tidak terpenuhi
			// maka tampilkan angka nya secara langsung.
		} else {
			fmt.Println(number)
		}

	}
}
