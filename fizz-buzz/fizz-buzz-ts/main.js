function fizzBuzz(n) {
    var result = [];
    for (var i = 1; i <= n; i++) {
        if (i % 15 === 0) {
            result.push("FizzBuzz");
            continue;
        }
        if (i % 3 === 0) {
            result.push("Fizz");
            continue;
        }
        if (i % 5 === 0) {
            result.push("Buzz");
            continue;
        }
        result.push(i.toString());
    }
    return result;
}
function test() {
    console.log(fizzBuzz(3)); // -> ["1","2","Fizz"]
    console.log(fizzBuzz(5)); // -> ["1","2","Fizz","4","Buzz"]
    console.log(fizzBuzz(15)); // -> ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
}
test();
