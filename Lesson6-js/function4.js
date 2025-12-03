 function createFunc() {
 const x = 20;
 function f() {
 return x; // this `x` refers to the local `x` above
 }
 return f;
 }
 var f1 = createFunc();
 var y = f1(); // 20
 var q = (x) => x*3;
 var triple = q(6); // 18
