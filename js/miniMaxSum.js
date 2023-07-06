function miniMaxSum(arr) {
  let sum = 0;
  let total = [];
  // Write your code here
  for (let i = 0; i < arr.length; i++) {
    sum += arr[i];
    total.push(sum);
  }

  total.map((value, index) => {
    console.log(arr[index]);
    // value -= arr[index];
    // console.log(value);
  });
}

miniMaxSum([7, 69, 2, 221, 8974]);
