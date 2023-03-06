function maximumWealth(accounts) {
    var richestNumber = 0;
    for (var _i = 0, accounts_1 = accounts; _i < accounts_1.length; _i++) {
        var account = accounts_1[_i];
        var sumAccount = account.reduce(function (sum, num) { return sum + num; }, 0);
        if (richestNumber < sumAccount) {
            richestNumber = sumAccount;
        }
    }
    return richestNumber;
}
;
function test() {
    console.log(maximumWealth([[1, 2, 3], [3, 2, 1]])); // -> 6
    console.log(maximumWealth([[1, 5], [7, 3], [3, 5]])); // -> 10
    console.log(maximumWealth([[2, 8, 7], [7, 1, 3], [1, 9, 5]])); // -> 17
}
;
test();
