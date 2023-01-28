/* eslint-disable indent */
/* eslint-disable no-shadow */
/* eslint-disable no-console */
/* eslint-disable arrow-parens */
/* eslint-disable max-len */
/* eslint-disable no-unused-vars */
const countries: string[] = [
  'Estonia',
  'Finland',
  'Sweden',
  'Denmark',
  'Norway',
  'IceLand',
];
const names: string[] = ['Asabeneh', 'Mathias', 'Elias', 'Brook'];
const numbers: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// your code here
// 1. Gunakan forEach untuk console.log setiap negara, nama, nomor dalam array.
// answer here
countries.forEach(e => {
  console.log(e);
});

names.forEach(e => {
  console.log(e);
});

numbers.forEach(e => {
  console.log(e);
});

// 2. Gunakan map untuk membuat array baru dengan mengubah setiap negara menjadi hurup besar di array negara.
// answer here

const newArr: string[] = countries.map((ctr) => ctr.toUpperCase());
console.log(newArr);

// 3. Gunakan map untuk membuat array baru dengan mengubah setiap angka menjadi kuadrat dalam array angka.
// answer here

const squares: number[] = numbers.map((num) => num * num);
console.log(squares);

// 4. Gunakan map untuk mengubah setiap nama menjadi hurup besar di array negara.
// answer here
const newName: string[] = names.map((name) => name.toUpperCase());
console.log(newName);

// 5. Gunakan filter untuk menyaring negara yang mengandung land.
// answer here
const countryHasLand: string[] = countries.map((country) => country.toLowerCase()).filter(words => words.includes('land'));
console.log(countryHasLand);

// 6. Gunakan filter untuk menyaring negara yang memiliki enam karakter.
// answer here
const countryHas6Character: string[] = countries.filter((ctr) => ctr.length === 6);
console.log(countryHas6Character);

// 7. Gunakan filter untuk memfilter negara yang berisi enam huruf dan lebih banyak lagi dalam array negara.
// answer here
const countryHas6CharacterOrMore: string[] = countries.filter((ctr) => ctr.length >= 6);
console.log(countryHas6CharacterOrMore);

// 8. Gunakan filter untuk menyaring negara yang dimulai dengan 'E'.
// answer here
const countryWithE: string[] = countries.filter((ctr) => ctr.charAt(0) === 'E');
console.log(countryWithE);

// 9 Rantai (chaining) dua atau lebih array iterator (mis. arr.map(callback).filter(callback).reduce(callback)) dan tampilkan hasilnya dalam console.
// usecase mendapatkan negara yang mulai dari vowel dan jumlah karakter >= 6
// answer here
const countryWithVowel: string[] = countries.filter(ctr => ctr.charAt(0).match(/[aiueoAIUEO]/i)).filter(ctr => ctr.length >= 6);
console.log(countryWithVowel);

// 10. deklarasikan fungsi yang disebut getStringLists yang mengambil array sebagian parameter dan kemudian mengembalikan array hanya dengan item string.
// answer here
const arr = [1, 'Supreme', true, 'pergi'];
const getStringLists = (newArr: (string | number | boolean)[]): string[] => {
  let newStr: string[] = [];
  newArr.forEach(e => {
    if (typeof e === 'string') {
      newStr.push(e);
    }
  }); 
  return newStr;
};
console.log(getStringLists(arr));

// 11. Gunakan reduce untuk menjumlahkan semua angka dalam array angka.
// answer here
const total: number = numbers.reduce((acc, curValue) => acc + curValue);
console.log(total);

// 12. Gunakan reduce untuk menggabungkan semua negara dan untuk menghasilkan kalimat ini: Estonia, Finlandia, Swedia, Denmark, Norwegia, dan IceLand adalah negara-negara Eropa utara.
// answer here
// eslint-disable-next-line arrow-body-style
const joinCountry: string = countries.reduce((arr, text, idx) => {
  if (idx === countries.length - 1) {
    return `${arr}dan ${text} adalah negara-negara Eropa utara.`;
  }
 return `${arr + text}, `;
}, '');
console.log(joinCountry);

// 13. Jelaskan perbedaan some dan every
// SOME: array some digunakan untuk mengecek minimal satu elemen pada array yang memenuhi kondisi
// EVERY: mengekcek semua elemen yang sesuai dengan kondisi

// 14. Gunakan some untuk memeriksa apakah panjang beberapa nama lebih besar dari tujuh dalam array nama.
// answer here
const someNameLongerThan7: boolean = names.some(name => name.length > 7);
console.log(someNameLongerThan7);

// 15. Gunakan every untuk memeriksa apakah semua negara mengandung kata land.
// answer here
const everyCountryContainsLand: boolean = countries.every(ctr => ctr.includes('land'));
console.log(everyCountryContainsLand);

// 16. Jelaskan perbedaan antara find dan findIndex
// find :mengembalikan elemen pertama pada array yang memenuhi kondisi dan return undefined jika elemen tidak ditemukan
// findIndex :mengembalikan index elemen pertama pada array yang memenuhi kondisi dan -1 jika tidak ditemukan

// 17. Gunakan find untuk menemukan negara pertama yang hanya berisi enam huruf dalam array negara
// answer here
const firstCountryContain6Letter: (string | undefined) = countries.find(ctr => ctr.length === 6);
console.log(firstCountryContain6Letter);

// 18. Gunakan findIndex untuk menemukan posisi negara pertama yang hanya berisi enam huruf dalam array negara.
// answer here
const firstIndexCountryContain6Letter: number = countries.findIndex(ctr => ctr.length === 6);
console.log(firstIndexCountryContain6Letter);
// 19. Gunakan findIndex untuk menemukan posisi Norwegia jika tidak ada dalam array anda akan mendapatkan -1.
// answer here
const norwayIndex: number = countries.findIndex(ctr => ctr === 'Norway');
console.log(norwayIndex);

// 20. Gunakan findIndex untuk menemukan posisi Rusia jika tidak ada dalam array anda akan mendapatkan -1.
// answer here
const russiaIndex: number = countries.findIndex(ctr => ctr === 'Russia');
console.log(russiaIndex);
