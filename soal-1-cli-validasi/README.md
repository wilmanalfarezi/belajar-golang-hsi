#soal 1 cli validasi
Program sederhana berbasis CLI yang meminta input
Nama dan Umur lalu memvalidasi apakah umur yang dimasukan minimal 18 tahun 

pastikan sudah tehubung ke remote git berikut : 
git remote https://github.com/wilmanalfarezi/belajar-golang-hsi.git

##cara menjalankan
go run main.go

## Input

Program akan meminta dua input:
- Nama (teks)
- Umur (angka bulat)

Contoh:
Masukan nama anda: Wilman
Masukan umur anda: 21

## Output

Jika umur valid (≥ 18):

Nama    : Wilman
Umur    : 21
>> Selamat datang, Wilman!


Jika umur < 18:

Nama    : Wilman
Umur    : 16
>> Error: umur tidak valid (minimal 18 tahun)

Jika input umur bukan angka:

Masukan umur anda: abc
Masukan angka!

## Catatan
- Pastikan input umur berupa angka agar tidak terjadi error.