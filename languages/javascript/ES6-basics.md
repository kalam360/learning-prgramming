## 1. var, let, const

## Class and object in js

- Use ```constructor``` function and ```new``` operator

```js
// create multiple object in js
const person1 = {
    name: "John",
    walk(){
        console.log(this.name)
    }
}

const person2 = {
    name: "Mitchel", 
    walk() {
        console.log(this.name)
    }
}

person1.walk()
person2.walk()


// Now replicating this with Class object
class Person {
    constructor(name){
        this.name = name
    }
    walk(){
        console.log(this.name)
    }
}

const person3 = new Person("John")
const person4 = new Person("Mitchel")

person3.walk()
person4.walk()

```


## Inheritence

- Remember what ```this``` is in parent and child class 
- use ```super``` function in child class to access the parent class

```js
// inheritence
     
class Person {
    constructor(name){
        this.name = name
    }
    talk(){
        console.log(`${this.name} is talking`)
    }
}

class Teacher extends Person {
    constructor(name, degree){
        super(name)
        this.degree= degree
    }

    teach(){
       console.log( `${this.name} is a Teacher and his degree is ${this.degree}`)
    }
}

const p1 = new Person("Siam")
const p2 = new Teacher("Rafsan", "MSc")

p1.talk()
p2.talk()
p2.teach()
```
## modules

- Split the code into modules. 
- understand ```export```, ```default```, and named ```export``` 