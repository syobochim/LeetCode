function runningSum(nums) {
    var result = [];
    var prevSum = 0;
    nums.forEach(function (num, i) {
        prevSum = prevSum + num;
        result.push(prevSum);
    });
    return result;
}
;
function test() {
    console.log(runningSum([1, 2, 3, 4]));
    console.log(runningSum([1, 1, 1, 1, 1]));
    console.log(runningSum([3, 1, 2, 10, 1]));
}
test();
