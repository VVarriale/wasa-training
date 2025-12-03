 function makeMultiple(x){
 var mult = 2;
 return mult * x;
 }
 //var wrong=mult; //multisfunction-scoped
 var y = makeMultiple(5); // 10
 var f = makeMultiple; // functionmakeMultiple(x){varmult=2;returnmult*x}
 var n = f(100); // 200
 • makeMultiple()referstothefunctioninvocation
 • makeMultiplereferstotheobjectfunction
