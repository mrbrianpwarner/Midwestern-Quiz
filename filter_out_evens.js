console.log("Filter out evens");
array = [3,5,2,4,4,4,4,49,6,1,8,54];
console.log("Pre-filter:\t" + array);
// Use the built in filter function. Only keep the odds
array = array.filter(num => num % 2 != 0);
console.log("Post-filter:\t" + array);