function timeConversion(s) {
  // Write your code here
  let final = "";
  if (s[0] + s[1] == "12" && s[8] + s[9] == "AM") {
    let first = s[0] + s[1];
    first = "00";
    final = first + s[2] + s[3] + s[4] + s[5] + s[6] + s[7];
  } else if (s[0] + s[1] >= 1 && s[0] + s[1] <= 12 && s[8] + s[9] == "AM") {
    final = s[0] + s[1] + s[2] + s[3] + s[4] + s[5] + s[6] + s[7];
  } else if (s[0] + s[1] >= 1 && s[0] + s[1] <= 11 && s[8] + s[9] == "PM") {
    let sum = parseInt(s[0] + s[1]) + 12;
    final = sum + s[2] + s[3] + s[4] + s[5] + s[6] + s[7];
  }

  return final;
}

// 12:00 am -> 00:00
// 01:00 - 12:00 am -> keep just remove am
// 01:00 - 11:00 pm -> sum to 12
timeConversion("12:00:00AM");
timeConversion("01:00:00AM");
timeConversion("07:05:45PM");
