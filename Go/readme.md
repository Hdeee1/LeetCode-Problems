# 3. Two Sum
- Temukan dua angka dalam array nums yang jika dijumlahkan hasilnya sama dengan target.
- Kembalikan indeks dari kedua angka tersebut dalam bentuk array.
- Setiap soal pasti memiliki satu jawaban benar. Satu indeks tidak boleh digunakan dua kali.

## Terjemahan Petunjuk & Tantangan:
- Tantangan: Buat algoritma dengan kompleksitas waktu lebih cepat dari O(n^2).
- Metode brute force menggunakan dua perulangan bersarang akan sangat lambat.
- Logikanya: Jika nilai saat ini adalah x, kamu harus mencari nilai y di mana y = target - x.
- Gunakan ruang penyimpanan tambahan berupa hash map (map di Go) untuk mengingat angka yang sudah dilewati. Ini mempercepat pencarian y menjadi O(1).

### Contoh:
```bash
// Inisialisasi map dengan key = nilai angka, value = indeks
riwayatAngka := make(map[int]int)

x := 2
y := target - x

// Cek apakah nilai 'y' sudah tersimpan di dalam map
indeksY, ditemukan := riwayatAngka[y]
```

# 4. Roman to Integer
Ubah teks simbol Romawi menjadi angka. Aturan utama: jika simbol yang lebih kecil berada tepat di depan simbol yang lebih besar, maka nilainya dikurangi (contoh: IV = 5 - 1). Jika tidak, nilainya ditambahkan.

## Alur Logika:
- Peta Nilai (*Map*): Buat sebuah *map* untuk menyimpan nilai setiap karakter ('I': 1, 'V': 5, dst).
- Variabel Total: Siapkan variabel bernilai 0 untuk menyimpan hasil akhir.
- Perulangan: Lakukan iterasi dari awal hingga akhir indeks teks.
- Validasi Batas: Pastikan indeks saat ini bukan indeks terakhir sebelum membandingkannya dengan indeks berikutnya.
- Pengecekan: Jika nilai karakter saat ini lebih kecil dari nilai karakter di sebelah kanannya, kurangi variabel total dengan nilai karakter saat ini.
- Penambahan: Jika lebih besar atau sama (atau sudah di indeks terakhir), tambahkan nilai karakter saat ini ke variabel total.

```bash
func romanToInt(s string) int {
    // 1. Deklarasi map nilai Romawi
    // 2. Deklarasi variabel total = 0
    // 3. For loop sepanjang len(s)
        // a. Cek apakah indeks+1 ada DAN nilai saat ini < nilai berikutnya
        // b. Jika ya: total -= nilai saat ini
        // c. Jika tidak: total += nilai saat ini
    // 4. Return total
}
```

# 5. Palindrome Number
Tentukan apakah angka x adalah palindrom tanpa mengubahnya menjadi teks (string). Angka negatif pasti bukan palindrom karena tanda minus (-121 dibalik menjadi 121-).

## Alur Logika (Matematis):
- Pengecekan Awal: Jika x negatif, atau x berakhiran angka 0 (kecuali angka 0 itu sendiri), langsung kembalikan false.
- Simpan Nilai Asli: Buat variabel original untuk menduplikasi nilai x, karena nilai x asli akan dipotong hingga habis.
- Inisialisasi Pembalik: Buat variabel reversed bernilai 0 untuk menyusun angka baru dari belakang.
- Perulangan Pembalikan: Selama x > 0:
-- Kalikan reversed dengan 10, lalu tambahkan dengan digit terakhir dari x (x % 10).
-- Buang digit terakhir x dengan pembagian bulat (x /= 10).
- Evaluasi Akhir: Kembalikan hasil perbandingan apakah original == reversed.