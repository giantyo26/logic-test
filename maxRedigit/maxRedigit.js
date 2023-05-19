const maxRedigit = (num) => {
    arrNum = [];
    if (num >= 100 && num <= 999) {
        while (num != 0) {
            arrNum.push(num % 10);
            num = Math.floor(num / 10);
        }
        console.log(arrNum.sort().reverse().join(''))
    } else {
        console.log(null)
    }
}
maxRedigit(472)
maxRedigit(-1)
maxRedigit(1000)
maxRedigit(123)
