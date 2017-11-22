function main() {
  var n = parseInt(readLine());
  ar = readLine().split(' ');
  ar = ar.map(Number);
  var result = birthdayCakeCandles(n, ar);
  process.stdout.write('' + result + '\n');
}

function birthdayCakeCandles(n, ar) {
  const tallest = ar.reduce((acc, curr) => {
    if (curr > acc) {
      return curr;
    }

    return acc;
  }, 0);

  return ar.filter(item => item === tallest).length;
}
