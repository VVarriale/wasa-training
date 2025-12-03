var a; //function-scoped
 var b = 1;
 b += 5; //reassigned,now it's6
 let c; //block-scoped
 let d = 3/2; //it's 1.5
 d = "Hello"; // reassigned,changedtype
 d += " world";//now"Helloworld"
 d = d + 1; //now"Helloworld1"
 const k = 30; //block-scoped,cannotbereassigned
 x = "foo"; // deprecated assignment(globalscope)
