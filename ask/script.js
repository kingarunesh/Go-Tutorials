let a = 1;

function main() {
    test();

    console.log(a);
    a = 2;
    console.log(a);
}

main();

function test() {
    console.log("test fun : ", a);
}
