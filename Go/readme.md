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