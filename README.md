# Soal no 3

Pada soal no 3 kompleksitas adalah O(N), dimana n adalah panjang dari input dalam kasus ini adalah panjang dari suatu string bracket. Lalu pada kode program harus mengecek sampai ke akhir dari suatu string sehingga harus dilakukan looping dari awal string sampai ke n dari string
```go
for i := 0; i <= len(listOfBrackets)-1; i++ {}
```
Dalam looping operasi yang paling sering dilakukan adalah operasi perbandingan dan operasi assignment yaitu untuk melakukan pengecekan apakah bracket memiliki pasangan yang sesuai atau tidak dan dilakukan pengecekan apakah bracket input sesuai dengan aturan yang ada atau tidak.
```go
	if listOfBrackets[i] != ")" && listOfBrackets[i] != "]" && listOfBrackets[i] != "}" && listOfBrackets[i] != "(" && listOfBrackets[i] != "[" && listOfBrackets[i] != "{" {
		return "NO"
	}

	if listOfBrackets[i] == ")" && stacks[0] == "(" {
		stacks = append([]string{}, stacks[1:]...)
	} else if listOfBrackets[i] == "]" && stacks[0] == "[" {
		stacks = append([]string{}, stacks[1:]...)
	} else if listOfBrackets[i] == "}" && stacks[0] == "{" {
		stacks = append([]string{}, stacks[1:]...)
	} else {
		stacks = append([]string{listOfBrackets[i]}, stacks...)
	}
```
Operasi bersifat konstan yaitu 1 kali dijalankan saja yaitu O(1) namun, secara keseluruhan loop berjalan n kali maka kompleksitas tetap O(n) sehingga kompleksitas dari algoritma ini adalah linier time atau O(n)
